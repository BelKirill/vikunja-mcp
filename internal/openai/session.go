package openai

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/charmbracelet/log"
)

// SessionStatus represents the current state of a conversation session
type SessionStatus string

const (
	SessionStatusActive  SessionStatus = "active"
	SessionStatusIdle    SessionStatus = "idle"
	SessionStatusExpired SessionStatus = "expired"
	SessionStatusError   SessionStatus = "error"
)

// IsValid checks if the session status is valid
func (s SessionStatus) IsValid() bool {
	log.Debug("Checking if session status is valid", "status", s)
	switch s {
	case SessionStatusActive, SessionStatusIdle, SessionStatusExpired, SessionStatusError:
		return true
	default:
		log.Warn("Invalid session status encountered", "status", s)
		return false
	}
}

func (s SessionStatus) String() string {
	return string(s)
}

// SessionContext tracks the current work state and context
type SessionContext struct {
	// Current focus session details
	CurrentEnergy    string    `json:"current_energy"`     // low, medium, high, social
	CurrentMode      string    `json:"current_mode"`       // quick, deep, admin
	CurrentProject   string    `json:"current_project"`    // active project name/ID
	SessionStartTime time.Time `json:"session_start_time"` // when current focus session started
	TargetDuration   int       `json:"target_duration"`    // planned session length in minutes

	// Recent work patterns for AI context
	RecentTasks    []int    `json:"recent_tasks"`    // last 5-10 completed task IDs
	RecentProjects []string `json:"recent_projects"` // projects worked on today
	CompletionRate float64  `json:"completion_rate"` // recent task completion percentage

	// ADHD-specific context
	LastBreakTime   time.Time `json:"last_break_time"`  // for hyperfocus management
	FocusStreak     int       `json:"focus_streak"`     // consecutive successful sessions
	ContextSwitches int       `json:"context_switches"` // count for current day

	// AI learning context
	PreferredTasks []string           `json:"preferred_tasks"` // task types that work well
	AvoidedTasks   []string           `json:"avoided_tasks"`   // task types consistently skipped
	EnergyPatterns map[string]float64 `json:"energy_patterns"` // energy level success rates
}

// SessionMetadata tracks conversation-specific information
type SessionMetadata struct {
	// OpenAI conversation tracking
	MessageCount    int       `json:"message_count"`     // total messages in thread
	LastMessageTime time.Time `json:"last_message_time"` // timestamp of last interaction
	TokensUsed      int       `json:"tokens_used"`       // approximate token usage

	// Conversation context
	TopicsDiscussed []string `json:"topics_discussed"` // main conversation topics
	TasksRequested  []int    `json:"tasks_requested"`  // tasks Claude recommended
	TasksCompleted  []int    `json:"tasks_completed"`  // tasks marked done in session

	// Quality metrics
	Sentiment        string  `json:"sentiment"`         // positive, neutral, frustrated
	SuccessRate      float64 `json:"success_rate"`      // task completion vs. requested
	UserSatisfaction string  `json:"user_satisfaction"` // high, medium, low
}

// SessionHealth tracks session monitoring and recovery information
type SessionHealth struct {
	// Connection health
	LastPingTime     time.Time `json:"last_ping_time"`     // last successful API call
	ConnectionErrors int       `json:"connection_errors"`  // error count since last reset
	LastErrorTime    time.Time `json:"last_error_time"`    // timestamp of most recent error
	LastErrorMessage string    `json:"last_error_message"` // error details for debugging

	// Performance metrics
	AverageResponseTime time.Duration `json:"average_response_time"` // API response time
	TimeoutCount        int           `json:"timeout_count"`         // timeout occurrences
	RetryCount          int           `json:"retry_count"`           // retry attempts

	// Recovery state
	RecoveryAttempts int       `json:"recovery_attempts"`  // attempts to restore session
	LastRecoveryTime time.Time `json:"last_recovery_time"` // when last recovery happened
	IsHealthy        bool      `json:"is_healthy"`         // overall health status
}

