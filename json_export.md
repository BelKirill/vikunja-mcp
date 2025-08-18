# JSON Export Design Document

This document outlines the design and implementation plan for adding local JSON export functionality to the vikunja-mcp system, allowing users to download project tasks and metadata to local files.

## Current Project-Based Task Extraction

### Existing Functionality

The system already has several components for project-based task operations:

1. **Project Filtering in Focus Sessions** (`internal/focus/service.go:111-112`):
   ```go
   rawTasksFiltered := filterTasksByProjects(rawTasks, opts.ExcludeProjects, opts.OnlyProjects)
   ```

2. **MCP Tools with Project Support**:
   - `daily-focus`: Has `only_projects` and `exclude_projects` parameters
   - `get-filtered-tasks`: Has `project_id` parameter for single project filtering
   - `upsert-task`: Creates/updates tasks with `project_id`

3. **Current Data Retrieval Methods**:
   - `GetAllTasks()`: Fetches all user tasks (paginated)
   - `GetFilteredTasks(filter)`: Uses Vikunja filter expressions like `project_id = 8`
   - `GetFullTaskData(taskID)`: Gets complete task data including project info

4. **Project Data Available**:
   - `GetAllProjects()`: Returns all user projects
   - `GetProject(projectID)`: Gets specific project details

## JSON Export Architecture

### Data Structures

**Primary Export Structure**:
```go
type ProjectExport struct {
    Project      models.PartialProject      `json:"project"`
    Tasks        []models.Task             `json:"tasks"`
    Labels       []models.PartialLabel     `json:"labels"`
    Comments     map[int64][]models.Comment `json:"comments"` // task_id -> comments
    ExportMeta   ExportMetadata            `json:"export_meta"`
}

type ExportMetadata struct {
    ExportedAt    time.Time `json:"exported_at"`
    TaskCount     int       `json:"task_count"`
    IncludesDone  bool      `json:"includes_done"`
    FilterUsed    string    `json:"filter_used,omitempty"`
    ExportVersion string    `json:"export_version"`
    ProjectTitle  string    `json:"project_title"`
}

type ExportOptions struct {
    IncludeCompleted bool   `json:"include_completed"`
    IncludeComments  bool   `json:"include_comments"`
    IncludeMetadata  bool   `json:"include_metadata"`  // AI-generated hyperfocus metadata
    CustomFilter     string `json:"custom_filter,omitempty"`
    OutputPath       string `json:"output_path"`
    PrettyPrint      bool   `json:"pretty_print"`
}
```

### Core Export Service

**Location**: `internal/focus/export.go` (new file)

**Main Export Function**:
```go
func (s *Service) ExportProjectToJSON(ctx context.Context, projectID int64, options ExportOptions) (*ProjectExport, error) {
    // 1. Fetch project metadata
    // 2. Fetch and filter tasks
    // 3. Enrich tasks with AI metadata if requested
    // 4. Fetch comments if requested
    // 5. Collect unique labels
    // 6. Build export structure
    // 7. Write to file
}
```

**Supporting Functions**:
```go
func (s *Service) ExportFilteredTasksToJSON(ctx context.Context, filter string, options ExportOptions) (*ProjectExport, error)
func (s *Service) ExportMultipleProjectsToJSON(ctx context.Context, projectIDs []int64, options ExportOptions) ([]ProjectExport, error)
func writeExportToFile(export *ProjectExport, outputPath string, prettyPrint bool) error
```

## MCP Tool Integration

### New MCP Tool: `export-project-json`

**Registration** (`internal/focus/mcp_tools.go`):
```go
server.RegisterTool(mcp.Tool{
    Name:        "export-project-json",
    Description: "Export all tasks from a project to local JSON file with enriched metadata",
    InputSchema: map[string]interface{}{
        "type": "object",
        "properties": map[string]interface{}{
            "project_id": map[string]interface{}{
                "type":        "integer",
                "description": "Project ID to export (omit to export all projects)",
            },
            "output_path": map[string]interface{}{
                "type":        "string",
                "description": "Local file path for JSON export (e.g., './export.json')",
            },
            "include_completed": map[string]interface{}{
                "type":        "boolean",
                "description": "Include completed tasks in export (default: false)",
                "default":     false,
            },
            "include_comments": map[string]interface{}{
                "type":        "boolean", 
                "description": "Include task comments in export (default: false)",
                "default":     false,
            },
            "include_metadata": map[string]interface{}{
                "type":        "boolean",
                "description": "Include AI-generated hyperfocus metadata (default: true)",
                "default":     true,
            },
            "custom_filter": map[string]interface{}{
                "type":        "string",
                "description": "Vikunja filter expression to limit export scope",
            },
            "pretty_print": map[string]interface{}{
                "type":        "boolean",
                "description": "Format JSON with indentation for readability (default: true)",
                "default":     true,
            },
        },
        "required": ["output_path"],
    },
})
```

