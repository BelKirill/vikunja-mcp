package openai

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/charmbracelet/log"
	openai "github.com/sashabaranov/go-openai"
)

// ThreadHealth represents the health status of a thread
type ThreadHealth struct {
	ThreadID         string        `json:"thread_id"`
	IsHealthy        bool          `json:"is_healthy"`
	LastChecked      time.Time     `json:"last_checked"`
	ResponseTime     time.Duration `json:"response_time"`
	ErrorCount       int           `json:"error_count"`
	LastError        string        `json:"last_error,omitempty"`
	RecoveryAttempts int           `json:"recovery_attempts"`
}

// RetryConfig defines retry behavior for health operations
type RetryConfig struct {
	MaxRetries    int           `json:"max_retries"`
	InitialDelay  time.Duration `json:"initial_delay"`
	MaxDelay      time.Duration `json:"max_delay"`
	BackoffFactor float64       `json:"backoff_factor"`
	EnableJitter  bool          `json:"enable_jitter"`
}

// ThreadError classifies different types of thread errors
type ThreadError struct {
	Type        string `json:"type"`        // "network", "auth", "not_found", "rate_limit", "unknown"
	Recoverable bool   `json:"recoverable"` // whether recovery should be attempted
	Message     string `json:"message"`
}

// ThreadManager handles OpenAI thread operations with health monitoring
type ThreadManager interface {
	// Basic thread operations
	CreateThread(ctx context.Context) (string, error)
	GetThread(ctx context.Context, threadID string) (*openai.Thread, error)
	DeleteThread(ctx context.Context, threadID string) error
	AddMessage(ctx context.Context, threadID, content string) error
	GetMessages(ctx context.Context, threadID string, limit int) ([]openai.Message, error)

	// Health monitoring and recovery
	HealthCheck(ctx context.Context, threadID string) (*ThreadHealth, error)
	ValidateThread(ctx context.Context, threadID string) error
	RecoverThread(ctx context.Context, sessionID, oldThreadID string) (string, error)
	StartHealthMonitoring(ctx context.Context, interval time.Duration, threadIDs []string)
	StopHealthMonitoring()
	GetHealthStatus(threadID string) (*ThreadHealth, error)
	ClassifyError(err error) *ThreadError
}

// ThreadManagerImpl implements ThreadManager with health monitoring
type ThreadManagerImpl struct {
	client        *openai.Client
	logger        *log.Logger
	retryConfig   RetryConfig
	healthStatus  map[string]*ThreadHealth
	healthMutex   sync.RWMutex
	monitorCancel context.CancelFunc
	monitorDone   chan struct{}
}

// NewThreadManager creates a new ThreadManager with health monitoring capabilities
func NewThreadManager(apiKey string) (ThreadManager, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("OpenAI API key is required")
	}

	client := openai.NewClient(apiKey)
	logger := log.New(os.Stderr)

	// Default retry configuration
	retryConfig := RetryConfig{
		MaxRetries:    3,
		InitialDelay:  1 * time.Second,
		MaxDelay:      30 * time.Second,
		BackoffFactor: 2.0,
		EnableJitter:  true,
	}

	return &ThreadManagerImpl{
		client:       client,
		logger:       logger,
		retryConfig:  retryConfig,
		healthStatus: make(map[string]*ThreadHealth),
		monitorDone:  make(chan struct{}),
	}, nil
}

// CreateThread creates a new empty thread
func (tm *ThreadManagerImpl) CreateThread(ctx context.Context) (string, error) {
	thread, err := tm.client.CreateThread(ctx, openai.ThreadRequest{})
	if err != nil {
		return "", fmt.Errorf("failed to create thread: %w", err)
	}

	tm.logger.Info("Thread created", "thread_id", thread.ID)

	// Initialize health status for new thread
	tm.healthMutex.Lock()
	tm.healthStatus[thread.ID] = &ThreadHealth{
		ThreadID:         thread.ID,
		IsHealthy:        true,
		LastChecked:      time.Now(),
		ErrorCount:       0,
		RecoveryAttempts: 0,
	}
	tm.healthMutex.Unlock()

	return thread.ID, nil
}

// GetThread retrieves thread information
func (tm *ThreadManagerImpl) GetThread(ctx context.Context, threadID string) (*openai.Thread, error) {
	if threadID == "" {
		return nil, fmt.Errorf("thread ID cannot be empty")
	}

	thread, err := tm.client.RetrieveThread(ctx, threadID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve thread %s: %w", threadID, err)
	}

	return &thread, nil
}

// DeleteThread deletes a thread
func (tm *ThreadManagerImpl) DeleteThread(ctx context.Context, threadID string) error {
	if threadID == "" {
		return fmt.Errorf("thread ID cannot be empty")
	}

	_, err := tm.client.DeleteThread(ctx, threadID)
	if err != nil {
		return fmt.Errorf("failed to delete thread %s: %w", threadID, err)
	}

	// Remove from health tracking
	tm.healthMutex.Lock()
	delete(tm.healthStatus, threadID)
	tm.healthMutex.Unlock()

	tm.logger.Info("Thread deleted", "thread_id", threadID)
	return nil
}

