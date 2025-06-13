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

	// Register focus recommendation tool
	server.RegisterTool(mcp.Tool{
		Name:        "get-focus-recommendation",
		Description: "Get the single best task recommendation for current focus session",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"energy": map[string]interface{}{
					"type":        "string",
					"enum":        []string{"low", "medium", "high", "social"},
					"description": "Current energy level",
					"default":     "medium",
				},
				"mode": map[string]interface{}{
					"type":        "string",
					"enum":        []string{"deep", "quick", "admin"},
					"description": "Work mode preference",
					"default":     "deep",
				},
				"max_minutes": map[string]interface{}{
					"type":        "integer",
					"description": "Available time in minutes",
					"minimum":     5,
					"maximum":     480,
					"default":     60,
				},
			},
		},
	}, func(args map[string]interface{}) (interface{}, error) {
		return handleGetFocusRecommendation(vikunjaSvc, args)
	})

	return nil
}

// handleDailyFocus processes the daily-focus tool call
func handleDailyFocus(service *Service, args map[string]interface{}) (interface{}, error) {
	opts := models.FocusOptions{
		Energy:     "medium", // Default values
		Mode:       "deep",
		MaxMinutes: 300, // 5 hours default
		MaxTasks:   10,
	}

	// Parse arguments with defaults
	if energy, ok := args["energy"].(string); ok {
		opts.Energy = energy
	}
	if mode, ok := args["mode"].(string); ok {
		opts.Mode = mode
	}
	if hours, ok := args["hours"].(float64); ok {
		opts.MaxMinutes = int(hours * 60) // Convert hours to minutes
	}
	if maxItems, ok := args["max_items"].(float64); ok {
		opts.MaxTasks = int(maxItems)
	}

	log.Info("MCP daily-focus request",
		"energy", opts.Energy,
		"mode", opts.Mode,
		"max_minutes", opts.MaxMinutes,
		"max_tasks", opts.MaxTasks)

	tasks, err := service.GetFocusTasks(context.Background(), opts)
	if err != nil {
		log.Error("Failed to get focus tasks", "error", err)
		return nil, err
	}

	// Format response for Claude
	return map[string]interface{}{
		"message": "Focus tasks retrieved successfully",
		"summary": map[string]interface{}{
			"total_tasks":   len(tasks),
			"energy_filter": opts.Energy,
			"mode_filter":   opts.Mode,
			"target_hours":  float64(opts.MaxMinutes) / 60.0,
		},
		"tasks": tasks,
	}, nil
}

// handleGetTaskMetadata retrieves metadata for a specific task
func handleGetTaskMetadata(service *Service, args map[string]interface{}) (interface{}, error) {
	taskIDFloat, ok := args["task_id"].(float64)
	if !ok {
		return nil, fmt.Errorf("task_id must be a number")
	}
	taskID := int64(taskIDFloat)

	log.Info("MCP get-task-metadata request", "task_id", taskID)

	task, err := service.GetTaskMetadata(context.Background(), taskID)
	if err != nil {
		log.Error("Failed to get task metadata", "task_id", taskID, "error", err)
		return nil, err
	}

	if task == nil {
		return map[string]interface{}{
			"task_id":             taskID,
			"title":               "",
			"description":         "",
			"done":                false,
			"priority":            0,
			"has_hyperfocus_data": false,
			"metadata":            nil,
		}, nil
	}

	// Return enriched task data
	return map[string]interface{}{
		"task_id":             task.RawTask.ID,
		"title":               task.RawTask.Title,
		"description":         task.CleanDescription,
		"done":                task.RawTask.Done,
		"priority":            task.RawTask.Priority,
		"project_id":          task.RawTask.ProjectID,
		"has_hyperfocus_data": task.Metadata != nil,
		"metadata":            task.Metadata,
		"created":             task.RawTask.Created,
		"updated":             task.RawTask.Updated,
	}, nil
}

// handleGetFocusRecommendation gets the single best task for focus session
func handleGetFocusRecommendation(service *Service, args map[string]interface{}) (interface{}, error) {
	opts := models.FocusOptions{
		Energy:     "medium",
		Mode:       "deep",
		MaxMinutes: 60,
		MaxTasks:   1, // Single recommendation
	}

	// Parse arguments
	if energy, ok := args["energy"].(string); ok {
		opts.Energy = energy
	}
	if mode, ok := args["mode"].(string); ok {
		opts.Mode = mode
	}
	if maxMinutes, ok := args["max_minutes"].(float64); ok {
		opts.MaxMinutes = int(maxMinutes)
	}

	log.Info("MCP focus recommendation request",
		"energy", opts.Energy,
		"mode", opts.Mode,
		"max_minutes", opts.MaxMinutes)

	recommendation, err := service.GetTaskRecommendation(context.Background(), opts)
	if err != nil {
		log.Error("Failed to get task recommendation", "error", err)
		return nil, err
	}

	if recommendation == nil {
		return map[string]interface{}{
			"message":        "No suitable tasks found",
			"recommendation": nil,
			"criteria": map[string]interface{}{
				"energy":      opts.Energy,
				"mode":        opts.Mode,
				"max_minutes": opts.MaxMinutes,
			},
		}, nil
	}

	// Calculate session recommendation
	sessionLength := service.EstimateSessionLength(*recommendation, opts.MaxMinutes)

	return map[string]interface{}{
		"message": "Task recommendation generated successfully",
		"recommendation": map[string]interface{}{
			"task":               recommendation,
			"recommended_length": sessionLength,
			"can_extend":         recommendation.Metadata.Extend && opts.MaxMinutes >= recommendation.Metadata.Estimate,
			"focus_score":        recommendation.FocusScore,
			"reasoning":          fmt.Sprintf("Selected for %s energy, %s mode (score: %.1f)", opts.Energy, opts.Mode, recommendation.FocusScore),
		},
		"criteria": map[string]interface{}{
			"energy":      opts.Energy,
			"mode":        opts.Mode,
			"max_minutes": opts.MaxMinutes,
		},
	}, nil
}

// handleUpsertTask processes the upsert_task tool call
func handleUpsertTask(service *Service, args map[string]interface{}) (interface{}, error) {
	var task models.MinimalTask

	// Parse task data from arguments
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

	action := "updated"
	if task.TaskID == 0 {
		action = "created"
	}

	log.Info("MCP upsert task request",
		"action", action,
		"task_id", task.TaskID,
		"title", task.Title,
		"project_id", task.Project)

	result, err := service.UpsertTask(context.Background(), task)
	if err != nil {
		log.Error("Failed to upsert task", "error", err, "action", action)
		return nil, err
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
