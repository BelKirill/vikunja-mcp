package models

import (
	"context"
)

// handleDailyFocus processes the daily-focus tool call
// FocusService defines the interface for focus-related service methods
// Use MinimalTask as the DTO for upsert operations
// Export ParseHyperFocusMetadata and CleanDescription for idiomatic Go
type FocusService interface {
	GetFocusTasks(ctx context.Context, opts FocusOptions) ([]FocusResult, error)
	ParseHyperFocusMetadata(desc string) *HyperfocusMetadata
	CleanDescription(desc string) string
	UpsertTask(ctx context.Context, task MinimalTask) (*MinimalTask, error)
}