// AddMessage adds a user message to a thread
func (tm *ThreadManagerImpl) AddMessage(ctx context.Context, threadID, content string) error {
	if threadID == "" {
		return fmt.Errorf("thread ID cannot be empty")
	}
	if content == "" {
		return fmt.Errorf("message content cannot be empty")
	}

	_, err := tm.client.CreateMessage(ctx, threadID, openai.MessageRequest{
		Role:    "user",
		Content: content,
	})

	if err != nil {
		return fmt.Errorf("failed to add message to thread %s: %w", threadID, err)
	}

	tm.logger.Debug("Message added", "thread_id", threadID, "content_length", len(content))
	return nil
}

// GetMessages retrieves messages from a thread
func (tm *ThreadManagerImpl) GetMessages(ctx context.Context, threadID string, limit int) ([]openai.Message, error) {
	if threadID == "" {
		return nil, fmt.Errorf("thread ID cannot be empty")
	}

	if limit <= 0 {
		limit = 20
	}

	messages, err := tm.client.ListMessage(ctx, threadID, &limit, nil, nil, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get messages from thread %s: %w", threadID, err)
	}

	return messages.Messages, nil
}

// HealthCheck performs a comprehensive health check on a thread
func (tm *ThreadManagerImpl) HealthCheck(ctx context.Context, threadID string) (*ThreadHealth, error) {
	startTime := time.Now()

	health := &ThreadHealth{
		ThreadID:    threadID,
		LastChecked: startTime,
		IsHealthy:   false,
	}

	// Get existing health status to preserve error count and recovery attempts
	tm.healthMutex.RLock()
	if existing, exists := tm.healthStatus[threadID]; exists {
		health.ErrorCount = existing.ErrorCount
		health.RecoveryAttempts = existing.RecoveryAttempts
	}
	tm.healthMutex.RUnlock()

	// Validate thread exists and is accessible
	err := tm.ValidateThread(ctx, threadID)
	responseTime := time.Since(startTime)
	health.ResponseTime = responseTime

	if err != nil {
		threadErr := tm.ClassifyError(err)
		health.LastError = threadErr.Message
		health.ErrorCount++
		tm.logger.Warn("Thread health check failed",
			"thread_id", threadID,
			"error", err,
			"error_type", threadErr.Type,
			"recoverable", threadErr.Recoverable)
	} else {
		health.IsHealthy = true
		health.ErrorCount = 0
		health.LastError = ""
	}

	// Store health status
	tm.healthMutex.Lock()
	tm.healthStatus[threadID] = health
	tm.healthMutex.Unlock()

	return health, err
}

// ValidateThread checks if a thread is accessible via OpenAI API
func (tm *ThreadManagerImpl) ValidateThread(ctx context.Context, threadID string) error {
	if threadID == "" {
		return fmt.Errorf("thread ID cannot be empty")
	}

	// Try to retrieve thread - this is the most reliable validation
	_, err := tm.client.RetrieveThread(ctx, threadID)
	return err
}

// RecoverThread creates a new thread to replace a failed one
func (tm *ThreadManagerImpl) RecoverThread(ctx context.Context, sessionID, oldThreadID string) (string, error) {
	tm.logger.Info("Attempting thread recovery",
		"session_id", sessionID,
		"old_thread_id", oldThreadID)

	// Increment recovery attempts for the old thread
	tm.healthMutex.Lock()
	if health, exists := tm.healthStatus[oldThreadID]; exists {
		health.RecoveryAttempts++
	}
	tm.healthMutex.Unlock()

	// Create new thread with retry logic
	newThreadID, err := tm.createThreadWithRetry(ctx)
	if err != nil {
		tm.logger.Error("Thread recovery failed",
			"session_id", sessionID,
			"old_thread_id", oldThreadID,
			"error", err)
		return "", fmt.Errorf("failed to recover thread: %w", err)
	}

	// Initialize health status for new thread
	tm.healthMutex.Lock()
	tm.healthStatus[newThreadID] = &ThreadHealth{
		ThreadID:         newThreadID,
		IsHealthy:        true,
		LastChecked:      time.Now(),
		ErrorCount:       0,
		RecoveryAttempts: 0,
	}
	// Remove old thread from health tracking
	delete(tm.healthStatus, oldThreadID)
	tm.healthMutex.Unlock()

	tm.logger.Info("Thread recovery successful",
		"session_id", sessionID,
		"old_thread_id", oldThreadID,
		"new_thread_id", newThreadID)

	return newThreadID, nil
}

