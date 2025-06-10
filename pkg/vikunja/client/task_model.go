package client

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}

type FileInfo struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Mime    string `json:"mime"`
	Size    int    `json:"size"`
	Created string `json:"created"`
}

type Attachment struct {
	ID        int      `json:"id"`
	TaskID    int      `json:"task_id"`
	Created   string   `json:"created"`
	CreatedBy User     `json:"created_by"`
	File      FileInfo `json:"file"`
}

type Label struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	HexColor    string `json:"hex_color"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
	CreatedBy   User   `json:"created_by"`
}

type Reminder struct {
	Reminder       string `json:"reminder"`
	RelativePeriod int    `json:"relative_period"`
	RelativeTo     string `json:"relative_to"` // e.g., "due_date"
}

type Subscription struct {
	ID       int    `json:"id"`
	Entity   int    `json:"entity"`
	EntityID int    `json:"entity_id"`
	Created  string `json:"created"`
}

type RawTask struct {
	ID                     int                      `json:"id"`
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
	ProjectID              int                      `json:"project_id"`
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

// HyperfocusMetadata - our 4-field system
type HyperfocusMetadata struct {
	Energy  string `json:"energy"`  // "low"|"medium"|"high"|"social"
	Mode    string `json:"mode"`    // "deep"|"quick"|"admin"
	Extend  bool   `json:"extend"`  // Can extend beyond 25min?
	Minutes int    `json:"minutes"` // Base pomodoro estimate
}

// Task - RawTask + our enhancements
type Task struct {
	*RawTask
	Metadata         *HyperfocusMetadata `json:"metadata,omitempty"`
	CleanDescription string              `json:"clean_description"`
}
