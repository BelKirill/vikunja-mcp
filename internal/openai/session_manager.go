package openai

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
)

// SessionManager defines the interface for managing OpenAI conversation sessions
type SessionManager interface {
	// Core CRUD operations
	CreateSession(ctx context.Context, opts SessionOptions) (*Session, error)
	GetSession(ctx context.Context, sessionID string) (*Session, error)
	UpdateSession(ctx context.Context, session *Session) error
	DeleteSession(ctx context.Context, sessionID string) error

	// Active session management
	GetActiveSession(ctx context.Context, userID string) (*Session, error)
	SetActiveSession(ctx context.Context, userID, sessionID string) error

	// Lifecycle management
	RefreshSession(ctx context.Context, sessionID string) error
	ExpireSession(ctx context.Context, sessionID string) error
	ExpireSessions(ctx context.Context) (int, error)
	CleanupExpiredSessions(ctx context.Context) (int, error)

	// Query operations
	ListSessions(ctx context.Context, userID string) ([]*Session, error)
	ListActiveSessions(ctx context.Context) ([]*Session, error)
	GetSessionStats(ctx context.Context) SessionStats

	// Health and monitoring
	HealthCheck(ctx context.Context) error
	GetMetrics(ctx context.Context) SessionMetrics

	// Shutdown and cleanup
	Shutdown(ctx context.Context) error
}

// SessionOptions contains parameters for creating new sessions
type SessionOptions struct {
	UserID          string          `json:"user_id"`
	ThreadID        string          `json:"thread_id"`        // optional: existing OpenAI thread
	TimeoutDuration time.Duration   `json:"timeout_duration"` // optional: custom timeout
	AutoSave        bool            `json:"auto_save"`        // optional: auto-persistence
	DebugMode       bool            `json:"debug_mode"`       // optional: enhanced logging
	InitialContext  *SessionContext `json:"initial_context"`  // optional: starting context
}

// SessionConfig contains configuration for the SessionManager
type SessionConfig struct {
	// Timing configuration
	DefaultTimeout      time.Duration `json:"default_timeout"`
	CleanupInterval     time.Duration `json:"cleanup_interval"`
	HealthCheckInterval time.Duration `json:"health_check_interval"`

	// Capacity limits
	MaxSessions        int           `json:"max_sessions"`
	MaxSessionsPerUser int           `json:"max_sessions_per_user"`
	MaxIdleTime        time.Duration `json:"max_idle_time"`

	// Behavior settings
	AutoCleanup   bool `json:"auto_cleanup"`
	EnableMetrics bool `json:"enable_metrics"`
	PersistToDisk bool `json:"persist_to_disk"`

	// Logging
	LogLevel             string `json:"log_level"`
	LogSessionOperations bool   `json:"log_session_operations"`
}

// SessionStats provides aggregate statistics about sessions
type SessionStats struct {
	TotalSessions   int           `json:"total_sessions"`
	ActiveSessions  int           `json:"active_sessions"`
	IdleSessions    int           `json:"idle_sessions"`
	ExpiredSessions int           `json:"expired_sessions"`
	ErrorSessions   int           `json:"error_sessions"`
	AverageLifetime time.Duration `json:"average_lifetime"`
	OldestSession   time.Time     `json:"oldest_session"`
	NewestSession   time.Time     `json:"newest_session"`
}

// SessionMetrics provides performance and health metrics
type SessionMetrics struct {
	// Operation counters
	SessionsCreated   int64 `json:"sessions_created"`
	SessionsDeleted   int64 `json:"sessions_deleted"`
	SessionsExpired   int64 `json:"sessions_expired"`
	SessionsRecovered int64 `json:"sessions_recovered"`

	// Performance metrics
	AverageCreateTime time.Duration `json:"average_create_time"`
	AverageAccessTime time.Duration `json:"average_access_time"`

	// Health metrics
	ErrorCount          int64     `json:"error_count"`
	LastCleanupTime     time.Time `json:"last_cleanup_time"`
	LastHealthCheckTime time.Time `json:"last_health_check_time"`

	// Resource usage
	MemoryUsage  int64   `json:"memory_usage_bytes"`
	CacheHitRate float64 `json:"cache_hit_rate"`
}

