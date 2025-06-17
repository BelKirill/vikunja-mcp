package openai

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestNewThreadManager(t *testing.T) {
	tests := []struct {
		name    string
		apiKey  string
		wantErr bool
	}{
		{
			name:    "valid api key",
			apiKey:  "test-key",
			wantErr: false,
		},
		{
			name:    "empty api key",
			apiKey:  "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm, err := NewThreadManager(tt.apiKey)

			if tt.wantErr {
				if err == nil {
					t.Error("NewThreadManager() expected error but got none")
				}
				if tm != nil {
					t.Error("NewThreadManager() expected nil ThreadManager on error")
				}
			} else {
				if err != nil {
					t.Errorf("NewThreadManager() unexpected error: %v", err)
				}
				if tm == nil {
					t.Error("NewThreadManager() returned nil ThreadManager")
				}
			}
		})
	}
}

func TestThreadManagerBasicOperations(t *testing.T) {
	// Skip this test if no OpenAI API key is available
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		t.Skip("Skipping integration test: OPENAI_API_KEY not set")
	}

	tm, err := NewThreadManager(apiKey)
	if err != nil {
		t.Fatalf("Failed to create ThreadManager: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test thread creation
	threadID, err := tm.CreateThread(ctx)
	if err != nil {
		t.Fatalf("Failed to create thread: %v", err)
	}

	if threadID == "" {
		t.Fatal("CreateThread returned empty thread ID")
	}

	t.Logf("Created thread: %s", threadID)

	// Test thread retrieval
	thread, err := tm.GetThread(ctx, threadID)
	if err != nil {
		t.Fatalf("Failed to get thread: %v", err)
	}

	if thread.ID != threadID {
		t.Errorf("Expected thread ID %s, got %s", threadID, thread.ID)
	}

	// Test adding a message
	err = tm.AddMessage(ctx, threadID, "Hello, this is a test message!")
	if err != nil {
		t.Fatalf("Failed to add message: %v", err)
	}

	// Test getting messages
	messages, err := tm.GetMessages(ctx, threadID, 10)
	if err != nil {
		t.Fatalf("Failed to get messages: %v", err)
	}

	if len(messages) == 0 {
		t.Error("Expected at least one message, got none")
	}

	// Test thread deletion
	err = tm.DeleteThread(ctx, threadID)
	if err != nil {
		t.Fatalf("Failed to delete thread: %v", err)
	}

	t.Logf("Successfully completed all basic operations for thread %s", threadID)
}

func TestThreadManagerErrorHandling(t *testing.T) {
	// Create ThreadManager with fake API key to test error handling
	tm, err := NewThreadManager("fake-key")
	if err != nil {
		t.Fatalf("Failed to create ThreadManager: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Test operations with invalid API key (should fail)
	_, err = tm.CreateThread(ctx)
	if err == nil {
		t.Error("Expected error with fake API key, but got none")
	}

	// Test operations with empty thread ID
	_, err = tm.GetThread(ctx, "")
	if err == nil {
		t.Error("Expected error with empty thread ID, but got none")
	}

	err = tm.DeleteThread(ctx, "")
	if err == nil {
		t.Error("Expected error with empty thread ID, but got none")
	}

	err = tm.AddMessage(ctx, "", "test")
	if err == nil {
		t.Error("Expected error with empty thread ID, but got none")
	}

	err = tm.AddMessage(ctx, "thread-123", "")
	if err == nil {
		t.Error("Expected error with empty message content, but got none")
	}

	_, err = tm.GetMessages(ctx, "", 10)
	if err == nil {
		t.Error("Expected error with empty thread ID, but got none")
	}
}
