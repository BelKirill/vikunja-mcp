package focus

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
