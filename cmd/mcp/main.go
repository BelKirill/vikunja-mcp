package main

import (
	"os"

	"github.com/BelKirill/vikunja-mcp/internal/config"
	"github.com/BelKirill/vikunja-mcp/internal/focus"
	"github.com/BelKirill/vikunja-mcp/pkg/mcp"
	"github.com/charmbracelet/log"
)

func main() {
	// Load configuration once (sets log level internally)
	_, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration", "error", err)
	}

	// Force logs to stderr for MCP (after config sets the level)
	log.SetOutput(os.Stderr)
	log.SetFormatter(log.JSONFormatter)

	log.Info("=== MCP SERVER STARTING ===")

	// Create MCP server
	server := mcp.NewServer("vikunja-mcp", "0.1.0")

	// Register tools (no config passing needed!)
	err = focus.RegisterMCPTools(server)
	if err != nil {
		log.Fatal("Failed to register MCP tools", "error", err)
	}

	log.Info("=== TOOLS REGISTERED, STARTING SERVE ===")

	// Run MCP server over stdio
	if err := server.Serve(); err != nil {
		log.Fatal("MCP server error", "error", err)
	}
}
