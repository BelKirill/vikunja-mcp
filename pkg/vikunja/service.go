package vikunja

import (
	"context"
	"fmt"

	"github.com/BelKirill/vikunja-mcp/pkg/vikunja/client"
	"github.com/BelKirill/vikunja-mcp/pkg/vikunja/models"
	"github.com/charmbracelet/log"
)

// Service provides methods to interact with the Vikunja API
type Service struct {
	client *client.Client
}

// NewService creates a new Vikunja service instance
func NewService() (*Service, error) {
	c, err := client.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create Vikunja client: %w", err)
	}
	return &Service{client: c}, nil
}

// GetTask retrieves a task by its ID
func (s *Service) GetTask(ctx context.Context, taskID string) (*models.Task, error) {
	log.Info("fetching task", "task_id", taskID)

	var task models.Task
	if err := s.client.Get(ctx, fmt.Sprintf("/api/v1/tasks/%s", taskID), &task); err != nil {
		return nil, fmt.Errorf("failed to get task: %w", err)
	}

	return &task, nil
}

// GetProjectTasks retrieves all tasks for a project
func (s *Service) GetProjectTasks(ctx context.Context, projectID string) ([]models.Task, error) {
	log.Info("fetching project tasks", "project_id", projectID)

	var tasks []models.Task
	if err := s.client.Get(ctx, fmt.Sprintf("/api/v1/projects/%s/tasks", projectID), &tasks); err != nil {
		return nil, fmt.Errorf("failed to get project tasks: %w", err)
	}

	return tasks, nil
}

// CreateTask creates a new task
func (s *Service) CreateTask(ctx context.Context, task *models.Task) (*models.Task, error) {
	log.Info("creating task", "title", task.Title, "project_id", task.ProjectID)

	var createdTask models.Task
	if err := s.client.Post(ctx, "/api/v1/tasks", task, &createdTask); err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}

	return &createdTask, nil
}

// UpdateTask updates an existing task
func (s *Service) UpdateTask(ctx context.Context, taskID string, task *models.Task) (*models.Task, error) {
	log.Info("updating task", "task_id", taskID)

	var updatedTask models.Task
	if err := s.client.Put(ctx, fmt.Sprintf("/api/v1/tasks/%s", taskID), task, &updatedTask); err != nil {
		return nil, fmt.Errorf("failed to update task: %w", err)
	}

	return &updatedTask, nil
}

// DeleteTask deletes a task
func (s *Service) DeleteTask(ctx context.Context, taskID string) error {
	log.Info("deleting task", "task_id", taskID)

	if err := s.client.Delete(ctx, fmt.Sprintf("/api/v1/tasks/%s", taskID)); err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	return nil
}

// GetTasksByDueDate retrieves tasks due on a specific date
func (s *Service) GetTasksByDueDate(ctx context.Context, date string) ([]models.Task, error) {
	log.Info("fetching tasks by due date", "date", date)

	var tasks []models.Task
	if err := s.client.Get(ctx, fmt.Sprintf("/api/v1/tasks?due_date=%s", date), &tasks); err != nil {
		return nil, fmt.Errorf("failed to get tasks by due date: %w", err)
	}

	return tasks, nil
}

// GetTasksByPriority retrieves tasks with a specific priority
func (s *Service) GetTasksByPriority(ctx context.Context, priority int) ([]models.Task, error) {
	log.Info("fetching tasks by priority", "priority", priority)

	var tasks []models.Task
	if err := s.client.Get(ctx, fmt.Sprintf("/api/v1/tasks?priority=%d", priority), &tasks); err != nil {
		return nil, fmt.Errorf("failed to get tasks by priority: %w", err)
	}

	return tasks, nil
}

// GetTasksByLabel retrieves tasks with a specific label
func (s *Service) GetTasksByLabel(ctx context.Context, labelID int) ([]models.Task, error) {
	log.Info("fetching tasks by label", "label_id", labelID)

	var tasks []models.Task
	if err := s.client.Get(ctx, fmt.Sprintf("/api/v1/tasks?label_id=%d", labelID), &tasks); err != nil {
		return nil, fmt.Errorf("failed to get tasks by label: %w", err)
	}

	return tasks, nil
}

// GetTasksByStatus retrieves tasks with a specific status (done/undone)
func (s *Service) GetTasksByStatus(ctx context.Context, done bool) ([]models.Task, error) {
	log.Info("fetching tasks by status", "done", done)

	var tasks []models.Task
	if err := s.client.Get(ctx, fmt.Sprintf("/api/v1/tasks?done=%v", done), &tasks); err != nil {
		return nil, fmt.Errorf("failed to get tasks by status: %w", err)
	}

	return tasks, nil
}

// GetTasksByTimeRange retrieves tasks within a time range
func (s *Service) GetTasksByTimeRange(ctx context.Context, startDate, endDate string) ([]models.Task, error) {
	log.Info("fetching tasks by time range", "start_date", startDate, "end_date", endDate)

	var tasks []models.Task
	if err := s.client.Get(ctx, fmt.Sprintf("/api/v1/tasks?start_date=%s&end_date=%s", startDate, endDate), &tasks); err != nil {
		return nil, fmt.Errorf("failed to get tasks by time range: %w", err)
	}

	return tasks, nil
}

// GetTasksByEstimate retrieves tasks with a specific time estimate
func (s *Service) GetTasksByEstimate(ctx context.Context, estimateMin int) ([]models.Task, error) {
	log.Info("fetching tasks by estimate", "estimate_min", estimateMin)

	var tasks []models.Task
	if err := s.client.Get(ctx, fmt.Sprintf("/api/v1/tasks?estimate=%d", estimateMin), &tasks); err != nil {
		return nil, fmt.Errorf("failed to get tasks by estimate: %w", err)
	}

	return tasks, nil
}