**Handler Function** (`internal/focus/mcp_handlers.go`):
```go
func handleExportProjectJSON(service *Service, args map[string]interface{}) (interface{}, error) {
    // Parse arguments
    // Validate output path
    // Call export service
    // Return success/failure response
}
```

## Implementation Plan

### Phase 1: Core Export Service
**Files to create/modify**:
- `internal/focus/export.go` (new)
- `models/export.go` (new)

**Functionality**:
1. Basic project export with tasks and metadata
2. File writing with JSON marshaling
3. Error handling and logging
4. Basic filtering support

### Phase 2: MCP Tool Integration
**Files to modify**:
- `internal/focus/mcp_tools.go` (add tool registration)
- `internal/focus/mcp_handlers.go` (add handler function)

**Functionality**:
1. MCP tool registration
2. Argument parsing and validation
3. Integration with core export service
4. Response formatting

### Phase 3: Advanced Features
**Additional functionality**:
1. **Batch Export**: Multiple projects in single operation
2. **Incremental Export**: Only tasks modified since timestamp
3. **Format Options**: CSV export, compressed JSON
4. **Validation**: JSON schema validation for exports
5. **Resume Capability**: Handle large exports with progress tracking

### Phase 4: Enhanced Data Collection
**Extended data inclusion**:
1. **Task Dependencies**: Blocked/blocking relationships
2. **File Attachments**: Metadata for attached files
3. **Task History**: Creation/modification timeline
4. **User Assignments**: Full user details for assigned tasks

## Usage Examples

### Basic Project Export
```bash
export-project-json --project_id=1 --output_path="./project1_export.json"
```

### Full Export with Comments
```bash
export-project-json --project_id=1 --output_path="./project1_full.json" --include_completed=true --include_comments=true
```

### Filtered Export
```bash
export-project-json --custom_filter="project_id = 1 && priority >= 3" --output_path="./high_priority_tasks.json"
```

### All Projects Export
```bash
export-project-json --output_path="./all_projects.json" --include_completed=false
```

## File Locations Summary

### New Files
- `internal/focus/export.go` - Core export service
- `models/export.go` - Export-related data models

### Modified Files
- `internal/focus/mcp_tools.go` - Tool registration (around line 200)
- `internal/focus/mcp_handlers.go` - Handler functions (around line 280)

## Expected Output Format

```json
{
  "project": {
    "id": 1,
    "title": "Web Development",
    "description": "Frontend and backend tasks"
  },
  "tasks": [
    {
      "task_id": 123,
      "title": "Implement authentication",
      "description": "OAuth2 integration",
      "priority": 4,
      "done": false,
      "metadata": {
        "energy": "medium",
        "mode": "deep",
        "hyper_focus_comp": 4,
        "minutes": 60,
        "extend": true
      },
      "labels": ["technical", "backend"]
    }
  ],
  "labels": [
    {"id": 1, "title": "technical", "description": "Technical work", "hex_color": "3498db"},
    {"id": 2, "title": "backend", "description": "Server-side code", "hex_color": "e74c3c"}
  ],
  "comments": {
    "123": [
      {
        "id": 456,
        "comment": "Started OAuth implementation",
        "author": {"name": "John Doe", "email": "john@example.com"},
        "created": "2024-01-15T10:30:00Z"
      }
    ]
  },
  "export_meta": {
    "exported_at": "2024-01-15T15:45:00Z",
    "task_count": 1,
    "includes_done": false,
    "export_version": "1.0",
    "project_title": "Web Development"
  }
}
```

## Benefits

1. **Data Backup**: Complete project data preservation
2. **External Integration**: Easy import into other tools
3. **Analysis**: Data mining and reporting capabilities  
4. **Migration**: Project transfer between Vikunja instances
5. **ADHD Optimization**: Preserves AI-generated hyperfocus metadata
6. **Flexibility**: Multiple export formats and filtering options

This design leverages the existing project filtering infrastructure while adding comprehensive JSON export capabilities suitable for backup, analysis, and integration with external productivity tools.