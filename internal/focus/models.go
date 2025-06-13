package focus

import "context"

// handleDailyFocus processes the daily-focus tool call
// FocusService defines the interface for focus-related service methods
type FocusService interface {
	GetFocusTasks(ctx context.Context, opts FocusOptions) ([]FocusResult, error)
	parseHyperFocusMetadata(desc string) *HyperfocusMetadata
	cleanDescription(desc string) string
}

// FocusRequest represents the daily focus request payload.
// swagger:model FocusRequest
type FocusRequest struct {
	// Date is the desired focus date. Defaults to tomorrow if not provided.
	// example: "2025-05-26"
	Date string `json:"date,omitempty"`
	// Hours is the number of hours to focus. Derived from calendar if zero.
	// example: 5
	Hours float64 `json:"hours,omitempty"`
}

// FocusResponseItem represents an item in the daily-focus endpoint response.
// swagger:model FocusResponseItem
type FocusResponseItem struct {
	// T represents the task ID.
	// example: "task-123"
	T string `json:"t"`
	// P represents the project identifier.
	// example: "project-a"
	P string `json:"p"`
	// Est represents the estimated duration of the task.
	// example: 2.5
	Est float64 `json:"est"`
}

// FocusResponse represents the daily focus endpoint response comprising a slice of focus response items.
// swagger:model FocusResponse
type FocusResponse struct {
	// Items is a list of focus response items.
	Items []FocusResponseItem `json:"items"`
}

// APIError represents a standard API error response.
// swagger:model APIError
type APIError struct {
	// Code is the HTTP status code.
	// example: 400
	Code int `json:"code"`
	// Message is a human-readable error message.
	// example: "invalid JSON body"
	Message string `json:"message"`
}

type FocusItem struct {
	TaskID  string
	Project string
}

type FocusOptions struct {
	Energy   string
	Mode     string
	Hours    float32
	MaxItems int
	Date     string
}

type FocusResult struct {
	TaskID      string
	Project     string
	Metadata    *HyperfocusMetadata
	Priority    int
	Title       string
	Done        bool
	Description string
}

type HyperfocusMetadata struct {
	Mode string
}