// Session represents a complete OpenAI conversation session with persistence
type Session struct {
	// Core identification
	ID       string `json:"id"`        // unique session identifier
	ThreadID string `json:"thread_id"` // OpenAI thread ID
	UserID   string `json:"user_id"`   // user identifier

	// Lifecycle management
	Status         SessionStatus `json:"status"`           // current session state
	CreatedAt      time.Time     `json:"created_at"`       // session creation time
	UpdatedAt      time.Time     `json:"updated_at"`       // last modification time
	LastAccessTime time.Time     `json:"last_access_time"` // last user interaction
	ExpiresAt      time.Time     `json:"expires_at"`       // when session expires

	// Session context and metadata
	Context  SessionContext  `json:"context"`  // work state and patterns
	Metadata SessionMetadata `json:"metadata"` // conversation tracking
	Health   SessionHealth   `json:"health"`   // monitoring and recovery

	// Configuration
	TimeoutDuration time.Duration `json:"timeout_duration"` // session timeout setting
	AutoSave        bool          `json:"auto_save"`        // automatic persistence
	DebugMode       bool          `json:"debug_mode"`       // enhanced logging

	// Version control for data migration
	Version int `json:"version"` // structure version
}

// NewSession creates a new session with default values
func NewSession(threadID, userID string) *Session {
	now := time.Now()
	log.Info("Creating new session", "thread_id", threadID, "user_id", userID)
	return &Session{
		ID:              generateSessionID(),
		ThreadID:        threadID,
		UserID:          userID,
		Status:          SessionStatusActive,
		CreatedAt:       now,
		UpdatedAt:       now,
		LastAccessTime:  now,
		ExpiresAt:       now.Add(24 * time.Hour), // 24 hour default
		Context:         NewSessionContext(),
		Metadata:        NewSessionMetadata(),
		Health:          NewSessionHealth(),
		TimeoutDuration: 30 * time.Minute,
		AutoSave:        true,
		DebugMode:       false,
		Version:         1,
	}
}

// NewSessionContext creates initialized session context
func NewSessionContext() SessionContext {
	return SessionContext{
		RecentTasks:    make([]int, 0),
		RecentProjects: make([]string, 0),
		PreferredTasks: make([]string, 0),
		AvoidedTasks:   make([]string, 0),
		EnergyPatterns: make(map[string]float64),
		CompletionRate: 1.0, // optimistic start
	}
}

// NewSessionMetadata creates initialized session metadata
func NewSessionMetadata() SessionMetadata {
	return SessionMetadata{
		TopicsDiscussed:  make([]string, 0),
		TasksRequested:   make([]int, 0),
		TasksCompleted:   make([]int, 0),
		Sentiment:        "neutral",
		SuccessRate:      1.0,
		UserSatisfaction: "medium",
	}
}

// NewSessionHealth creates initialized session health
func NewSessionHealth() SessionHealth {
	return SessionHealth{
		LastPingTime:        time.Now(),
		AverageResponseTime: 1 * time.Second,
		IsHealthy:           true,
	}
}

// IsExpired checks if the session has expired
func (s *Session) IsExpired() bool {
	expired := time.Now().After(s.ExpiresAt)
	log.Debug("Checking if session is expired", "session_id", s.ID, "expired", expired)
	return expired
}

// IsActive checks if the session is currently active
func (s *Session) IsActive() bool {
	active := s.Status == SessionStatusActive && !s.IsExpired()
	log.Debug("Checking if session is active", "session_id", s.ID, "active", active)
	return active
}

// UpdateLastAccess updates the last access time and session status
func (s *Session) UpdateLastAccess() {
	log.Info("Updating last access time", "session_id", s.ID)
	s.LastAccessTime = time.Now()
	s.UpdatedAt = time.Now()
	if s.Status == SessionStatusIdle {
		log.Debug("Session was idle, setting to active", "session_id", s.ID)
		s.Status = SessionStatusActive
	}
}

