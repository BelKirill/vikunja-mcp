package main

import (
	"github.com/BelKirill/vikunja-mcp/internal/focus"
	"github.com/BelKirill/vikunja-mcp/pkg/mcp"
	"github.com/charmbracelet/log"
)

func main() {

	// Create MCP server
	server := mcp.NewServer("vikunja-mcp", "0.1.0")

	// Register tools
	err := focus.RegisterMCPTools(server)
	if err != nil {
		log.Fatal("Failed to register MCP tools", "error", err)
	}

	// Run MCP server over stdio
	if err := server.Serve(); err != nil {
		log.Fatal("MCP server error", "error", err)
	}
}
