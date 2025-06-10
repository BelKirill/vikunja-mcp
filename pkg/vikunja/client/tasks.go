package client

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"
)

// Task represents a minimal Vikunja task model.
// Add more fields later as needed.
type Task struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Done        bool   `json:"done"`
}

// ListTasks returns all tasks visible to the authenticated user.
//
//	GET /api/v1/tasks
func (c *Client) GetAllTasks(ctx context.Context) ([]Task, error) {
	var tasks []Task
	if err := c.Get(ctx, "/api/v1/tasks/all", &tasks); err != nil {
		return nil, err
	}
	log.Info("tasks fetched", "count", len(tasks))
	return tasks, nil
}

// GetTask returns a single task by its ID.
//
//	GET /api/v1/tasks/{id}
func (c *Client) GetTask(ctx context.Context, id int64) (*Task, error) {
	endpoint := fmt.Sprintf("/api/v1/tasks/%d", id)

	var t Task
	if err := c.Get(ctx, endpoint, &t); err != nil {
		return nil, err
	}
	log.Info("task fetched", "id", id, "title", t.Title)
	return &t, nil
}
