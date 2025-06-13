package models

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

// FocusOptions represents the criteria for selecting focus tasks
type FocusOptions struct {
	Energy     string `json:"energy"`      // "low"|"medium"|"high"|"social"
	Mode       string `json:"mode"`        // "deep"|"quick"|"admin"
	MaxMinutes int    `json:"max_minutes"` // Maximum available time for session
	MaxTasks   int    `json:"max_tasks"`   // Maximum number of tasks to return
	Date       string `json:"date"`        // Target date for focus session
}

// FocusResult represents a task recommended for a focus session with enriched metadata
type FocusResult struct {
	TaskID      int64               `json:"task_id"`
	Project     int64               `json:"project"`
	Metadata    *HyperFocusMetadata `json:"metadata"`
	Priority    int                 `json:"priority"`
	Title       string              `json:"title"`
	Done        bool                `json:"done"`
	Description string              `json:"description"`
	Estimate    int                 `json:"estimate"`    // Estimated duration in minutes
	FocusScore  float64             `json:"focus_score"` // Calculated suitability score
}

// HyperFocusMetadata represents the ADHD-optimized task metadata
type HyperFocusMetadata struct {
	Energy                  string `json:"energy"`           // "low"|"medium"|"high"|"social"
	Mode                    string `json:"mode"`             // "deep"|"quick"|"admin"
	Extend                  bool   `json:"extend"`           // Can extend beyond 25min?
	Minutes                 int    `json:"minutes"`          // Base pomodoro work unit
	Estimate                int    `json:"estimate"`         // Total estimated duration in minutes
	HyperFocusCompatability int    `json:"hyper_focus_comp"` // Scale 1-5, hyperfocus compatibility
}

// MinimalTask represents the essential task data for MCP operations
type MinimalTask struct {
	TaskID      int64               `json:"task_id"`
	Project     int64               `json:"project"`
	Metadata    *HyperFocusMetadata `json:"metadata"`
	Priority    int                 `json:"priority"`
	Title       string              `json:"title"`
	Done        bool                `json:"done"`
	Description string              `json:"description"`
}

// User represents a Vikunja user
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}

// FileInfo represents file metadata for attachments
type FileInfo struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Mime    string `json:"mime"`
	Size    int    `json:"size"`
	Created string `json:"created"`
}

// Attachment represents a file attached to a task
type Attachment struct {
	ID        int      `json:"id"`
	TaskID    int      `json:"task_id"`
	Created   string   `json:"created"`
	CreatedBy User     `json:"created_by"`
	File      FileInfo `json:"file"`
}

// Label represents a task label/tag
type Label struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	HexColor    string `json:"hex_color"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
	CreatedBy   User   `json:"created_by"`
}

// Reminder represents a task reminder configuration
type Reminder struct {
	Reminder       string `json:"reminder"`
	RelativePeriod int    `json:"relative_period"`
	RelativeTo     string `json:"relative_to"` // e.g., "due_date"
}

// Subscription represents a user's subscription to task updates
type Subscription struct {
	ID       int    `json:"id"`
	Entity   int    `json:"entity"`
	EntityID int    `json:"entity_id"`
	Created  string `json:"created"`
}

// RawTask represents the complete task data from Vikunja API
type RawTask struct {
	ID                     int64                    `json:"id"`
	Title                  string                   `json:"title"`
	Description            string                   `json:"description"`
	Done                   bool                     `json:"done"`
	DoneAt                 string                   `json:"done_at"`
	DueDate                string                   `json:"due_date"`
	StartDate              string                   `json:"start_date"`
	EndDate                string                   `json:"end_date"`
	HexColor               string                   `json:"hex_color"`
	Identifier             string                   `json:"identifier"`
	IsFavorite             bool                     `json:"is_favorite"`
	Index                  int                      `json:"index"`
	Position               int                      `json:"position"`
	Priority               int                      `json:"priority"`
	PercentDone            int                      `json:"percent_done"`
	RepeatAfter            int                      `json:"repeat_after"`
	RepeatMode             int                      `json:"repeat_mode"`
	BucketID               int                      `json:"bucket_id"`
	ProjectID              int64                    `json:"project_id"`
	CoverImageAttachmentID int                      `json:"cover_image_attachment_id"`
	Created                string                   `json:"created"`
	Updated                string                   `json:"updated"`
	CreatedBy              User                     `json:"created_by"`
	Assignees              []User                   `json:"assignees"`
	Attachments            []Attachment             `json:"attachments"`
	Labels                 []Label                  `json:"labels"`
	Reminders              []Reminder               `json:"reminders"`
	Reactions              map[string][]User        `json:"reactions"`
	RelatedTasks           map[string][]interface{} `json:"related_tasks"`
	Subscription           Subscription             `json:"subscription"`
}

// Task represents an enriched task with parsed hyperfocus metadata
type Task struct {
	RawTask          *RawTask            `json:"raw_task"`
	Metadata         *HyperFocusMetadata `json:"metadata,omitempty"`
	CleanDescription string              `json:"clean_description"`
}

// SessionRecommendation represents an optimal focus session recommendation
type SessionRecommendation struct {
	Task              *FocusResult `json:"task"`
	RecommendedLength int          `json:"recommended_length"` // Minutes
	CanExtend         bool         `json:"can_extend"`
	Reasoning         string       `json:"reasoning"` // Why this task was chosen
}

// FocusSession represents an active or completed focus session
type FocusSession struct {
	ID            string `json:"id"`
	TaskID        int64  `json:"task_id"`
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time,omitempty"`
	PlannedLength int    `json:"planned_length"` // Minutes
	ActualLength  int    `json:"actual_length"`  // Minutes
	Completed     bool   `json:"completed"`
	Extended      bool   `json:"extended"`
	Notes         string `json:"notes,omitempty"`
	Energy        string `json:"energy"`        // User's energy at start
	Mode          string `json:"mode"`          // Chosen mode
	Effectiveness int    `json:"effectiveness"` // 1-5 self-rating
}

// EnergyPattern represents learned patterns about user's energy levels
type EnergyPattern struct {
	TimeOfDay     string  `json:"time_of_day"`    // "morning"|"afternoon"|"evening"
	DayOfWeek     string  `json:"day_of_week"`    // "monday"|"tuesday"|etc
	TypicalEnergy string  `json:"typical_energy"` // "low"|"medium"|"high"|"social"
	Confidence    float64 `json:"confidence"`     // 0.0-1.0 confidence in pattern
	SampleSize    int     `json:"sample_size"`    // Number of sessions contributing to pattern
}

// TaskPerformance represents how well a task performs in focus sessions
type TaskPerformance struct {
	TaskID           int64   `json:"task_id"`
	SessionCount     int     `json:"session_count"`
	CompletionRate   float64 `json:"completion_rate"`    // 0.0-1.0
	AverageEffective float64 `json:"average_effective"`  // Average effectiveness rating
	PreferredEnergy  string  `json:"preferred_energy"`   // Energy level that works best
	PreferredMode    string  `json:"preferred_mode"`     // Mode that works best
	AvgSessionLength int     `json:"avg_session_length"` // Average successful session length
}