// SessionManagerImpl implements the SessionManager interface
type SessionManagerImpl struct {
	config SessionConfig
	logger *log.Logger

	// Session storage and synchronization
	sessions       map[string]*Session // sessionID -> Session
	userSessions   map[string][]string // userID -> []sessionID
	activeSessions map[string]string   // userID -> activeSessionID
	mutex          sync.RWMutex

	// Background processes
	cleanupTicker *time.Ticker
	healthTicker  *time.Ticker
	shutdownChan  chan struct{}

	// Metrics and monitoring
	metrics      SessionMetrics
	metricsMutex sync.RWMutex
}

// NewSessionManager creates a new SessionManager with the given configuration
func NewSessionManager(config SessionConfig, logger *log.Logger) SessionManager {
	// Pipe logger to stderr if not provided
	if logger == nil {
		logger = log.New(os.Stderr)
	}
	// Apply defaults to config
	if config.DefaultTimeout == 0 {
		config.DefaultTimeout = 30 * time.Minute
	}
	if config.CleanupInterval == 0 {
		config.CleanupInterval = 5 * time.Minute
	}
	if config.MaxSessions == 0 {
		config.MaxSessions = 1000
	}
	if config.MaxSessionsPerUser == 0 {
		config.MaxSessionsPerUser = 5
	}
	if config.MaxIdleTime == 0 {
		config.MaxIdleTime = 1 * time.Hour
	}

	sm := &SessionManagerImpl{
		config:         config,
		logger:         logger,
		sessions:       make(map[string]*Session),
		userSessions:   make(map[string][]string),
		activeSessions: make(map[string]string),
		shutdownChan:   make(chan struct{}),
		metrics:        SessionMetrics{},
	}

	// Start background processes if auto-cleanup is enabled
	if config.AutoCleanup {
		sm.startBackgroundProcesses()
	}

	logger.Info("SessionManager initialized",
		"max_sessions", config.MaxSessions,
		"default_timeout", config.DefaultTimeout,
		"auto_cleanup", config.AutoCleanup)

	return sm
}

// CreateSession creates a new session with the given options
func (sm *SessionManagerImpl) CreateSession(ctx context.Context, opts SessionOptions) (*Session, error) {
	startTime := time.Now()
	defer func() {
		sm.updateMetrics(func(m *SessionMetrics) {
			m.AverageCreateTime = time.Since(startTime)
		})
	}()

	// Validate options
	if opts.UserID == "" {
		return nil, fmt.Errorf("user ID is required")
	}

	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	// Check session limits
	if len(sm.sessions) >= sm.config.MaxSessions {
		return nil, ErrSessionLimitExceeded
	}

	userSessionCount := len(sm.userSessions[opts.UserID])
	if userSessionCount >= sm.config.MaxSessionsPerUser {
		return nil, ErrUserSessionLimitExceeded
	}

	// Generate session ID if not provided
	sessionID := generateSecureSessionID()

	// Generate thread ID if not provided
	threadID := opts.ThreadID
	if threadID == "" {
		threadID = generateThreadID()
	}

	// Apply defaults
	timeout := opts.TimeoutDuration
	if timeout == 0 {
		timeout = sm.config.DefaultTimeout
	}

	// Create session
	session := NewSession(threadID, opts.UserID)
	session.ID = sessionID
	session.TimeoutDuration = timeout
	session.AutoSave = opts.AutoSave
	session.DebugMode = opts.DebugMode

	// Set initial context if provided
	if opts.InitialContext != nil {
		session.Context = *opts.InitialContext
	}

	// Store session
	sm.sessions[sessionID] = session
	sm.userSessions[opts.UserID] = append(sm.userSessions[opts.UserID], sessionID)

	// Set as active session if user has no active session
	if sm.activeSessions[opts.UserID] == "" {
		sm.activeSessions[opts.UserID] = sessionID
	}

	// Update metrics
	sm.updateMetrics(func(m *SessionMetrics) {
		m.SessionsCreated++
	})

	if sm.config.LogSessionOperations {
		sm.logger.Info("Session created",
			"session_id", sessionID,
			"user_id", opts.UserID,
			"thread_id", threadID,
			"timeout", timeout)
	}

	return session, nil
}

