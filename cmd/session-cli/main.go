package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/BelKirill/vikunja-mcp/internal/openai"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var (
	// Global session manager
	sessionManager openai.SessionManager
	logger         *log.Logger
	
	// Global flags
	verbose bool
	timeout int
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "session-cli",
	Short: "Session Manager CLI Tool",
	Long: `A command-line tool for testing and debugging the OpenAI session manager.
	
This tool provides direct access to session CRUD operations, metrics,
and debugging capabilities for the vikunja-mcp session management system.`,
	PersistentPreRun: initSessionManager,
	PersistentPostRun: cleanupSessionManager,
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose logging")
	rootCmd.PersistentFlags().IntVarP(&timeout, "timeout", "t", 30, "Default session timeout in minutes")
	
	// Add all subcommands
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(statsCmd)
	rootCmd.AddCommand(cleanupCmd)
	rootCmd.AddCommand(healthCmd)
	rootCmd.AddCommand(shutdownCmd)
}

func initSessionManager(cmd *cobra.Command, args []string) {
	// Initialize logger
	logger = log.New(os.Stderr)
	if verbose {
		logger.SetLevel(log.DebugLevel)
	} else {
		logger.SetLevel(log.InfoLevel)
	}

	// Initialize session manager
	config := openai.SessionConfig{
		DefaultTimeout:       time.Duration(timeout) * time.Minute,
		CleanupInterval:      1 * time.Minute,
		MaxSessions:          100,
		MaxSessionsPerUser:   5,
		AutoCleanup:          true,
		EnableMetrics:        true,
		LogSessionOperations: verbose,
	}

	sessionManager = openai.NewSessionManager(config, logger)
}

func cleanupSessionManager(cmd *cobra.Command, args []string) {
	if sessionManager != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		sessionManager.Shutdown(ctx)
	}
}

// Create command
var createCmd = &cobra.Command{
	Use:   "create <user_id>",
	Short: "Create a new session",
	Long: `Create a new session for the specified user.
	
You can optionally specify a custom timeout and thread ID.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		userID := args[0]
		
		timeoutMinutes, _ := cmd.Flags().GetInt("timeout")
		threadID, _ := cmd.Flags().GetString("thread")
		energy, _ := cmd.Flags().GetString("energy")
		mode, _ := cmd.Flags().GetString("mode")
		project, _ := cmd.Flags().GetString("project")
		
		opts := openai.SessionOptions{
			UserID: userID,
			TimeoutDuration: time.Duration(timeoutMinutes) * time.Minute,
			ThreadID: threadID,
			InitialContext: &openai.SessionContext{
				CurrentEnergy:   energy,
				CurrentMode:     mode,
				CurrentProject:  project,
				CompletionRate:  1.0,
				EnergyPatterns:  make(map[string]float64),
				RecentTasks:     []int{},
				RecentProjects:  []string{},
				PreferredTasks:  []string{},
				AvoidedTasks:    []string{},
			},
		}

		session, err := sessionManager.CreateSession(context.Background(), opts)
		if err != nil {
			return fmt.Errorf("failed to create session: %w", err)
		}

		if cmd.Flag("output").Value.String() == "json" {
			return printJSON(session)
		}
		
		printSessionSummary(session, "Session created successfully")
		return nil
	},
}

func init() {
	createCmd.Flags().IntP("timeout", "t", 30, "Session timeout in minutes")
	createCmd.Flags().StringP("thread", "", "", "OpenAI thread ID")
	createCmd.Flags().StringP("energy", "e", "medium", "Initial energy level (low, medium, high, social)")
	createCmd.Flags().StringP("mode", "m", "quick", "Initial work mode (quick, deep, admin)")
	createCmd.Flags().StringP("project", "p", "", "Current project name")
	createCmd.Flags().StringP("output", "o", "summary", "Output format (summary, json)")
}

// Get command
var getCmd = &cobra.Command{
	Use:   "get <session_id>",
	Short: "Get session details",
	Long:  `Retrieve detailed information about a specific session including all metadata.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		sessionID := args[0]
		
		session, err := sessionManager.GetSession(context.Background(), sessionID)
		if err != nil {
			return fmt.Errorf("failed to get session: %w", err)
		}

		if cmd.Flag("output").Value.String() == "json" {
			return printJSON(session)
		}
		
		printSessionDetailed(session)
		return nil
	},
}

