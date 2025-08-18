package models

import "time"

// ProjectExport represents a complete project export with tasks and metadata
type ProjectExport struct {
	Project    *PartialProject     `json:"project"`
	Tasks      []Task              `json:"tasks"`
	Labels     []PartialLabel      `json:"labels"`
	Comments   map[int64][]Comment `json:"comments"` // task_id -> comments
	ExportMeta ExportMetadata      `json:"export_meta"`
}

// ExportMetadata contains information about the export operation
type ExportMetadata struct {
	ExportedAt    time.Time `json:"exported_at"`
	TaskCount     int       `json:"task_count"`
	IncludesDone  bool      `json:"includes_done"`
	FilterUsed    string    `json:"filter_used,omitempty"`
	ExportVersion string    `json:"export_version"`
	ProjectTitle  string    `json:"project_title"`
}

// ExportOptions configures how the export should be performed
type ExportOptions struct {
	IncludeCompleted bool   `json:"include_completed"`
	IncludeComments  bool   `json:"include_comments"`
	IncludeMetadata  bool   `json:"include_metadata"` // AI-generated hyperfocus metadata
	CustomFilter     string `json:"custom_filter,omitempty"`
	OutputPath       string `json:"output_path"`
	PrettyPrint      bool   `json:"pretty_print"`
}

// MultiProjectExport represents an export containing multiple projects
type MultiProjectExport struct {
	Projects   []ProjectExport `json:"projects"`
	ExportMeta ExportMetadata  `json:"export_meta"`
}

// ExportResult represents the result of an export operation
type ExportResult struct {
	Success      bool   `json:"success"`
	FilePath     string `json:"file_path"`
	TaskCount    int    `json:"task_count"`
	ProjectCount int    `json:"project_count"`
	ErrorMessage string `json:"error_message,omitempty"`
	FileSize     int64  `json:"file_size,omitempty"`
}