// GetSession retrieves a session by ID
func (sm *SessionManagerImpl) GetSession(ctx context.Context, sessionID string) (*Session, error) {
	startTime := time.Now()
	defer func() {
		sm.updateMetrics(func(m *SessionMetrics) {
			m.AverageAccessTime = time.Since(startTime)
		})
	}()

	if sessionID == "" {
		return nil, ErrInvalidSessionID
	}

	sm.mutex.RLock()
	session, exists := sm.sessions[sessionID]
	sm.mutex.RUnlock()

	if !exists {
		return nil, ErrSessionNotFound
	}

	// Update last access time
	session.UpdateLastAccess()

	// Check if session is expired
	if session.IsExpired() {
		sm.logger.Warn("Attempted to access expired session", "session_id", sessionID)
		return nil, ErrSessionExpired
	}

	return session, nil
}

// UpdateSession updates an existing session
func (sm *SessionManagerImpl) UpdateSession(ctx context.Context, session *Session) error {
	if session == nil {
		return ErrInvalidSession
	}

	if err := session.Validate(); err != nil {
		return fmt.Errorf("invalid session: %w", err)
	}

	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	_, exists := sm.sessions[session.ID]
	if !exists {
		return ErrSessionNotFound
	}

	// Update timestamps
	session.UpdatedAt = time.Now()

	// Store updated session
	sm.sessions[session.ID] = session

	if sm.config.LogSessionOperations {
		sm.logger.Debug("Session updated",
			"session_id", session.ID,
			"status", session.Status,
			"user_id", session.UserID)
	}

	return nil
}

// DeleteSession removes a session
func (sm *SessionManagerImpl) DeleteSession(ctx context.Context, sessionID string) error {
	if sessionID == "" {
		return ErrInvalidSessionID
	}

	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	session, exists := sm.sessions[sessionID]
	if !exists {
		return ErrSessionNotFound
	}

	userID := session.UserID

	// Remove from main sessions map
	delete(sm.sessions, sessionID)

	// Remove from user sessions
	userSessions := sm.userSessions[userID]
	for i, id := range userSessions {
		if id == sessionID {
			sm.userSessions[userID] = append(userSessions[:i], userSessions[i+1:]...)
			break
		}
	}

	// Remove from active sessions if this was the active one
	if sm.activeSessions[userID] == sessionID {
		delete(sm.activeSessions, userID)
	}

	// Update metrics
	sm.updateMetrics(func(m *SessionMetrics) {
		m.SessionsDeleted++
	})

	if sm.config.LogSessionOperations {
		sm.logger.Info("Session deleted",
			"session_id", sessionID,
			"user_id", userID)
	}

	return nil
}

// GetActiveSession retrieves the currently active session for a user
func (sm *SessionManagerImpl) GetActiveSession(ctx context.Context, userID string) (*Session, error) {
	if userID == "" {
		return nil, ErrInvalidUserID
	}

	sm.mutex.RLock()
	activeSessionID, hasActive := sm.activeSessions[userID]
	sm.mutex.RUnlock()

	if !hasActive {
		return nil, ErrNoActiveSession
	}

	return sm.GetSession(ctx, activeSessionID)
}

// SetActiveSession sets the active session for a user
func (sm *SessionManagerImpl) SetActiveSession(ctx context.Context, userID, sessionID string) error {
	if userID == "" {
		return ErrInvalidUserID
	}
	if sessionID == "" {
		return ErrInvalidSessionID
	}

	// Verify session exists and belongs to user
	session, err := sm.GetSession(ctx, sessionID)
	if err != nil {
		return err
	}

	if session.UserID != userID {
		return ErrSessionNotOwnedByUser
	}

	sm.mutex.Lock()
	sm.activeSessions[userID] = sessionID
	sm.mutex.Unlock()

	if sm.config.LogSessionOperations {
		sm.logger.Info("Active session set",
			"user_id", userID,
			"session_id", sessionID)
	}

	return nil
}

// RefreshSession extends the session expiration time
func (sm *SessionManagerImpl) RefreshSession(ctx context.Context, sessionID string) error {
	session, err := sm.GetSession(ctx, sessionID)
	if err != nil {
		return err
	}

	session.Extend(sm.config.DefaultTimeout)

	return sm.UpdateSession(ctx, session)
}