func init() {
	getCmd.Flags().StringP("output", "o", "detailed", "Output format (detailed, json)")
}

// Update command
var updateCmd = &cobra.Command{
	Use:   "update <session_id>",
	Short: "Update session context",
	Long: `Update session context data such as energy level, work mode, and focus metrics.
	
Available update fields:
  - energy: low, medium, high, social
  - mode: quick, deep, admin  
  - project: project name
  - duration: target session duration in minutes
  - focus_streak: number of consecutive successful sessions
  - completion_rate: task completion percentage (0.0-1.0)
  - context_switches: number of context switches today`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		sessionID := args[0]
		
		session, err := sessionManager.GetSession(context.Background(), sessionID)
		if err != nil {
			return fmt.Errorf("failed to get session: %w", err)
		}

		// Get update flags
		if energy, _ := cmd.Flags().GetString("energy"); energy != "" {
			session.Context.CurrentEnergy = energy
		}
		if mode, _ := cmd.Flags().GetString("mode"); mode != "" {
			session.Context.CurrentMode = mode
		}
		if project, _ := cmd.Flags().GetString("project"); project != "" {
			session.Context.CurrentProject = project
		}
		if duration, _ := cmd.Flags().GetInt("duration"); duration > 0 {
			session.Context.TargetDuration = duration
		}
		if streak, _ := cmd.Flags().GetInt("focus-streak"); streak >= 0 {
			session.Context.FocusStreak = streak
		}
		if rate, _ := cmd.Flags().GetFloat64("completion-rate"); rate >= 0 {
			session.Context.CompletionRate = rate
		}
		if switches, _ := cmd.Flags().GetInt("context-switches"); switches >= 0 {
			session.Context.ContextSwitches = switches
		}

		// Parse key=value pairs from positional args if any
		kvPairs, _ := cmd.Flags().GetStringArray("set")
		for _, pair := range kvPairs {
			parts := strings.SplitN(pair, "=", 2)
			if len(parts) == 2 {
				updateSessionField(session, parts[0], parts[1])
			}
		}

		session.UpdatedAt = time.Now()
		
		err = sessionManager.UpdateSession(context.Background(), session)
		if err != nil {
			return fmt.Errorf("failed to update session: %w", err)
		}

		if cmd.Flag("output").Value.String() == "json" {
			return printJSON(session)
		}
		
		printSessionSummary(session, "Session updated successfully")
		return nil
	},
}

func init() {
	updateCmd.Flags().StringP("energy", "e", "", "Energy level")
	updateCmd.Flags().StringP("mode", "m", "", "Work mode")
	updateCmd.Flags().StringP("project", "p", "", "Project name")
	updateCmd.Flags().IntP("duration", "d", 0, "Target duration in minutes")
	updateCmd.Flags().IntP("focus-streak", "s", -1, "Focus streak count")
	updateCmd.Flags().Float64P("completion-rate", "r", -1, "Completion rate (0.0-1.0)")
	updateCmd.Flags().IntP("context-switches", "c", -1, "Context switches count")
	updateCmd.Flags().StringArrayP("set", "", []string{}, "Set custom key=value pairs")
	updateCmd.Flags().StringP("output", "o", "summary", "Output format (summary, json)")
}

// Delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <session_id>",
	Short: "Delete a session",
	Long:  `Remove a session from the session manager.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		sessionID := args[0]
		
		// Optionally get session info before deletion
		if confirm, _ := cmd.Flags().GetBool("confirm"); !confirm {
			session, err := sessionManager.GetSession(context.Background(), sessionID)
			if err != nil {
				return fmt.Errorf("session not found: %w", err)
			}
			
			fmt.Printf("About to delete session:\n")
			fmt.Printf("  Session ID: %s\n", session.ID)
			fmt.Printf("  User ID: %s\n", session.UserID)
			fmt.Printf("  Status: %s\n", session.Status)
			fmt.Printf("  Created: %s\n", session.CreatedAt.Format("2006-01-02 15:04:05"))
			
			fmt.Print("Continue? (y/N): ")
			var response string
			fmt.Scanln(&response)
			if strings.ToLower(response) != "y" && strings.ToLower(response) != "yes" {
				fmt.Println("Deletion cancelled")
				return nil
			}
		}
		
		err := sessionManager.DeleteSession(context.Background(), sessionID)
		if err != nil {
			return fmt.Errorf("failed to delete session: %w", err)
		}

		fmt.Printf("Session %s deleted successfully\n", sessionID)
		return nil
	},
}

func init() {
	deleteCmd.Flags().BoolP("confirm", "y", false, "Skip confirmation prompt")
}

// List command
var listCmd = &cobra.Command{
	Use:   "list [user_id]",
	Short: "List sessions",
	Long: `List all active sessions or sessions for a specific user.
	
Without arguments, lists all active sessions.
With a user_id argument, lists only that user's sessions.`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var sessions []*openai.Session
		var err error

		if len(args) > 0 {
			userID := args[0]
			sessions, err = sessionManager.ListSessions(context.Background(), userID)
			if err != nil {
				return fmt.Errorf("failed to list sessions for user %s: %w", userID, err)
			}
		} else {
			sessions, err = sessionManager.ListActiveSessions(context.Background())
			if err != nil {
				return fmt.Errorf("failed to list active sessions: %w", err)
			}
		}

		output, _ := cmd.Flags().GetString("output")
		
		switch output {
		case "json":
			return printJSON(sessions)
		case "table":
			printSessionTable(sessions)
		case "simple":
			for _, session := range sessions {
				fmt.Printf("%s\t%s\t%s\n", session.ID, session.UserID, session.Status)
			}
		default:
			printSessionTable(sessions)
		}
		
		return nil
	},
}

func init() {
	listCmd.Flags().StringP("output", "o", "table", "Output format (table, json, simple)")
	listCmd.Flags().BoolP("all", "a", false, "Include expired sessions")
}

// Stats command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show session manager statistics",
	Long:  `Display comprehensive statistics about the session manager including metrics and performance data.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		stats := sessionManager.GetSessionStats(context.Background())
		metrics := sessionManager.GetMetrics(context.Background())

		output, _ := cmd.Flags().GetString("output")
		
		if output == "json" {
			data := map[string]interface{}{
				"stats":   stats,
				"metrics": metrics,
			}
			return printJSON(data)
		}

		printStats(stats, metrics)
		return nil
	},
}

func init() {
	statsCmd.Flags().StringP("output", "o", "detailed", "Output format (detailed, json)")
}

// Cleanup command
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Clean up expired sessions",
	Long:  `Force cleanup of expired sessions and show cleanup statistics.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dryRun, _ := cmd.Flags().GetBool("dry-run")
		
		if dryRun {
			fmt.Println("DRY RUN: Would perform cleanup operations")
		}
		
		// Expire sessions first
		expired, err := sessionManager.ExpireSessions(context.Background())
		if err != nil {
			return fmt.Errorf("failed to expire sessions: %w", err)
		}

		if !dryRun {
			// Then cleanup expired sessions
			cleaned, err := sessionManager.CleanupExpiredSessions(context.Background())
			if err != nil {
				return fmt.Errorf("failed to cleanup sessions: %w", err)
			}
			
			fmt.Printf("Expired %d sessions\n", expired)
			fmt.Printf("Cleaned up %d sessions\n", cleaned)
		} else {
			fmt.Printf("Would expire %d sessions\n", expired)
		}
		
		return nil
	},
}

func init() {
	cleanupCmd.Flags().BoolP("dry-run", "n", false, "Show what would be cleaned without doing it")
}

