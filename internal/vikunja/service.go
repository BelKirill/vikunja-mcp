package vikunja

import (
	"context"

	"github.com/BelKirill/vikunja-mcp/models"
	"github.com/BelKirill/vikunja-mcp/pkg/vikunja/client"
)

// Service provides business logic on top of the Vikunja API client.
type Service struct {
	Client *client.Client
}

// NewService creates a new Service with the given Vikunja client.
func NewService() (*Service, error) {
	vikClient, err := client.NewClient()
	if err != nil {
		return nil, err
	}
	return &Service{Client: vikClient}, nil
}

// GetUserTasks fetches all tasks for the current user.
func (s *Service) GetUserTasks(ctx context.Context) ([]client.RawTask, error) {
	tasks, err := s.Client.GetAllTasks(ctx)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// GetTaskByID fetches a single task by its ID.
func (s *Service) GetTaskByID(ctx context.Context, id int64) (*client.RawTask, error) {
	return s.Client.GetTask(ctx, id)
}

// GetIncompleteTasks returns all tasks that are not marked as done.
func (s *Service) GetIncompleteTasks(ctx context.Context) ([]client.RawTask, error) {
	tasks, err := s.GetUserTasks(ctx)
	if err != nil {
		return nil, err
	}
	var result []client.RawTask
	for _, t := range tasks {
		if !t.Done {
			result = append(result, t)
		}
	}
	return result, nil
}

// Example: Find tasks by label.
func (s *Service) GetTasksByLabel(ctx context.Context, label string) ([]client.RawTask, error) {
	tasks, err := s.GetUserTasks(ctx)
	if err != nil {
		return nil, err
	}
	var result []client.RawTask
	for _, t := range tasks {
		for _, l := range t.Labels {
			if l.Title == label {
				result = append(result, t)
				break
			}
		}
	}
	return result, nil
}

// UpsertTask creates or updates a task using MinimalTask as the DTO.
func (s *Service) UpsertTask(ctx context.Context, task models.MinimalTask) (*models.MinimalTask, error) {
	var raw client.RawTask
	// Convert MinimalTask to RawTask
	raw.ID = int(task.TaskID)
	raw.ProjectID = int64(task.Project)
	raw.Title = task.Title
	raw.Description = task.Description
	raw.Priority = task.Priority
	raw.Done = task.Done
	// Metadata is ignored for now

	result, err := s.Client.UpsertTask(ctx, raw)
	if err != nil {
		return nil, err
	}
	// Convert RawTask back to MinimalTask
	return &models.MinimalTask{
		TaskID:      int64(result.ID),
		Project:     int64(result.ProjectID),
		Title:       result.Title,
		Description: result.Description,
		Priority:    result.Priority,
		Done:        result.Done,
		// Metadata: nil (not handled)
	}, nil
}