// MarkIdle transitions session to idle state
func (s *Session) MarkIdle() {
	log.Info("Marking session as idle", "session_id", s.ID)
	s.Status = SessionStatusIdle
	s.UpdatedAt = time.Now()
}

// MarkError sets session to error state with message
func (s *Session) MarkError(errorMsg string) {
	log.Error("Marking session as error", "session_id", s.ID, "error", errorMsg)
	s.Status = SessionStatusError
	s.Health.LastErrorMessage = errorMsg
	s.Health.LastErrorTime = time.Now()
	s.Health.ConnectionErrors++
	s.Health.IsHealthy = false
	s.UpdatedAt = time.Now()
}

// Extend prolongs the session expiration time
func (s *Session) Extend(duration time.Duration) {
	log.Info("Extending session expiration", "session_id", s.ID, "duration", duration)
	s.ExpiresAt = time.Now().Add(duration)
	s.UpdatedAt = time.Now()
}

// Validate checks session data integrity
func (s *Session) Validate() error {
	log.Debug("Validating session data", "session_id", s.ID)
	if s.ID == "" {
		log.Error("Session ID cannot be empty")
		return fmt.Errorf("session ID cannot be empty")
	}
	if s.ThreadID == "" {
		log.Error("OpenAI thread ID cannot be empty")
		return fmt.Errorf("OpenAI thread ID cannot be empty")
	}
	if s.UserID == "" {
		log.Error("User ID cannot be empty")
		return fmt.Errorf("user ID cannot be empty")
	}
	if !s.Status.IsValid() {
		log.Error("Invalid session status", "status", s.Status)
		return fmt.Errorf("invalid session status: %s", s.Status)
	}
	if s.Version < 1 {
		log.Error("Invalid session version", "version", s.Version)
		return fmt.Errorf("invalid version: %d", s.Version)
	}
	return nil
}

// String returns a string representation for debugging
func (s *Session) String() string {
	return fmt.Sprintf("Session{ID: %s, ThreadID: %s, Status: %s, User: %s, Created: %s}",
		s.ID, s.ThreadID, s.Status, s.UserID, s.CreatedAt.Format(time.RFC3339))
}

// generateSessionID creates a unique session identifier
func generateSessionID() string {
	// Use current timestamp + random component for uniqueness
	return fmt.Sprintf("sess_%d", time.Now().UnixNano())
}

// SessionList represents a collection of sessions for management
type SessionList struct {
	Sessions []Session `json:"sessions"`
	Total    int       `json:"total"`
	Active   int       `json:"active"`
	Expired  int       `json:"expired"`
}

// Custom JSON marshaling for time.Duration fields
func (s *Session) MarshalJSON() ([]byte, error) {
	log.Debug("Marshaling session to JSON", "session_id", s.ID)
	type Alias Session
	return json.Marshal(&struct {
		*Alias
		TimeoutDurationMs int64 `json:"timeout_duration_ms"`
		AvgResponseTimeMs int64 `json:"avg_response_time_ms"`
	}{
		Alias:             (*Alias)(s),
		TimeoutDurationMs: s.TimeoutDuration.Milliseconds(),
		AvgResponseTimeMs: s.Health.AverageResponseTime.Milliseconds(),
	})
}

func (s *Session) UnmarshalJSON(data []byte) error {
	log.Debug("Unmarshaling session from JSON")
	type Alias Session
	aux := &struct {
		*Alias
		TimeoutDurationMs int64 `json:"timeout_duration_ms"`
		AvgResponseTimeMs int64 `json:"avg_response_time_ms"`
	}{
		Alias: (*Alias)(s),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		log.Error("Failed to unmarshal session JSON", "error", err)
		return err
	}

	s.TimeoutDuration = time.Duration(aux.TimeoutDurationMs) * time.Millisecond
	s.Health.AverageResponseTime = time.Duration(aux.AvgResponseTimeMs) * time.Millisecond

	return nil
}