// Health command
var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Check session manager health",
	Long:  `Perform a health check on the session manager and display current status.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := sessionManager.HealthCheck(context.Background())
		if err != nil {
			fmt.Printf("❌ Health check FAILED: %v\n", err)
			return err
		}

		fmt.Println("✅ Health check PASSED")
		
		// Show current status
		stats := sessionManager.GetSessionStats(context.Background())
		fmt.Printf("📊 Status: %d active, %d total sessions\n", 
			stats.ActiveSessions, stats.TotalSessions)
		
		return nil
	},
}

// Shutdown command
var shutdownCmd = &cobra.Command{
	Use:   "shutdown",
	Short: "Shutdown session manager",
	Long:  `Gracefully shutdown the session manager, cleaning up resources.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		timeoutSecs, _ := cmd.Flags().GetInt("timeout")
		
		fmt.Println("Shutting down session manager...")
		
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSecs)*time.Second)
		defer cancel()

		err := sessionManager.Shutdown(ctx)
		if err != nil {
			return fmt.Errorf("shutdown failed: %w", err)
		}

		fmt.Println("✅ Session manager shutdown complete")
		return nil
	},
}

func init() {
	shutdownCmd.Flags().IntP("timeout", "t", 5, "Shutdown timeout in seconds")
}

// Helper functions
func updateSessionField(session *openai.Session, key, value string) {
	switch key {
	case "energy":
		session.Context.CurrentEnergy = value
	case "mode":
		session.Context.CurrentMode = value
	case "project":
		session.Context.CurrentProject = value
	case "duration":
		if minutes, err := strconv.Atoi(value); err == nil {
			session.Context.TargetDuration = minutes
		}
	case "focus_streak":
		if streak, err := strconv.Atoi(value); err == nil {
			session.Context.FocusStreak = streak
		}
	case "completion_rate":
		if rate, err := strconv.ParseFloat(value, 64); err == nil {
			session.Context.CompletionRate = rate
		}
	case "context_switches":
		if switches, err := strconv.Atoi(value); err == nil {
			session.Context.ContextSwitches = switches
		}
	}
}

func printJSON(data interface{}) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func printSessionSummary(session *openai.Session, message string) {
	fmt.Printf("✅ %s\n", message)
	fmt.Printf("📝 Session ID: %s\n", session.ID)
	fmt.Printf("👤 User ID: %s\n", session.UserID)
	fmt.Printf("🔗 Thread ID: %s\n", session.ThreadID)
	fmt.Printf("📊 Status: %s\n", session.Status)
	fmt.Printf("⚡ Energy: %s\n", session.Context.CurrentEnergy)
	fmt.Printf("🎯 Mode: %s\n", session.Context.CurrentMode)
	if session.Context.CurrentProject != "" {
		fmt.Printf("📁 Project: %s\n", session.Context.CurrentProject)
	}
	fmt.Printf("⏰ Expires: %v (in %v)\n", 
		session.ExpiresAt.Format("15:04:05"), 
		time.Until(session.ExpiresAt).Round(time.Second))
}

