package openai

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	openai "github.com/sashabaranov/go-openai"
)

// ThreadManager handles basic OpenAI thread operations
type ThreadManager interface {
	CreateThread(ctx context.Context) (string, error)
	GetThread(ctx context.Context, threadID string) (*openai.Thread, error)
	DeleteThread(ctx context.Context, threadID string) error
	AddMessage(ctx context.Context, threadID, content string) error
	GetMessages(ctx context.Context, threadID string, limit int) ([]openai.Message, error)
}

// ThreadManagerImpl implements basic thread management
type ThreadManagerImpl struct {
	client *openai.Client
	logger *log.Logger
}

// NewThreadManager creates a basic ThreadManager
func NewThreadManager(apiKey string) (ThreadManager, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("OpenAI API key is required")
	}

	client := openai.NewClient(apiKey)
	logger := log.New(os.Stderr)

	return &ThreadManagerImpl{
		client: client,
		logger: logger,
	}, nil
}

// CreateThread creates a new empty thread
func (tm *ThreadManagerImpl) CreateThread(ctx context.Context) (string, error) {
	thread, err := tm.client.CreateThread(ctx, openai.ThreadRequest{})
	if err != nil {
		return "", fmt.Errorf("failed to create thread: %w", err)
	}

	tm.logger.Info("Thread created", "thread_id", thread.ID)
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

	// Based on the test code, ListMessage signature is:
	// ListMessage(ctx, threadID, limit, order, after, before, runID)
	messages, err := tm.client.ListMessage(ctx, threadID, &limit, nil, nil, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get messages from thread %s: %w", threadID, err)
	}

	return messages.Messages, nil
}