// ExpireSession marks a session as expired
func (sm *SessionManagerImpl) ExpireSession(ctx context.Context, sessionID string) error {
	session, err := sm.GetSession(ctx, sessionID)
	if err != nil {
		return err
	}

	session.Status = SessionStatusExpired
	session.ExpiresAt = time.Now()

	sm.updateMetrics(func(m *SessionMetrics) {
		m.SessionsExpired++
	})

	return sm.UpdateSession(ctx, session)
}

// ExpireSessions marks all expired sessions as expired
func (sm *SessionManagerImpl) ExpireSessions(ctx context.Context) (int, error) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	expired := 0
	now := time.Now()

	for _, session := range sm.sessions {
		if session.IsExpired() && session.Status != SessionStatusExpired {
			session.Status = SessionStatusExpired
			session.UpdatedAt = now
			expired++
		}
	}

	sm.updateMetrics(func(m *SessionMetrics) {
		m.SessionsExpired += int64(expired)
	})

	if expired > 0 {
		sm.logger.Info("Sessions expired", "count", expired)
	}

	return expired, nil
}

// CleanupExpiredSessions removes expired sessions from memory
func (sm *SessionManagerImpl) CleanupExpiredSessions(ctx context.Context) (int, error) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	cleaned := 0
	now := time.Now()

	// Collect expired session IDs
	var expiredIDs []string
	for id, session := range sm.sessions {
		if session.IsExpired() || (session.Status == SessionStatusExpired) {
			expiredIDs = append(expiredIDs, id)
		}
	}

	// Remove expired sessions
	for _, sessionID := range expiredIDs {
		session := sm.sessions[sessionID]
		userID := session.UserID

		// Remove from maps
		delete(sm.sessions, sessionID)

		// Remove from user sessions
		userSessions := sm.userSessions[userID]
		for i, id := range userSessions {
			if id == sessionID {
				sm.userSessions[userID] = append(userSessions[:i], userSessions[i+1:]...)
				break
			}
		}

		// Remove from active sessions
		if sm.activeSessions[userID] == sessionID {
			delete(sm.activeSessions, userID)
		}

		cleaned++
	}

	sm.updateMetrics(func(m *SessionMetrics) {
		m.LastCleanupTime = now
	})

	if cleaned > 0 {
		sm.logger.Info("Expired sessions cleaned up", "count", cleaned)
	}

	return cleaned, nil
}

// ListSessions returns all sessions for a user
func (sm *SessionManagerImpl) ListSessions(ctx context.Context, userID string) ([]*Session, error) {
	if userID == "" {
		return nil, ErrInvalidUserID
	}

	sm.mutex.RLock()
	sessionIDs := sm.userSessions[userID]
	sm.mutex.RUnlock()

	sessions := make([]*Session, 0, len(sessionIDs))
	for _, sessionID := range sessionIDs {
		if session, err := sm.GetSession(ctx, sessionID); err == nil {
			sessions = append(sessions, session)
		}
	}

	return sessions, nil
}

// ListActiveSessions returns all currently active sessions
func (sm *SessionManagerImpl) ListActiveSessions(ctx context.Context) ([]*Session, error) {
	sm.mutex.RLock()
	activeSessionIDs := make([]string, 0, len(sm.activeSessions))
	for _, sessionID := range sm.activeSessions {
		activeSessionIDs = append(activeSessionIDs, sessionID)
	}
	sm.mutex.RUnlock()

	sessions := make([]*Session, 0, len(activeSessionIDs))
	for _, sessionID := range activeSessionIDs {
		if session, err := sm.GetSession(ctx, sessionID); err == nil {
			sessions = append(sessions, session)
		}
	}

	return sessions, nil
}

// GetSessionStats returns aggregate statistics
func (sm *SessionManagerImpl) GetSessionStats(ctx context.Context) SessionStats {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()

	stats := SessionStats{}
	now := time.Now()

	var totalLifetime time.Duration
	var oldestTime, newestTime time.Time

	for _, session := range sm.sessions {
		stats.TotalSessions++

		switch session.Status {
		case SessionStatusActive:
			stats.ActiveSessions++
		case SessionStatusIdle:
			stats.IdleSessions++
		case SessionStatusExpired:
			stats.ExpiredSessions++
		case SessionStatusError:
			stats.ErrorSessions++
		}

		// Calculate lifetime
		lifetime := now.Sub(session.CreatedAt)
		totalLifetime += lifetime

		// Track oldest and newest
		if oldestTime.IsZero() || session.CreatedAt.Before(oldestTime) {
			oldestTime = session.CreatedAt
		}
		if newestTime.IsZero() || session.CreatedAt.After(newestTime) {
			newestTime = session.CreatedAt
		}
	}

	if stats.TotalSessions > 0 {
		stats.AverageLifetime = totalLifetime / time.Duration(stats.TotalSessions)
	}

	stats.OldestSession = oldestTime
	stats.NewestSession = newestTime

	return stats
}

