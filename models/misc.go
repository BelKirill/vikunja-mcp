package models

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
	Entity   string `json:"entity"`
	EntityID int    `json:"entity_id"`
	Created  string `json:"created"`
}