func printSessionDetailed(session *openai.Session) {
	printSessionSummary(session, "Session Details")
	
	fmt.Printf("\n📈 Context:\n")
	fmt.Printf("  🔥 Focus Streak: %d\n", session.Context.FocusStreak)
	fmt.Printf("  ✅ Completion Rate: %.1f%%\n", session.Context.CompletionRate*100)
	fmt.Printf("  🔄 Context Switches: %d\n", session.Context.ContextSwitches)
	fmt.Printf("  ⏱️  Target Duration: %d minutes\n", session.Context.TargetDuration)
	
	fmt.Printf("\n🕐 Timestamps:\n")
	fmt.Printf("  Created: %s\n", session.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("  Updated: %s\n", session.UpdatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("  Last Access: %s\n", session.LastAccessTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("  Expires: %s\n", session.ExpiresAt.Format("2006-01-02 15:04:05"))
	
	fmt.Printf("\n🏥 Health:\n")
	fmt.Printf("  Is Active: %t\n", session.IsActive())
	fmt.Printf("  Is Expired: %t\n", session.IsExpired())
	fmt.Printf("  Is Healthy: %t\n", session.Health.IsHealthy)
	if session.Health.ConnectionErrors > 0 {
		fmt.Printf("  ⚠️  Connection Errors: %d\n", session.Health.ConnectionErrors)
	}
}

func printSessionTable(sessions []*openai.Session) {
	if len(sessions) == 0 {
		fmt.Println("No sessions found")
		return
	}

	fmt.Printf("\n%-15s %-15s %-10s %-8s %-8s %-20s\n", 
		"SESSION_ID", "USER_ID", "STATUS", "ENERGY", "MODE", "EXPIRES")
	fmt.Println(strings.Repeat("-", 90))

	for _, session := range sessions {
		sessionIDShort := session.ID
		if len(sessionIDShort) > 12 {
			sessionIDShort = sessionIDShort[:12] + "..."
		}

		userIDShort := session.UserID
		if len(userIDShort) > 12 {
			userIDShort = userIDShort[:12] + "..."
		}

		expiresIn := time.Until(session.ExpiresAt).Round(time.Minute)
		expiresStr := expiresIn.String()
		if expiresIn < 0 {
			expiresStr = "EXPIRED"
		}

		fmt.Printf("%-15s %-15s %-10s %-8s %-8s %-20s\n",
			sessionIDShort,
			userIDShort,
			session.Status,
			session.Context.CurrentEnergy,
			session.Context.CurrentMode,
			expiresStr)
	}

	fmt.Printf("\nTotal: %d sessions\n", len(sessions))
}

func printStats(stats openai.SessionStats, metrics openai.SessionMetrics) {
	fmt.Println("📊 SESSION MANAGER STATISTICS")
	fmt.Println(strings.Repeat("=", 40))
	
	fmt.Printf("📈 Overview:\n")
	fmt.Printf("  Total Sessions:    %d\n", stats.TotalSessions)
	fmt.Printf("  Active Sessions:   %d\n", stats.ActiveSessions)
	fmt.Printf("  Idle Sessions:     %d\n", stats.IdleSessions)
	fmt.Printf("  Expired Sessions:  %d\n", stats.ExpiredSessions)
	fmt.Printf("  Error Sessions:    %d\n", stats.ErrorSessions)
	fmt.Printf("  Average Lifetime:  %v\n", stats.AverageLifetime)
	
	if !stats.OldestSession.IsZero() {
		fmt.Printf("  Oldest Session:    %v\n", stats.OldestSession.Format("15:04:05"))
	}
	if !stats.NewestSession.IsZero() {
		fmt.Printf("  Newest Session:    %v\n", stats.NewestSession.Format("15:04:05"))
	}

	fmt.Printf("\n⚡ Performance Metrics:\n")
	fmt.Printf("  Sessions Created:    %d\n", metrics.SessionsCreated)
	fmt.Printf("  Sessions Deleted:    %d\n", metrics.SessionsDeleted)
	fmt.Printf("  Sessions Expired:    %d\n", metrics.SessionsExpired)
	fmt.Printf("  Sessions Recovered:  %d\n", metrics.SessionsRecovered)
	fmt.Printf("  Error Count:         %d\n", metrics.ErrorCount)
	fmt.Printf("  Avg Create Time:     %v\n", metrics.AverageCreateTime)
	fmt.Printf("  Avg Access Time:     %v\n", metrics.AverageAccessTime)
	
	if metrics.MemoryUsage > 0 {
		fmt.Printf("  Memory Usage:        %d bytes\n", metrics.MemoryUsage)
	}
	if metrics.CacheHitRate > 0 {
		fmt.Printf("  Cache Hit Rate:      %.1f%%\n", metrics.CacheHitRate*100)
	}
	
	if !metrics.LastCleanupTime.IsZero() {
		fmt.Printf("  Last Cleanup:        %v\n", metrics.LastCleanupTime.Format("15:04:05"))
	}
	if !metrics.LastHealthCheckTime.IsZero() {
		fmt.Printf("  Last Health Check:   %v\n", metrics.LastHealthCheckTime.Format("15:04:05"))
	}
}