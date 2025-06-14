package client

import (
	"context"
	"fmt"

	"github.com/BelKirill/vikunja-mcp/models"
	"github.com/charmbracelet/log"
)

// ListTasks returns all tasks visible to the authenticated user.
//
//	GET /api/v1/tasks
func (c *Client) GetAllTasks(ctx context.Context) ([]models.MinimalTask, error) {
	log.Info("GetAllTasks called")
	var tasks []models.RawTask
	if err := c.Get(ctx, "/api/v1/tasks/all?limit=1000", &tasks); err != nil {
		log.Error("Failed to fetch all tasks", "error", err)
		return nil, err
	}
	log.Info("tasks fetched", "count", len(tasks))
	minimalTasks := make([]models.MinimalTask, 0, len(tasks))
	for _, t := range tasks {
		log.Debug("Converting RawTask to MinimalTask", "id", t.ID)
		minimalTasks = append(minimalTasks, models.MinimalTask{
			TaskID:      int64(t.ID),
			Project:     int64(t.ProjectID),
			Title:       t.Title,
			Description: t.Description,
			Priority:    t.Priority,
			Done:        t.Done,
		})
	}
	log.Debug("Returning minimalTasks", "count", len(minimalTasks))
	return minimalTasks, nil
}

// GetTask returns a single task by its ID.
//
//	GET /api/v1/tasks/{id}
func (c *Client) GetTask(ctx context.Context, id int64) (*models.RawTask, error) {
	log.Info("GetTask called", "id", id)
	endpoint := fmt.Sprintf("/api/v1/tasks/%d", id)
	var t models.RawTask
	if err := c.Get(ctx, endpoint, &t); err != nil {
		log.Error("Failed to fetch task", "id", id, "error", err)
		return nil, err
	}
	log.Info("task fetched", "id", id, "title", t.Title)
	log.Debug("Fetched task details", "task", t)
	return &t, nil
}

// UpsertTask creates a new task or updates an existing one using MinimalTask
func (c *Client) UpsertTask(ctx context.Context, taskData models.MinimalTask) (*models.MinimalTask, error) {
	log.Info("UpsertTask called", "task_id", taskData.TaskID)
	log.Debug("UpsertTask details", "taskData", taskData)
	if taskData.TaskID == 0 {
		return c.createTask(ctx, &taskData)
	}
	return c.updateTask(ctx, &taskData)
}

func (c *Client) createTask(ctx context.Context, taskData *models.MinimalTask) (*models.MinimalTask, error) {
	log.Info("createTask called", "project", taskData.Project)
	log.Debug("createTask details", "taskData", taskData)
	endpoint := fmt.Sprintf("/api/v1/projects/%d/tasks", taskData.Project)
	requestBody := map[string]interface{}{
		"title":       taskData.Title,
		"description": taskData.Description,
		"priority":    taskData.Priority,
		"done":        taskData.Done,
	}
	var result models.RawTask
	if err := c.Put(ctx, endpoint, requestBody, &result); err != nil {
		log.Error("Failed to create task", "error", err)
		return nil, err
	}
	log.Info("task created", "id", result.ID, "title", result.Title)
	log.Debug("Created task details", "task", result)
	return &models.MinimalTask{
		TaskID:      result.ID,
		Project:     result.ProjectID,
		Title:       result.Title,
		Description: result.Description,
		Priority:    result.Priority,
		Done:        result.Done,
	}, nil
}

func (c *Client) updateTask(ctx context.Context, taskData *models.MinimalTask) (*models.MinimalTask, error) {
	log.Info("updateTask called", "task_id", taskData.TaskID)
	log.Debug("updateTask details", "taskData", taskData)
	endpoint := fmt.Sprintf("/api/v1/tasks/%d", taskData.TaskID)
	requestBody := map[string]interface{}{
		"title":       taskData.Title,
		"description": taskData.Description,
		"priority":    taskData.Priority,
		"done":        taskData.Done,
	}
	var result models.RawTask
	if err := c.Post(ctx, endpoint, requestBody, &result); err != nil {
		log.Error("Failed to update task", "error", err)
		return nil, err
	}
	log.Info("task updated", "id", result.ID, "title", result.Title)
	log.Debug("Updated task details", "task", result)
	return &models.MinimalTask{
		TaskID:      result.ID,
		Project:     result.ProjectID,
		Title:       result.Title,
		Description: result.Description,
		Priority:    result.Priority,
		Done:        result.Done,
	}, nil
}
