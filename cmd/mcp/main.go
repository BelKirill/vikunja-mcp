package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BelKirill/vikunja-mcp/internal/focus"
	"github.com/BelKirill/vikunja-mcp/internal/openai"
	"github.com/BelKirill/vikunja-mcp/pkg/mcp"
	"github.com/charmbracelet/log"
)

func main() {
	// Force logs to stderr for MCP
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)

	log.Info("=== MCP SERVER STARTING ===")

	// Initialize session management components
	sessionManager, threadManager, err := initializeSessionComponents()
	if err != nil {
		log.Error("Failed to initialize session components", "error", err)
		// Continue without sessions if initialization fails
		sessionManager = nil
		threadManager = nil
	}

	// Create MCP server
	server := mcp.NewServer("vikunja-mcp", "0.1.0")

	// Create focus server config with session support
	focusConfig := focus.ServerConfig{
		SessionManager:  sessionManager,
		ThreadManager:   threadManager,
		EnableSessions:  sessionManager != nil,
		// SessionDefaults: getDefaultSessionConfig(),
	}

	// Register tools with session support
	if sessionManager != nil {
		log.Info("=== SESSION MANAGEMENT ENABLED ===")
		// err = focus.RegisterMCPTools(server) // Your existing function
		err = focus.RegisterMCPToolsWithSessions(server, focusConfig)
	} else {
		log.Info("=== RUNNING WITHOUT SESSIONS ===")
		err = focus.RegisterMCPTools(server) // Your existing function
	}

	if err != nil {
		log.Fatal("Failed to register MCP tools", "error", err)
	}

	log.Info("=== TOOLS REGISTERED, STARTING SERVE ===")

	// Setup graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle shutdown signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Info("=== SHUTDOWN SIGNAL RECEIVED ===")
		cancel()
	}()

	// Start server in a goroutine
	serverDone := make(chan error, 1)
	go func() {
		serverDone <- server.Serve()
	}()

	// Wait for shutdown signal or server error
	select {
	case err := <-serverDone:
		if err != nil {
			log.Error("MCP server error", "error", err)
		}
	case <-ctx.Done():
		log.Info("=== INITIATING GRACEFUL SHUTDOWN ===")
	}

	// Cleanup session components
	cleanupSessionComponents(sessionManager, threadManager)

	log.Info("=== MCP SERVER SHUTDOWN COMPLETE ===")
}

// initializeSessionComponents initializes SessionManager and ThreadManager
func initializeSessionComponents() (openai.SessionManager, openai.ThreadManager, error) {
	// Get OpenAI API key
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Warn("OPENAI_API_KEY not set, session management will be disabled")
		return nil, nil, nil
	}

	log.Info("Initializing session management components")

	// Create ThreadManager
	threadManager, err := openai.NewThreadManager(apiKey)
	if err != nil {
		return nil, nil, err
	}

	// Create SessionManager with production configuration
	sessionConfig := openai.SessionConfig{
		DefaultTimeout:       30 * time.Minute,
		CleanupInterval:      5 * time.Minute,
		HealthCheckInterval:  2 * time.Minute,
		MaxSessions:          1000,
		MaxSessionsPerUser:   5,
		MaxIdleTime:          1 * time.Hour,
		AutoCleanup:          true,
		EnableMetrics:        true,
		PersistToDisk:        false, // In-memory for now
		LogLevel:             "info",
		LogSessionOperations: true,
	}

	logger := log.New(os.Stderr)
	sessionManager := openai.NewSessionManager(sessionConfig, logger)

	// Test session components
	if err := testSessionComponents(sessionManager, threadManager); err != nil {
		log.Error("Session components test failed", "error", err)
		return nil, nil, err
	}

	log.Info("Session management components initialized successfully")
	return sessionManager, threadManager, nil
}

// testSessionComponents performs a basic health check on session components
func testSessionComponents(sessionManager openai.SessionManager, threadManager openai.ThreadManager) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Test SessionManager health
	if err := sessionManager.HealthCheck(ctx); err != nil {
		return err
	}

	// Test ThreadManager by creating and immediately deleting a test thread
	threadID, err := threadManager.CreateThread(ctx)
	if err != nil {
		return err
	}

	// Clean up test thread
	if err := threadManager.DeleteThread(ctx, threadID); err != nil {
		log.Warn("Failed to cleanup test thread", "thread_id", threadID, "error", err)
	}

	return nil
}

// getDefaultSessionConfig returns default session configuration
// func getDefaultSessionConfig() focus.SessionDefaults {
// 	return focus.SessionDefaults{
// 		AutoCreateSessions: true,
// 		DefaultUserID:      "claude-user", // Default user ID for Claude interactions
// 		SessionTimeout:     30 * time.Minute,
// 		EnableContextual:   true,
// 		MaxContextHistory:  10, // Keep last 10 task completions in context
// 	}
// }

// cleanupSessionComponents performs graceful cleanup of session components
func cleanupSessionComponents(sessionManager openai.SessionManager, threadManager openai.ThreadManager) {
	if sessionManager == nil && threadManager == nil {
		return
	}

	log.Info("Cleaning up session components...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Stop health monitoring if running
	if threadManager != nil {
		threadManager.StopHealthMonitoring()
	}

	// Shutdown session manager
	if sessionManager != nil {
		if err := sessionManager.Shutdown(ctx); err != nil {
			log.Error("Error shutting down session manager", "error", err)
		}
	}

	log.Info("Session components cleanup complete")
}
