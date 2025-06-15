package focus

import (
	"context"
	"fmt"

	"github.com/BelKirill/vikunja-mcp/models"
	"github.com/charmbracelet/log"
)

// handleDailyFocus processes the daily-focus tool call
func handleDailyFocus(service *Service, args map[string]interface{}) (interface{}, error) {
	log.Debug("handleDailyFocus called", "args", args)
	opts := models.FocusOptions{
		Energy:       "medium", // Default values
		Mode:         "deep",
		MaxMinutes:   300, // 5 hours default
		MaxTasks:     10,
		Instructions: "General request, give a good assortment of tasks",
	}

	// Parse arguments with defaults
	if energy, ok := args["energy"].(string); ok {
		log.Debug("Parsed energy from args", "energy", energy)
		opts.Energy = energy
	}
	if mode, ok := args["mode"].(string); ok {
		log.Debug("Parsed mode from args", "mode", mode)
		opts.Mode = mode
	}
	if hours, ok := args["hours"].(float64); ok {
		log.Debug("Parsed hours from args", "hours", hours)
		opts.MaxMinutes = int(hours * 60) // Convert hours to minutes
	}
	if maxItems, ok := args["max_items"].(float64); ok {
		log.Debug("Parsed max_items from args", "max_items", maxItems)
		opts.MaxTasks = int(maxItems)
	}
	if instructions, ok := args["instructions"].(string); ok {
		log.Debug("Parsed instructions from args", "instructions", instructions)
		opts.Instructions = instructions
	}

	log.Debug("Calling service.GetFocusTasks", "opts", opts)
	tasks, err := service.GetFocusTasks(context.Background(), &opts)
	if err != nil {
		log.Error("Failed to get focus tasks", "error", err)
		return nil, err
	}
	log.Debug("service.GetFocusTasks returned", "count", len(tasks))

	// Format response for Claude
	resp := map[string]interface{}{
		"message": "Focus tasks retrieved successfully",
		"summary": map[string]interface{}{
			"total_tasks":   len(tasks),
			"energy_filter": opts.Energy,
			"mode_filter":   opts.Mode,
			"target_hours":  float64(opts.MaxMinutes) / 60.0,
		},
		"tasks": tasks,
	}
	log.Debug("handleDailyFocus response ready", "resp", resp)
	return resp, nil
}