// GetMetrics returns performance metrics
func (sm *SessionManagerImpl) GetMetrics(ctx context.Context) SessionMetrics {
	sm.metricsMutex.RLock()
	defer sm.metricsMutex.RUnlock()

	return sm.metrics
}

// HealthCheck validates SessionManager health
func (sm *SessionManagerImpl) HealthCheck(ctx context.Context) error {
	sm.mutex.RLock()
	sessionCount := len(sm.sessions)
	sm.mutex.RUnlock()

	// Check if we're approaching limits
	if sessionCount > int(float64(sm.config.MaxSessions)*0.9) {
		sm.logger.Warn("Session count approaching limit",
			"current", sessionCount,
			"max", sm.config.MaxSessions)
	}

	sm.updateMetrics(func(m *SessionMetrics) {
		m.LastHealthCheckTime = time.Now()
	})

	return nil
}

// Shutdown gracefully shuts down the SessionManager
func (sm *SessionManagerImpl) Shutdown(ctx context.Context) error {
	sm.logger.Info("SessionManager shutting down...")

	// Stop background processes
	if sm.cleanupTicker != nil {
		sm.cleanupTicker.Stop()
	}
	if sm.healthTicker != nil {
		sm.healthTicker.Stop()
	}

	close(sm.shutdownChan)

	// Final cleanup
	_, err := sm.CleanupExpiredSessions(ctx)
	if err != nil {
		sm.logger.Error("Error during final cleanup", "error", err)
	}

	sm.logger.Info("SessionManager shutdown complete")
	return nil
}

// Helper methods

// startBackgroundProcesses starts cleanup and health check routines
func (sm *SessionManagerImpl) startBackgroundProcesses() {
	sm.cleanupTicker = time.NewTicker(sm.config.CleanupInterval)
	sm.healthTicker = time.NewTicker(sm.config.HealthCheckInterval)

	go func() {
		for {
			select {
			case <-sm.cleanupTicker.C:
				ctx := context.Background()
				if _, err := sm.CleanupExpiredSessions(ctx); err != nil {
					sm.logger.Error("Background cleanup failed", "error", err)
				}
			case <-sm.healthTicker.C:
				ctx := context.Background()
				if err := sm.HealthCheck(ctx); err != nil {
					sm.logger.Error("Health check failed", "error", err)
				}
			case <-sm.shutdownChan:
				return
			}
		}
	}()
}

// updateMetrics safely updates metrics with the given function
func (sm *SessionManagerImpl) updateMetrics(updateFunc func(*SessionMetrics)) {
	sm.metricsMutex.Lock()
	defer sm.metricsMutex.Unlock()
	updateFunc(&sm.metrics)
}

// generateSecureSessionID creates a secure, unique session identifier
func generateSecureSessionID() string {
	return "sess_" + uuid.New().String()
}

// generateThreadID creates a thread identifier for OpenAI
func generateThreadID() string {
	return "thread_" + uuid.New().String()
}

// Error definitions
var (
	ErrSessionNotFound          = fmt.Errorf("session not found")
	ErrSessionExpired           = fmt.Errorf("session expired")
	ErrInvalidSessionID         = fmt.Errorf("invalid session ID")
	ErrInvalidUserID            = fmt.Errorf("invalid user ID")
	ErrInvalidSession           = fmt.Errorf("invalid session")
	ErrNoActiveSession          = fmt.Errorf("no active session for user")
	ErrSessionLimitExceeded     = fmt.Errorf("session limit exceeded")
	ErrUserSessionLimitExceeded = fmt.Errorf("user session limit exceeded")
	ErrSessionNotOwnedByUser    = fmt.Errorf("session not owned by user")
)