// GetTasksByMetadata retrieves tasks with specific hyperfocus metadata
func (s *Service) GetTasksByMetadata(ctx context.Context, metadata *models.HyperfocusMetadata) ([]models.Task, error) {
	log.Info("fetching tasks by metadata",
		"energy", metadata.Energy,
		"mode", metadata.Mode,
		"extend", metadata.Extend,
		"minutes", metadata.Minutes)

	// Note: This is a custom endpoint that would need to be implemented on the Vikunja side
	// For now, we'll fetch all tasks and filter them client-side
	var tasks []models.Task
	if err := s.client.Get(ctx, "/api/v1/tasks", &tasks); err != nil {
		return nil, fmt.Errorf("failed to get tasks: %w", err)
	}

	// Filter tasks by metadata
	var filteredTasks []models.Task
	for _, task := range tasks {
		if task.Metadata != nil &&
			task.Metadata.Energy == metadata.Energy &&
			task.Metadata.Mode == metadata.Mode &&
			task.Metadata.Extend == metadata.Extend &&
			task.Metadata.Minutes == metadata.Minutes {
			filteredTasks = append(filteredTasks, task)
		}
	}

	return filteredTasks, nil
}

// GetTasksByProjectAndDueDate retrieves tasks for a project due on a specific date
func (s *Service) GetTasksByProjectAndDueDate(ctx context.Context, projectID, date string) ([]models.Task, error) {
	log.Info("fetching project tasks by due date", "project_id", projectID, "date", date)

	var tasks []models.Task
	if err := s.client.Get(ctx, fmt.Sprintf("/api/v1/projects/%s/tasks?due_date=%s", projectID, date), &tasks); err != nil {
		return nil, fmt.Errorf("failed to get project tasks by due date: %w", err)
	}

	return tasks, nil
}

// GetTasksByProjectAndPriority retrieves tasks for a project with a specific priority
func (s *Service) GetTasksByProjectAndPriority(ctx context.Context, projectID string, priority int) ([]models.Task, error) {
	log.Info("fetching project tasks by priority", "project_id", projectID, "priority", priority)

	var tasks []models.Task
	if err := s.client.Get(ctx, fmt.Sprintf("/api/v1/projects/%s/tasks?priority=%d", projectID, priority), &tasks); err != nil {
		return nil, fmt.Errorf("failed to get project tasks by priority: %w", err)
	}

	return tasks, nil
}

// GetTasksByProjectAndStatus retrieves tasks for a project with a specific status
func (s *Service) GetTasksByProjectAndStatus(ctx context.Context, projectID string, done bool) ([]models.Task, error) {
	log.Info("fetching project tasks by status", "project_id", projectID, "done", done)

	var tasks []models.Task
	if err := s.client.Get(ctx, fmt.Sprintf("/api/v1/projects/%s/tasks?done=%v", projectID, done), &tasks); err != nil {
		return nil, fmt.Errorf("failed to get project tasks by status: %w", err)
	}

	return tasks, nil
}

// GetTasksByProjectAndTimeRange retrieves tasks for a project within a time range
func (s *Service) GetTasksByProjectAndTimeRange(ctx context.Context, projectID, startDate, endDate string) ([]models.Task, error) {
	log.Info("fetching project tasks by time range",
		"project_id", projectID,
		"start_date", startDate,
		"end_date", endDate)

	var tasks []models.Task
	if err := s.client.Get(ctx, fmt.Sprintf("/api/v1/projects/%s/tasks?start_date=%s&end_date=%s",
		projectID, startDate, endDate), &tasks); err != nil {
		return nil, fmt.Errorf("failed to get project tasks by time range: %w", err)
	}

	return tasks, nil
}

// GetTasksByProjectAndEstimate retrieves tasks for a project with a specific time estimate
func (s *Service) GetTasksByProjectAndEstimate(ctx context.Context, projectID string, estimateMin int) ([]models.Task, error) {
	log.Info("fetching project tasks by estimate", "project_id", projectID, "estimate_min", estimateMin)

	var tasks []models.Task
	if err := s.client.Get(ctx, fmt.Sprintf("/api/v1/projects/%s/tasks?estimate=%d", projectID, estimateMin), &tasks); err != nil {
		return nil, fmt.Errorf("failed to get project tasks by estimate: %w", err)
	}

	return tasks, nil
}

// GetTasksByProjectAndMetadata retrieves tasks for a project with specific hyperfocus metadata
func (s *Service) GetTasksByProjectAndMetadata(ctx context.Context, projectID string, metadata *models.HyperfocusMetadata) ([]models.Task, error) {
	log.Info("fetching project tasks by metadata",
		"project_id", projectID,
		"energy", metadata.Energy,
		"mode", metadata.Mode,
		"extend", metadata.Extend,
		"minutes", metadata.Minutes)

	// Note: This is a custom endpoint that would need to be implemented on the Vikunja side
	// For now, we'll fetch all project tasks and filter them client-side
	var tasks []models.Task
	if err := s.client.Get(ctx, fmt.Sprintf("/api/v1/projects/%s/tasks", projectID), &tasks); err != nil {
		return nil, fmt.Errorf("failed to get project tasks: %w", err)
	}

	// Filter tasks by metadata
	var filteredTasks []models.Task
	for _, task := range tasks {
		if task.Metadata != nil &&
			task.Metadata.Energy == metadata.Energy &&
			task.Metadata.Mode == metadata.Mode &&
			task.Metadata.Extend == metadata.Extend &&
			task.Metadata.Minutes == metadata.Minutes {
			filteredTasks = append(filteredTasks, task)
		}
	}

	return filteredTasks, nil
}
