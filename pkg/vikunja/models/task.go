package models

import "time"

// RawTask represents the actual JSON structure from Vikunja API
type RawTask struct {
	// Core fields
	ID          string `json:"id"`
	Title       string `json:"title"`       // Almost certainly missing from your current struct
	Description string `json:"description"` // Almost certainly missing from your current struct

	// Your existing fields
	ProjectID   string `json:"project_id"`
	DueDate     string `json:"due_date"` // ISO 8601
	Priority    int    `json:"priority"`
	LabelIDs    []int  `json:"label_ids"`
	EstimateMin int    `json:"estimate"` // minutes

	// Standard task fields that likely exist
	Done      bool      `json:"done,omitempty"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`

	// Possible additional fields (test and remove if not present)
	BucketID    string  `json:"bucket_id,omitempty"`
	Position    float64 `json:"position,omitempty"`
	PercentDone float64 `json:"percent_done,omitempty"`
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
