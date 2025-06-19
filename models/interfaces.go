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
	ParseHyperFocusMetadata(desc string) (*HyperFocusMetadata, error)
	CleanDescription(desc string) string
	UpsertTask(ctx context.Context, task Task) (*Task, error)
}

// DecisionEngine interface allows pluggable AI decision backends
type DecisionEngine interface {
	RankTasks(ctx context.Context, request *DecisionRequest) (*DecisionResponse, error)
	SuggestFilter(ctx context.Context, request *string) (*FilterSuggestionResponse, error)
}

// ContextualFilter applies business rules before AI ranking
type ContextualFilter interface {
	Filter(ctx context.Context, tasks []Task, opts *FocusOptions) []Task
	Name() string
}

// FallbackStrategy handles cases when AI is unavailable
type FallbackStrategy interface {
	SelectTasks(ctx context.Context, tasks []Task, opts *FocusOptions) []RankedTask
	GetRecommendation(ctx context.Context, tasks []Task, opts *FocusOptions) *TaskRecommendation
}
