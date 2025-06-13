package mcp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/charmbracelet/log"
)

// MCP Protocol types
type JSONRPCRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id,omitempty"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

type JSONRPCResponse struct {
	JSONRPC string        `json:"jsonrpc"`
	ID      interface{}   `json:"id,omitempty"`
	Result  interface{}   `json:"result,omitempty"`
	Error   *JSONRPCError `json:"error,omitempty"`
}

type JSONRPCError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Tool struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	InputSchema map[string]interface{} `json:"inputSchema"`
}

type ToolHandler func(args map[string]interface{}) (interface{}, error)

type Server struct {
	name     string
	version  string
	tools    map[string]Tool
	handlers map[string]ToolHandler
	reader   *bufio.Scanner
	writer   io.Writer
}

func NewServer(name, version string) *Server {
	return &Server{
		name:     name,
		version:  version,
		tools:    make(map[string]Tool),
		handlers: make(map[string]ToolHandler),
		reader:   bufio.NewScanner(os.Stdin),
		writer:   os.Stdout,
	}
}

func (s *Server) RegisterTool(tool Tool, handler ToolHandler) {
	s.tools[tool.Name] = tool
	s.handlers[tool.Name] = handler
	log.Debug("registered MCP tool", "name", tool.Name)
}

func (s *Server) Serve() error {
	log.Info("starting MCP server", "name", s.name, "version", s.version)

	for s.reader.Scan() {
		line := s.reader.Text()
		if line == "" {
			continue
		}

		log.Debug("received MCP message", "raw", line)

		var req JSONRPCRequest
		if err := json.Unmarshal([]byte(line), &req); err != nil {
			log.Error("failed to parse JSON-RPC request", "error", err, "line", line)
			s.sendError(nil, -32700, "Parse error", nil)
			continue
		}

		response := s.handleRequest(req)
		if err := s.sendResponse(response); err != nil {
			log.Error("failed to send response", "error", err)
		}
	}

	if err := s.reader.Err(); err != nil {
		return fmt.Errorf("stdin read error: %w", err)
	}

	return nil
}

func (s *Server) handleRequest(req JSONRPCRequest) JSONRPCResponse {
	log.Info("handling MCP request", "method", req.Method, "id", req.ID)

	switch req.Method {
	case "initialize":
		return s.handleInitialize(req)
	case "tools/list":
		return s.handleToolsList(req)
	case "tools/call":
		return s.handleToolCall(req)
	default:
		return JSONRPCResponse{
			JSONRPC: "2.0",
			ID:      req.ID,
			Error: &JSONRPCError{
				Code:    -32601,
				Message: "Method not found",
			},
		}
	}
}

func (s *Server) handleInitialize(req JSONRPCRequest) JSONRPCResponse {
	// Parse params to get client info
	var params map[string]interface{}
	if req.Params != nil {
		if p, ok := req.Params.(map[string]interface{}); ok {
			params = p
		}
	}

	clientInfo := ""
	if params != nil {
		if name, ok := params["clientInfo"].(map[string]interface{}); ok {
			if n, ok := name["name"].(string); ok {
				clientInfo = n
			}
		}
	}

	log.Info("MCP client connected", "client", clientInfo)

	return JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      req.ID,
		Result: map[string]interface{}{
			"protocolVersion": "2024-11-05",
			"capabilities": map[string]interface{}{
				"tools": map[string]interface{}{},
			},
			"serverInfo": map[string]interface{}{
				"name":    s.name,
				"version": s.version,
			},
		},
	}
}

func (s *Server) handleToolsList(req JSONRPCRequest) JSONRPCResponse {
	tools := make([]Tool, 0, len(s.tools))
	for _, tool := range s.tools {
		tools = append(tools, tool)
	}

	log.Info("listing MCP tools", "count", len(tools))

	return JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      req.ID,
		Result: map[string]interface{}{
			"tools": tools,
		},
	}
}

func (s *Server) handleToolCall(req JSONRPCRequest) JSONRPCResponse {
	// Parse tool call parameters
	params, ok := req.Params.(map[string]interface{})
	if !ok {
		return s.errorResponse(req.ID, -32602, "Invalid params")
	}

	toolName, ok := params["name"].(string)
	if !ok {
		return s.errorResponse(req.ID, -32602, "Missing tool name")
	}

	handler, exists := s.handlers[toolName]
	if !exists {
		return s.errorResponse(req.ID, -32601, fmt.Sprintf("Tool not found: %s", toolName))
	}

	// Extract arguments
	args := make(map[string]interface{})
	if arguments, ok := params["arguments"].(map[string]interface{}); ok {
		args = arguments
	}

	log.Info("calling MCP tool", "tool", toolName, "args", args)

	// Call the handler
	result, err := handler(args)
	if err != nil {
		log.Error("tool execution failed", "tool", toolName, "error", err)
		return s.errorResponse(req.ID, -32603, fmt.Sprintf("Tool execution failed: %v", err))
	}

	log.Info("tool execution completed", "tool", toolName)

	return JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      req.ID,
		Result: map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": s.formatToolResult(result),
				},
			},
		},
	}
}

func (s *Server) errorResponse(id interface{}, code int, message string) JSONRPCResponse {
	return JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      id,
		Error: &JSONRPCError{
			Code:    code,
			Message: message,
		},
	}
}

func (s *Server) sendResponse(response JSONRPCResponse) error {
	data, err := json.Marshal(response)
	if err != nil {
		return err
	}

	log.Debug("sending MCP response", "response", string(data))

	_, err = fmt.Fprintf(s.writer, "%s\n", data)
	return err
}

func (s *Server) sendError(id interface{}, code int, message string, data interface{}) {
	response := JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      id,
		Error: &JSONRPCError{
			Code:    code,
			Message: message,
			Data:    data,
		},
	}
	err := s.sendResponse(response)
	if err != nil {
		log.Error("Error in sending response", "response", response, "err", err)
	}

}

func (s *Server) formatToolResult(result interface{}) string {
	// Format the result as a nice JSON string for Claude
	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error formatting result: %v", err)
	}
	return string(data)
}
