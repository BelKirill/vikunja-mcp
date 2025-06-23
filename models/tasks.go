package models

// FocusOptions represents the criteria for selecting focus tasks
type FocusOptions struct {
	Energy          string  `json:"energy"`           // "low"|"medium"|"high"|"social"
	Mode            string  `json:"mode"`             // "deep"|"quick"|"admin"
	MaxMinutes      int     `json:"max_minutes"`      // Maximum available time for session
	MaxTasks        int     `json:"max_tasks"`        // Maximum number of tasks to return
	Date            string  `json:"date"`             // Target date for focus session
	HyperFocus      int     `json:"hyperfocus_level"` // Target hyperfocus level
	Instructions    string  `json:"instructions"`     // Free text instructions for selecting the tasks
	ExcludeProjects []int64 `json:"exclude_projects,omitempty"`
	OnlyProjects    []int64 `json:"only_projects,omitempty"`
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

type Dependencies struct {
	Blocked_by []int `json:"blocked_by"`
	Blocks     []int `json:"blocks"`
}

type ContextualHints struct {
	PartOfChain bool   `json:"is_part_of_chain"`  // true,
	Next        []int  `json:"next_in_chain"`     // [9],
	Progress    string `json:"chain_progress"`    // "1/2",
	Name        string `json:"chain_name"`        // "Quality Assurance",
	Description string `json:"chain_description"` // "Testing and validation for dependency system reliability"
}

// HyperFocusMetadata represents the ADHD-optimized task metadata
type HyperFocusMetadata struct {
	Energy                  string          `json:"energy"`           // "low"|"medium"|"high"|"social"
	Mode                    string          `json:"mode"`             // "deep"|"quick"|"admin"
	Extend                  bool            `json:"extend"`           // Can extend beyond 25min?
	Minutes                 int             `json:"minutes"`          // Base pomodoro work unit
	Estimate                int             `json:"estimate"`         // Total estimated duration in minutes
	HyperFocusCompatability int             `json:"hyper_focus_comp"` // Scale 1-5, hyperfocus compatibility
	Instructions            string          `json:"instructions"`     // Instructions and considerations on selecting this task
	Dependencies            Dependencies    `json:"dependencies"`     // Other tasks blocking or being blocked by this task
	ContextualHints         ContextualHints `json:"contextual_hints"` // Currently: Chain tracking for tasks
}

// FullRawTask represents the complete task data from Vikunja API
type FullRawTask struct {
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

// RawTask represents the essential task data from Vikunja API (optimized for cost reduction)
type RawTask struct {
	ID          int64          `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Done        bool           `json:"done"`
	HexColor    string         `json:"hex_color"`
	Identifier  string         `json:"identifier"`
	Priority    int            `json:"priority"`
	ProjectID   int64          `json:"project_id"`
	Created     string         `json:"created"`
	Updated     string         `json:"updated"`
	Labels      []PartialLabel `json:"labels"`
}

// Task represents an enriched task with parsed hyperfocus metadata
type Task struct {
	Identifier       string              `json:"identifier"`
	RawTask          *RawTask            `json:"raw_task"`
	Metadata         *HyperFocusMetadata `json:"metadata,omitempty"`
	Comments         []Comment           `json:"comments,omitempty"`
	FocusScore       float64             `json:"focus_score,omitempty"`
	CleanDescription string              `json:"clean_description"`
}

// SessionRecommendation represents an optimal focus session recommendation
type SessionRecommendation struct {
	Task              *Task   `json:"task"`
	Score             float64 `json:"recommendation_score"`
	RecommendedLength int     `json:"recommended_length"` // Minutes
	CanExtend         bool    `json:"can_extend"`
	Reasoning         string  `json:"reasoning"` // Why this task was chosen
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
