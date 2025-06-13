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
	var tasks []RawTask
	if err := c.Get(ctx, "/api/v1/tasks/all", &tasks); err != nil {
		return nil, err
	}
	log.Info("tasks fetched", "count", len(tasks))
	minimalTasks := make([]models.MinimalTask, 0, len(tasks))
	for _, t := range tasks {
		minimalTasks = append(minimalTasks, models.MinimalTask{
			TaskID:      int64(t.ID),
			Project:     int64(t.ProjectID),
			Title:       t.Title,
			Description: t.Description,
			Priority:    t.Priority,
			Done:        t.Done,
		})
	}
	return minimalTasks, nil
}

// GetTask returns a single task by its ID.
//
//	GET /api/v1/tasks/{id}
func (c *Client) GetTask(ctx context.Context, id int64) (*models.MinimalTask, error) {
	endpoint := fmt.Sprintf("/api/v1/tasks/%d", id)

	var t RawTask
	if err := c.Get(ctx, endpoint, &t); err != nil {
		return nil, err
	}
	log.Info("task fetched", "id", id, "title", t.Title)
	return &models.MinimalTask{
		TaskID:      int64(t.ID),
		Project:     int64(t.ProjectID),
		Title:       t.Title,
		Description: t.Description,
		Priority:    t.Priority,
		Done:        t.Done,
	}, nil
}

// UpsertTask creates a new task or updates an existing one using MinimalTask
func (c *Client) UpsertTask(ctx context.Context, taskData models.MinimalTask) (*models.MinimalTask, error) {
	if taskData.TaskID == 0 {
		return c.createTask(ctx, &taskData)
	}
	return c.updateTask(ctx, &taskData)
}

func (c *Client) createTask(ctx context.Context, taskData *models.MinimalTask) (*models.MinimalTask, error) {
	endpoint := fmt.Sprintf("/api/v1/projects/%d/tasks", taskData.Project)
	
	// Create request body without ID (let Vikunja assign it)
	requestBody := RawTask{
		// ID: 0, // Don't set ID for creation - Vikunja assigns it
		ProjectID:   taskData.Project,
		Title:       taskData.Title,
		Description: taskData.Description,
		Priority:    taskData.Priority,
		Done:        taskData.Done,
	}
	
	var result RawTask
	// Use POST for creation (not PUT)
	if err := c.Put(ctx, endpoint, &requestBody, &result); err != nil {
		return nil, err
	}
	
	log.Info("task created", "id", result.ID, "title", result.Title)
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
	endpoint := fmt.Sprintf("/api/v1/tasks/%d", taskData.TaskID)
	
	// Create request body with existing ID
	requestBody := RawTask{
		ID:          taskData.TaskID,
		ProjectID:   taskData.Project,
		Title:       taskData.Title,
		Description: taskData.Description,
		Priority:    taskData.Priority,
		Done:        taskData.Done,
	}
	
	var result RawTask
	// Use PUT for updates (not POST)
	if err := c.Post(ctx, endpoint, &requestBody, &result); err != nil {
		return nil, err
	}
	
	log.Info("task updated", "id", result.ID, "title", result.Title)
	return &models.MinimalTask{
		TaskID:      result.ID,
		Project:     result.ProjectID,
		Title:       result.Title,
		Description: result.Description,
		Priority:    result.Priority,
		Done:        result.Done,
	}, nil
}