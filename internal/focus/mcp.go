package focus

import (
	"context"
	"fmt"

	"github.com/BelKirill/vikunja-mcp/models"
	"github.com/BelKirill/vikunja-mcp/pkg/mcp"
	"github.com/charmbracelet/log"
)

// RegisterMCPTools registers focus-related MCP tools
func RegisterMCPTools(server *mcp.Server) error {
	vikunjaSvc, err := NewService()
	if err != nil {
		return err
	}

	// Register the daily-focus tool
	server.RegisterTool(mcp.Tool{
		Name:        "daily-focus",
		Description: "Get AI-recommended tasks for focus session based on energy/mode",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"energy": map[string]interface{}{
					"type":        "string",
					"enum":        []string{"low", "medium", "high", "social"},
					"description": "Current energy level",
				},
				"mode": map[string]interface{}{
					"type":        "string",
					"enum":        []string{"deep", "quick", "admin"},
					"description": "Work mode preference",
				},
				"hours": map[string]interface{}{
					"type":        "integer",
					"description": "Target work hours (default: 5)",
					"minimum":     1,
					"maximum":     12,
				},
				"max_items": map[string]interface{}{
					"type":        "integer",
					"description": "Maximum tasks to return (default: 10)",
					"minimum":     1,
					"maximum":     50,
				},
			},
		},
	}, func(args map[string]interface{}) (interface{}, error) {
		return handleDailyFocus(vikunjaSvc, args)
	})

	// Register task metadata reader
	server.RegisterTool(mcp.Tool{
		Name:        "get-task-metadata",
		Description: "Extract hyperfocus metadata from a specific task",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"task_id": map[string]interface{}{
					"type":        "integer",
					"description": "Vikunja task ID",
				},
			},
			"required": []string{"task_id"},
		},
	}, func(args map[string]interface{}) (interface{}, error) {
		return handleGetTaskMetadata(vikunjaSvc, args)
	})

	// Register the upsert_task tool
	server.RegisterTool(mcp.Tool{
		Name:        "upsert_task",
		Description: "Create a new task or update an existing task (including marking complete)",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"task_id": map[string]interface{}{
					"type":        "integer",
					"description": "Task ID to update (omit to create new task)",
				},
				"title": map[string]interface{}{
					"type":        "string",
					"description": "Task title",
				},
				"done": map[string]interface{}{
					"type":        "boolean",
					"description": "Mark task as complete",
				},
				"priority": map[string]interface{}{
					"type":        "integer",
					"minimum":     1,
					"maximum":     5,
					"description": "Task priority (1-5)",
				},
				"description": map[string]interface{}{
					"type":        "string",
					"description": "Task description",
				},
				"project_id": map[string]interface{}{
					"type":        "integer",
					"description": "Project ID",
				},
			},
		},
	}, func(args map[string]interface{}) (interface{}, error) {
		return handleUpsertTask(vikunjaSvc, args)
	})

	return nil
}

// handleDailyFocus processes the daily-focus tool call
func handleDailyFocus(service models.FocusService, args map[string]interface{}) (interface{}, error) {
	opts := models.FocusOptions{}

	// Parse arguments with defaults
	if energy, ok := args["energy"].(string); ok {
		opts.Energy = energy
	}
	if mode, ok := args["mode"].(string); ok {
		opts.Mode = mode
	}
	if hours, ok := args["hours"].(float64); ok {
		opts.Hours = float32(hours)
	}
	if maxItems, ok := args["max_items"].(float64); ok {
		opts.MaxItems = int(maxItems)
	}

	log.Info("MCP daily-focus request", "energy", opts.Energy, "mode", opts.Mode, "hours", opts.Hours)

	tasks, err := service.GetFocusTasks(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	// Format response for Claude
	return map[string]interface{}{
		"message": "Focus tasks retrieved successfully",
		"summary": map[string]interface{}{
			"total_tasks":   len(tasks),
			"energy_filter": opts.Energy,
			"mode_filter":   opts.Mode,
			"target_hours":  opts.Hours,
		},
		"tasks": tasks,
	}, nil
}

// handleGetTaskMetadata retrieves metadata for a specific task
func handleGetTaskMetadata(service models.FocusService, args map[string]interface{}) (interface{}, error) {
	taskIDFloat, ok := args["task_id"].(float64)
	if !ok {
		return nil, fmt.Errorf("task_id must be a number")
	}
	taskID := int64(taskIDFloat)

	// For demonstration, just return a stub result
	metadata := service.ParseHyperFocusMetadata("")
	cleanDesc := service.CleanDescription("")

	return map[string]interface{}{
		"task_id":             taskID,
		"title":               "",
		"description":         cleanDesc,
		"priority":            0,
		"done":                false,
		"metadata":            metadata,
		"has_hyperfocus_data": metadata != nil,
	}, nil
}

// handleUpsertTask processes the upsert_task tool call
func handleUpsertTask(service models.FocusService, args map[string]interface{}) (interface{}, error) {
	var task models.MinimalTask
	if v, ok := args["task_id"].(float64); ok {
		task.TaskID = int64(v)
	}
	if v, ok := args["project_id"].(float64); ok {
		task.Project = int64(v)
	}
	if v, ok := args["title"].(string); ok {
		task.Title = v
	}
	if v, ok := args["description"].(string); ok {
		task.Description = v
	}
	if v, ok := args["priority"].(float64); ok {
		task.Priority = int(v)
	}
	if v, ok := args["done"].(bool); ok {
		task.Done = v
	}

	result, err := service.UpsertTask(context.Background(), task)
	if err != nil {
		return nil, err
	}

	action := "updated"
	if task.TaskID == 0 {
		action = "created"
	}

	return map[string]interface{}{
		"success": true,
		"action":  action,
		"task": map[string]interface{}{
			"task_id":     result.TaskID,
			"title":       result.Title,
			"done":        result.Done,
			"priority":    result.Priority,
			"description": result.Description,
			"project_id":  result.Project,
		},
		"message": fmt.Sprintf("Task %s successfully", action),
	}, nil
}