// createThreadWithRetry creates a thread with exponential backoff retry
func (tm *ThreadManagerImpl) createThreadWithRetry(ctx context.Context) (string, error) {
	var lastErr error

	for attempt := 0; attempt <= tm.retryConfig.MaxRetries; attempt++ {
		if attempt > 0 {
			delay := tm.calculateBackoffDelay(attempt)
			tm.logger.Debug("Retrying thread creation",
				"attempt", attempt,
				"delay", delay)

			select {
			case <-time.After(delay):
			case <-ctx.Done():
				return "", ctx.Err()
			}
		}

		threadID, err := tm.CreateThread(ctx)
		if err == nil {
			return threadID, nil
		}

		lastErr = err
		threadErr := tm.ClassifyError(err)

		// Don't retry non-recoverable errors
		if !threadErr.Recoverable {
			break
		}
	}

	return "", fmt.Errorf("thread creation failed after %d attempts: %w",
		tm.retryConfig.MaxRetries+1, lastErr)
}

// calculateBackoffDelay calculates exponential backoff delay with optional jitter
func (tm *ThreadManagerImpl) calculateBackoffDelay(attempt int) time.Duration {
	delay := time.Duration(float64(tm.retryConfig.InitialDelay) *
		math.Pow(tm.retryConfig.BackoffFactor, float64(attempt-1)))

	if delay > tm.retryConfig.MaxDelay {
		delay = tm.retryConfig.MaxDelay
	}

	// Add jitter if enabled
	if tm.retryConfig.EnableJitter {
		jitter := time.Duration(rand.Float64() * float64(delay) * 0.1)
		delay += jitter
	}

	return delay
}

// StartHealthMonitoring begins periodic health checks for specified threads
func (tm *ThreadManagerImpl) StartHealthMonitoring(ctx context.Context, interval time.Duration, threadIDs []string) {
	if tm.monitorCancel != nil {
		tm.StopHealthMonitoring()
	}

	monitorCtx, cancel := context.WithCancel(ctx)
	tm.monitorCancel = cancel

	go func() {
		defer close(tm.monitorDone)

		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				for _, threadID := range threadIDs {
					// Health check each thread (don't block on failures)
					go func(tid string) {
						_, err := tm.HealthCheck(monitorCtx, tid)
						if err != nil {
							tm.logger.Debug("Background health check failed",
								"thread_id", tid, "error", err)
						}
					}(threadID)
				}
			case <-monitorCtx.Done():
				tm.logger.Info("Health monitoring stopped")
				return
			}
		}
	}()

	tm.logger.Info("Health monitoring started",
		"interval", interval,
		"thread_count", len(threadIDs))
}

// StopHealthMonitoring stops the background health monitoring
func (tm *ThreadManagerImpl) StopHealthMonitoring() {
	if tm.monitorCancel != nil {
		tm.monitorCancel()
		<-tm.monitorDone
		tm.monitorCancel = nil
	}
}

// GetHealthStatus retrieves the current health status for a thread
func (tm *ThreadManagerImpl) GetHealthStatus(threadID string) (*ThreadHealth, error) {
	tm.healthMutex.RLock()
	defer tm.healthMutex.RUnlock()

	health, exists := tm.healthStatus[threadID]
	if !exists {
		return nil, fmt.Errorf("no health status found for thread %s", threadID)
	}

	// Return a copy to avoid race conditions
	healthCopy := *health
	return &healthCopy, nil
}

// ClassifyError categorizes errors for recovery decisions
func (tm *ThreadManagerImpl) ClassifyError(err error) *ThreadError {
	if err == nil {
		return nil
	}

	errStr := err.Error()

	// Network/timeout errors - usually recoverable
	if strings.Contains(errStr, "timeout") ||
		strings.Contains(errStr, "connection") ||
		strings.Contains(errStr, "network") {
		return &ThreadError{
			Type:        "network",
			Recoverable: true,
			Message:     errStr,
		}
	}

	// Authentication errors - not recoverable without fixing API key
	if strings.Contains(errStr, "unauthorized") ||
		strings.Contains(errStr, "authentication") ||
		strings.Contains(errStr, "api key") {
		return &ThreadError{
			Type:        "auth",
			Recoverable: false,
			Message:     errStr,
		}
	}

	// Not found errors - thread deleted/invalid, recoverable by creating new thread
	if strings.Contains(errStr, "not found") ||
		strings.Contains(errStr, "404") {
		return &ThreadError{
			Type:        "not_found",
			Recoverable: true,
			Message:     errStr,
		}
	}

	// Rate limiting - recoverable with backoff
	if strings.Contains(errStr, "rate limit") ||
		strings.Contains(errStr, "429") {
		return &ThreadError{
			Type:        "rate_limit",
			Recoverable: true,
			Message:     errStr,
		}
	}

	// Default: unknown error, attempt recovery
	return &ThreadError{
		Type:        "unknown",
		Recoverable: true,
		Message:     errStr,
	}
}
