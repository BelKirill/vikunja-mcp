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
func (s *Service) GetUserTasks(ctx context.Context) ([]models.MinimalTask, error) {
	return s.Client.GetAllTasks(ctx)
}

// GetTaskByID fetches a single task by its ID.
func (s *Service) GetTaskByID(ctx context.Context, id int64) (*models.MinimalTask, error) {
	return s.Client.GetTask(ctx, id)
}

// GetIncompleteTasks returns all tasks that are not marked as done.
func (s *Service) GetIncompleteTasks(ctx context.Context) ([]models.MinimalTask, error) {
	tasks, err := s.GetUserTasks(ctx)
	if err != nil {
		return nil, err
	}
	var result []models.MinimalTask
	for _, t := range tasks {
		if !t.Done {
			result = append(result, t)
		}
	}
	return result, nil
}

// Example: Find tasks by label.
func (s *Service) GetTasksByLabel(ctx context.Context, label string) ([]models.MinimalTask, error) {
	tasks, err := s.GetUserTasks(ctx)
	if err != nil {
		return nil, err
	}
	// This is a stub; label filtering would require label info in MinimalTask
	return tasks, nil
}

func (s *Service) UpsertTask(ctx context.Context, task models.MinimalTask) (*models.MinimalTask, error) {
	return s.Client.UpsertTask(ctx, task)
}