// handleGetTaskMetadata retrieves metadata for a specific task
func handleGetTaskMetadata(service *Service, args map[string]interface{}) (interface{}, error) {
	log.Debug("handleGetTaskMetadata called", "args", args)
	taskIDFloat, ok := args["task_id"].(float64)
	if !ok {
		log.Error("task_id missing or not a number", "args", args)
		return nil, fmt.Errorf("task_id must be a number")
	}
	taskID := int64(taskIDFloat)

	log.Debug("Calling service.GetTaskMetadata", "task_id", taskID)
	task, err := service.GetTaskMetadata(context.Background(), taskID)
	if err != nil {
		log.Error("Failed to get task metadata", "task_id", taskID, "error", err)
		return nil, err
	}

	if task == nil {
		log.Debug("No metadata found for task", "task_id", taskID)
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

	log.Debug("Task metadata found", "task_id", taskID, "metadata", task.Metadata)
	// Return enriched task data
	resp := map[string]interface{}{
		"task_id":             task.RawTask.ID,
		"title":               task.RawTask.Title,
		"description":         task.CleanDescription,
		"hex_color":           task.RawTask.HexColor,
		"done":                task.RawTask.Done,
		"priority":            task.RawTask.Priority,
		"project_id":          task.RawTask.ProjectID,
		"has_hyperfocus_data": task.Metadata != nil,
		"metadata":            task.Metadata,
		"created":             task.RawTask.Created,
		"updated":             task.RawTask.Updated,
	}
	log.Debug("handleGetTaskMetadata response ready", "resp", resp)
	return resp, nil
}

// handleUpsertTask processes the upsert_task tool call
func handleUpsertTask(service *Service, args map[string]interface{}) (interface{}, error) {
	log.Debug("handleUpsertTask called", "args", args)
	var task models.RawTask

	// Parse task data from arguments
	if v, ok := args["task_id"].(float64); ok {
		log.Debug("Parsed task_id from args", "task_id", v)
		task.ID = int64(v)
	}
	if v, ok := args["project_id"].(float64); ok {
		log.Debug("Parsed project_id from args", "project_id", v)
		task.ProjectID = int64(v)
	}
	if v, ok := args["title"].(string); ok {
		log.Debug("Parsed title from args", "title", v)
		task.Title = v
	}
	if v, ok := args["description"].(string); ok {
		log.Debug("Parsed description from args", "description", v)
		task.Description = v
	}
	if v, ok := args["priority"].(float64); ok {
		log.Debug("Parsed priority from args", "priority", v)
		task.Priority = int(v)
	}
	if v, ok := args["hex_color"].(string); ok {
		log.Debug("Parsed Hex Color from args", "done", v)
		task.HexColor = v
	}
	if v, ok := args["done"].(bool); ok {
		log.Debug("Parsed done from args", "done", v)
		task.Done = v
	}

	action := "updated"
	if task.ID == 0 {
		log.Debug("No task_id provided, will create new task")
		action = "created"
	}

	log.Debug("Calling service.UpsertTask", "action", action, "task", task)
	result, err := service.UpsertTask(context.Background(), &task)
	if err != nil {
		log.Error("Failed to upsert task", "error", err, "action", action)
		return nil, err
	}

	resp := map[string]interface{}{
		"success": true,
		"action":  action,
		"task": map[string]interface{}{
			"task_id":     result.ID,
			"title":       result.Title,
			"done":        result.Done,
			"priority":    result.Priority,
			"description": result.Description,
			"hex_color":   result.HexColor,
			"project_id":  result.ProjectID,
		},
		"message": fmt.Sprintf("Task %s successfully", action),
	}
	log.Debug("handleUpsertTask response ready", "resp", resp)
	return resp, nil
}

// handleGetFilteredTasks retrieves tasks using filter expressions or natural language
func handleGetFilteredTasks(service *Service, args map[string]interface{}) (interface{}, error) {
	log.Debug("handleGetFilteredTasks called", "args", args)

	var filter string
	var useAI bool

	// Check if we have a direct filter expression
	if filterArg, ok := args["filter"].(string); ok && filterArg != "" {
		filter = filterArg
		useAI = false
		log.Debug("Using direct filter expression", "filter", filter)
	} else if naturalRequest, ok := args["natural_request"].(string); ok && naturalRequest != "" {
		filter = naturalRequest
		useAI = true
		log.Debug("Using natural language request", "request", naturalRequest)
	} else {
		log.Error("Neither filter nor natural_request provided", "args", args)
		return nil, fmt.Errorf("either 'filter' or 'natural_request' must be provided")
	}

	log.Debug("Calling service.GetFilteredTasks", "filter", filter, "useAI", useAI)
	tasks, err := service.GetFilteredTasks(context.Background(), filter, useAI)
	if err != nil {
		log.Error("Failed to get filtered tasks", "error", err, "filter", filter)
		return nil, err
	}

	log.Debug("service.GetFilteredTasks returned", "count", len(tasks))

	// Format response for Claude
	resp := map[string]interface{}{
		"message": "Filtered tasks retrieved successfully",
		"summary": map[string]interface{}{
			"total_tasks": len(tasks),
			"filter_type": map[bool]string{true: "natural_language", false: "expression"}[useAI],
			"filter_used": filter,
		},
		"tasks": tasks,
	}

	log.Debug("handleGetFilteredTasks response ready", "resp", resp)
	return resp, nil
}
