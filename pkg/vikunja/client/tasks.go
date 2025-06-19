package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/BelKirill/vikunja-mcp/models"
	"github.com/charmbracelet/log"
)

// ListTasks returns all tasks visible to the authenticated user.
//
//	GET /api/v1/tasks
func (c *Client) GetAllTasks(ctx *context.Context) ([]models.RawTask, error) {
	log.Info("GetAllTasks called")
	var allTasks []models.RawTask
	page := 1
	for {
		var tasks []models.RawTask
		endpoint := fmt.Sprintf("/api/v1/tasks/all?page=%d", page)
		resp, err := c.getWithResponse(ctx, endpoint, &tasks)
		if err != nil {
			log.Error("Failed to fetch tasks", "page", page, "error", err)
			return nil, err
		}
		allTasks = append(allTasks, tasks...)

		totalPages := 1
		if resp != nil {
			totalPagesStr := resp.Header.Get("x-pagination-total-pages")
			if totalPagesStr != "" {
				if _, err := fmt.Sscanf(totalPagesStr, "%d", &totalPages); err != nil {
					log.Warn("Failed to parse x-pagination-total-pages header", "value", totalPagesStr, "error", err)
				}
			}
		}
		if page >= totalPages || len(tasks) == 0 {
			break
		}
		page++
	}
	log.Info("tasks fetched", "count", len(allTasks))
	log.Debug("Returning RawTasks", "count", len(allTasks))
	return allTasks, nil
}

// ListTasks returns all tasks visible to the authenticated user and passing the filter.
//
//	GET /api/v1/tasks
func (c *Client) GetFilteredTasks(ctx *context.Context, filter *string) ([]models.RawTask, error) {
	log.Info("GetFilteredTasks called")
	var allTasks []models.RawTask
	page := 1
	for {
		var tasks []models.RawTask
		endpoint := fmt.Sprintf("/api/v1/tasks/all?page=%d&filter=%s", page, url.QueryEscape(*filter))
		log.Debug("Fetching tasks with endpoint", "endpoint", endpoint)
		resp, err := c.getWithResponse(ctx, endpoint, &tasks)
		if err != nil {
			log.Error("Failed to fetch tasks", "page", page, "error", err)
			return nil, err
		}
		allTasks = append(allTasks, tasks...)

		totalPages := 1
		if resp != nil {
			totalPagesStr := resp.Header.Get("x-pagination-total-pages")
			if totalPagesStr != "" {
				if _, err := fmt.Sscanf(totalPagesStr, "%d", &totalPages); err != nil {
					log.Warn("Failed to parse x-pagination-total-pages header", "value", totalPagesStr, "error", err)
				}
			}
		}
		if page >= totalPages || len(tasks) == 0 {
			break
		}
		page++
	}
	log.Info("tasks fetched", "count", len(allTasks))
	log.Debug("Returning RawTasks", "count", len(allTasks))
	return allTasks, nil
}

// GetTask returns a single task by its ID.
//
//	GET /api/v1/tasks/{id}
func (c *Client) GetTask(ctx *context.Context, id int64) (*models.RawTask, error) {
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

// UpsertTask creates a new task or updates an existing one using RawTask
func (c *Client) UpsertTask(ctx *context.Context, taskData models.RawTask) (*models.RawTask, error) {
	log.Info("UpsertTask called", "task_id", taskData.ID)
	log.Debug("UpsertTask details", "taskData", taskData)
	if taskData.ID == 0 {
		return c.createTask(ctx, &taskData)
	}
	return c.updateTask(ctx, &taskData)
}

func (c *Client) createTask(ctx *context.Context, taskData *models.RawTask) (*models.RawTask, error) {
	log.Info("createTask called", "project", taskData.ProjectID)
	log.Debug("createTask details", "taskData", taskData)
	endpoint := fmt.Sprintf("/api/v1/projects/%d/tasks", taskData.ProjectID)
	requestBody := map[string]interface{}{
		"title":       taskData.Title,
		"description": taskData.Description,
		"hex_color":   taskData.HexColor,
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
	return &result, nil
}

func (c *Client) updateTask(ctx *context.Context, taskData *models.RawTask) (*models.RawTask, error) {
	log.Info("updateTask called", "task_id", taskData.ID)
	log.Debug("updateTask details", "taskData", taskData)
	endpoint := fmt.Sprintf("/api/v1/tasks/%d", taskData.ID)
	requestBody := map[string]interface{}{
		"title":       taskData.Title,
		"description": taskData.Description,
		"hex_color":   taskData.HexColor,
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
	return &result, nil
}

// getWithResponse performs a GET request and returns the http.Response for header inspection.
func (c *Client) getWithResponse(ctx *context.Context, endpoint string, result interface{}) (*http.Response, error) {
	log.Info("GET request (with response)", "endpoint", endpoint)
	fullURL := c.baseURL + endpoint
	req, err := c.newRequest(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := resp.Body.Close()
		if cerr != nil {
			log.Error("error closing response body", "error", cerr)
		}
	}()
	if resp.StatusCode >= 400 {
		return resp, fmt.Errorf("vikunja API error: %s", resp.Status)
	}
	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return resp, err
		}
	}
	return resp, nil
}
