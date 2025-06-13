package mcp

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"strings"
	"testing"
)

// Helper to create a test server with custom reader/writer
func newTestServer() (*Server, *bytes.Buffer) {
	out := &bytes.Buffer{}
	s := &Server{
		name:     "test-server",
		version:  "0.1.0",
		tools:    make(map[string]Tool),
		handlers: make(map[string]ToolHandler),
		writer:   out,
	}
	return s, out
}

func TestRegisterToolAndList(t *testing.T) {
	s, _ := newTestServer()
	tool := Tool{
		Name:        "echo",
		Description: "Echo tool",
		InputSchema: map[string]interface{}{"type": "object"},
	}
	s.RegisterTool(tool, func(args map[string]interface{}) (interface{}, error) {
		return args, nil
	})

	resp := s.handleToolsList(JSONRPCRequest{ID: 1})
	result, ok := resp.Result.(map[string]interface{})
	if !ok {
		t.Fatalf("expected map result")
	}
	tools, ok := result["tools"].([]Tool)
	if !ok {
		// fallback: decode from []interface{}
		rawTools, ok := result["tools"].([]interface{})
		if !ok || len(rawTools) != 1 {
			t.Fatalf("expected 1 tool, got %v", result["tools"])
		}
		toolMap, ok := rawTools[0].(map[string]interface{})
		if !ok || toolMap["name"] != "echo" {
			t.Fatalf("tool not found in list")
		}
	} else if len(tools) != 1 || tools[0].Name != "echo" {
		t.Fatalf("tool not found in list")
	}
}

func TestHandleInitialize(t *testing.T) {
	s, _ := newTestServer()
	req := JSONRPCRequest{
		ID:     1,
		Method: "initialize",
		Params: map[string]interface{}{
			"clientInfo": map[string]interface{}{
				"name": "test-client",
			},
		},
	}
	resp := s.handleInitialize(req)
	result, ok := resp.Result.(map[string]interface{})
	if !ok {
		t.Fatalf("expected map result")
	}
	serverInfo, ok := result["serverInfo"].(map[string]interface{})
	if !ok || serverInfo["name"] != "test-server" {
		t.Fatalf("expected serverInfo with name")
	}
}

func TestHandleToolCall_Success(t *testing.T) {
	s, _ := newTestServer()
	s.RegisterTool(Tool{
		Name:        "add",
		Description: "Add numbers",
		InputSchema: map[string]interface{}{},
	}, func(args map[string]interface{}) (interface{}, error) {
		a, _ := args["a"].(float64)
		b, _ := args["b"].(float64)
		return map[string]interface{}{"sum": a + b}, nil
	})

	req := JSONRPCRequest{
		ID:     2,
		Method: "tools/call",
		Params: map[string]interface{}{
			"name": "add",
			"arguments": map[string]interface{}{
				"a": 2,
				"b": 3,
			},
		},
	}
	resp := s.handleToolCall(req)
	if resp.Error != nil {
		t.Fatalf("unexpected error: %v", resp.Error)
	}
	result, ok := resp.Result.(map[string]interface{})
	if !ok {
		t.Fatalf("expected map result")
	}
	_, ok = result["content"].([]map[string]interface{})
	if !ok && result["content"] != nil {
		// fallback: decode from []interface{}
		rawContent, ok := result["content"].([]interface{})
		if !ok || len(rawContent) == 0 {
			t.Fatalf("expected content")
		}
	}
}

func TestHandleToolCall_ToolNotFound(t *testing.T) {
	s, _ := newTestServer()
	req := JSONRPCRequest{
		ID:     3,
		Method: "tools/call",
		Params: map[string]interface{}{
			"name": "missing",
		},
	}
	resp := s.handleToolCall(req)
	if resp.Error == nil || resp.Error.Code != -32601 {
		t.Fatalf("expected tool not found error")
	}
}

func TestHandleToolCall_InvalidParams(t *testing.T) {
	s, _ := newTestServer()
	req := JSONRPCRequest{
		ID:     4,
		Method: "tools/call",
		Params: "not a map",
	}
	resp := s.handleToolCall(req)
	if resp.Error == nil || resp.Error.Code != -32602 {
		t.Fatalf("expected invalid params error")
	}
}

func TestHandleToolCall_HandlerError(t *testing.T) {
	s, _ := newTestServer()
	s.RegisterTool(Tool{
		Name:        "fail",
		Description: "Always fails",
		InputSchema: map[string]interface{}{},
	}, func(args map[string]interface{}) (interface{}, error) {
		return nil, errors.New("fail error")
	})
	req := JSONRPCRequest{
		ID:     5,
		Method: "tools/call",
		Params: map[string]interface{}{
			"name":      "fail",
			"arguments": map[string]interface{}{},
		},
	}
	resp := s.handleToolCall(req)
	if resp.Error == nil || resp.Error.Code != -32603 {
		t.Fatalf("expected handler error")
	}
}

func TestHandleRequest_MethodNotFound(t *testing.T) {
	s, _ := newTestServer()
	req := JSONRPCRequest{
		ID:     6,
		Method: "unknown",
	}
	resp := s.handleRequest(req)
	if resp.Error == nil || resp.Error.Code != -32601 {
		t.Fatalf("expected method not found error")
	}
}

func TestSendResponse(t *testing.T) {
	s, out := newTestServer()
	resp := JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      7,
		Result:  map[string]interface{}{"ok": true},
	}
	err := s.sendResponse(resp)
	if err != nil {
		t.Fatalf("sendResponse failed: %v", err)
	}
	var decoded JSONRPCResponse
	if err := json.Unmarshal(out.Bytes(), &decoded); err != nil {
		t.Fatalf("invalid JSON output: %v", err)
	}
	if decoded.ID != 7 {
		t.Fatalf("wrong response ID")
	}
}

func TestServe_ParsesAndHandlesRequests(t *testing.T) {
	// Simulate two requests: one valid, one invalid
	in := strings.NewReader(
		`{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"clientInfo":{"name":"cli"}}}
invalid json
`)
	out := &bytes.Buffer{}
	s := &Server{
		name:     "test",
		version:  "0.1",
		tools:    make(map[string]Tool),
		handlers: make(map[string]ToolHandler),
		reader:   bufio.NewScanner(in),
		writer:   out,
	}
	_ = s.Serve()
	lines := strings.Split(strings.TrimSpace(out.String()), "\n")
	if len(lines) < 2 {
		t.Fatalf("expected 2 responses, got %d", len(lines))
	}
	var resp1, resp2 JSONRPCResponse
	_ = json.Unmarshal([]byte(lines[0]), &resp1)
	_ = json.Unmarshal([]byte(lines[1]), &resp2)
	if resp1.ID != float64(1) || resp2.Error == nil {
		t.Fatalf("unexpected responses: %+v %+v", resp1, resp2)
	}
}

func TestFormatToolResult_Error(t *testing.T) {
	s, _ := newTestServer()
	// Channels can't be marshaled to JSON
	ch := make(chan int)
	out := s.formatToolResult(ch)
	if !strings.Contains(out, "Error formatting result") {
		t.Fatalf("expected error formatting result, got: %s", out)
	}
}
