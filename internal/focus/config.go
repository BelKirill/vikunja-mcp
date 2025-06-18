package focus

import (
	"time"

	"github.com/BelKirill/vikunja-mcp/internal/openai"
)

// ServerConfig contains configuration for the focus server with session support
type ServerConfig struct {
	SessionManager  openai.SessionManager `json:"-"` // Don't serialize
	ThreadManager   openai.ThreadManager  `json:"-"` // Don't serialize
	EnableSessions  bool                  `json:"enable_sessions"`
	SessionDefaults SessionDefaults       `json:"session_defaults"`
}

// SessionDefaults contains default session behavior settings
type SessionDefaults struct {
	AutoCreateSessions bool          `json:"auto_create_sessions"`
	DefaultUserID      string        `json:"default_user_id"`
	SessionTimeout     time.Duration `json:"session_timeout"`
	EnableContextual   bool          `json:"enable_contextual"`
	MaxContextHistory  int           `json:"max_context_history"`
}

// SessionContext represents session-specific context for recommendations
type SessionContext struct {
	SessionID           string                 `json:"session_id"`
	UserID              string                 `json:"user_id"`
	ThreadID            string                 `json:"thread_id"`
	IsActive            bool                   `json:"is_active"`
	TaskHistory         []TaskCompletion       `json:"task_history"`
	CurrentProject      string                 `json:"current_project,omitempty"`
	CurrentContext      map[string]interface{} `json:"current_context,omitempty"`
	RecommendationCount int                    `json:"recommendation_count"`
	SessionStartTime    time.Time              `json:"session_start_time"`
	LastActivity        time.Time              `json:"last_activity"`
}

// TaskCompletion represents a completed task in session history
type TaskCompletion struct {
	TaskID      string                 `json:"task_id"`
	Title       string                 `json:"title"`
	CompletedAt time.Time              `json:"completed_at"`
	Duration    time.Duration          `json:"duration,omitempty"`
	Energy      string                 `json:"energy,omitempty"`
	Mode        string                 `json:"mode,omitempty"`
	Success     bool                   `json:"success"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// SessionAwareResponse wraps regular responses with session information
type SessionAwareResponse struct {
	Data           interface{}     `json:"data"`
	SessionContext *SessionContext `json:"session_context,omitempty"`
	SessionEnabled bool            `json:"session_enabled"`
	Message        string          `json:"message,omitempty"`
}

// FocusServer represents the enhanced focus server with session support
type FocusServer struct {
	config ServerConfig
	// Add your existing clients here when integrating
}

// NewFocusServer creates a new focus server with session support
func NewFocusServer(config ServerConfig) *FocusServer {
	return &FocusServer{
		config: config,
		// Initialize your existing clients here
	}
}
