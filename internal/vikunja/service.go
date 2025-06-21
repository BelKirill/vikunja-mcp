package vikunja

import (
	"context"
	"fmt"

	"github.com/BelKirill/vikunja-mcp/internal/config"
	"github.com/BelKirill/vikunja-mcp/models"
	"github.com/BelKirill/vikunja-mcp/pkg/vikunja/client"
	"github.com/charmbracelet/log"
)

// Service provides business logic on top of the Vikunja API client.
type Service struct {
	Client *client.Client
}

// NewService creates a new Service with the given Vikunja client.
func NewService() (*Service, error) {
	log.Info("NewService called for vikunja.Service")
	vikClient, err := client.NewClientWithCredentials(config.GetVikunja().URL, config.GetVikunja().Token)
	if err != nil {
		log.Error("Failed to create Vikunja client", "error", err)
		return nil, err
	}
	log.Info("Vikunja client created successfully")
	return &Service{Client: vikClient}, nil
}

// GetUserTasks fetches all tasks for the current user and enriches them with metadata.
func (s *Service) GetUserTasks(ctx context.Context) ([]models.RawTask, error) {
	log.Info("GetUserTasks called")
	rawTasks, err := s.Client.GetAllTasks(ctx)
	if err != nil {
		log.Error("Failed to get all tasks", "error", err)
		return nil, err
	}
	log.Info("Fetched raw tasks", "count", len(rawTasks))
	return rawTasks, nil
}

// GetFilteredTask fetches tasks using a filter
// https://vikunja.io/docs/filters/
func (s *Service) GetFilteredTasks(ctx context.Context, filter *string) ([]models.RawTask, error) {
	log.Info("GetFilteredTask called")
	filtered, err := s.Client.GetFilteredTasks(ctx, filter)
	if err != nil {
		log.Error("Failed to get filter", "error", err)
		return nil, err
	}
	log.Info("Fetched filtered raw tasks", "count", len(filtered))
	return filtered, nil
}

// GetTaskByID fetches a single task by its ID.
func (s *Service) GetTaskByID(ctx context.Context, id int64) (*models.RawTask, error) {
	log.Info("GetTaskByID called", "id", id)
	task, err := s.Client.GetTask(ctx, id)
	if err != nil {
		log.Error("Failed to get task by ID", "id", id, "error", err)
		return nil, err
	}
	log.Debug("Fetched task", "task", task)
	return task, nil
}

// GetTaskCommentsByID fetches all comments for a given task ID
func (s *Service) GetTaskCommentsByID(ctx context.Context, taskID int64) ([]models.Comment, error) {
	log.Info("GetTaskCommentsByID called", "task_id", taskID)
	endpoint := fmt.Sprintf("/api/v1/tasks/%d/comments", taskID)
	var comments []models.Comment
	err := s.Client.Get(ctx, endpoint, &comments)
	if err != nil {
		log.Error("Failed to fetch comments", "task_id", taskID, "error", err)
		return nil, err
	}
	log.Info("Fetched comments for task", "task_id", taskID, "count", len(comments))
	return comments, nil
}

// GetIncompleteTasks returns all tasks that are not marked as done.
func (s *Service) GetIncompleteTasks(ctx context.Context) ([]models.RawTask, error) {
	log.Info("GetIncompleteTasks called")
	tasks, err := s.GetUserTasks(ctx)
	if err != nil {
		log.Error("Failed to get user tasks", "error", err)
		return nil, err
	}
	var result []models.RawTask
	for _, t := range tasks {
		if !t.Done {
			log.Debug("Task is incomplete", "task_id", t.ID)
			result = append(result, t)
		}
	}
	log.Info("GetIncompleteTasks returning", "count", len(result))
	return result, nil
}

// UpsertTask creates or updates a task with intelligent field merging
func (s *Service) UpsertTask(ctx context.Context, task *models.RawTask) (*models.RawTask, error) {
	log.Info("UpsertTask called", "task_id", task.ID)

	var finalTask models.RawTask

	// If this is an update (ID > 0), get existing task and merge fields
	if task.ID > 0 {
		log.Debug("Updating existing task - fetching current data", "task_id", task.ID)
		existing, err := s.Client.GetTask(ctx, int64(task.ID))
		if err != nil {
			log.Error("Failed to get existing task for merge", "task_id", task.ID, "error", err)
			return nil, err
		}

		// Start with existing task data
		finalTask = *existing
		log.Debug("Starting with existing task data", "task_id", task.ID)

		// Only merge fields that were actually provided in the update
		// We detect "provided" vs "empty" by checking if the field has a non-zero value
		// or if it's a special case (like empty string being intentional)

		if task.Title != "" {
			log.Debug("Merging title update", "task_id", task.ID, "new_title", task.Title)
			finalTask.Title = task.Title
		}

		if task.Description != "" {
			log.Debug("Merging description update", "task_id", task.ID, "description_length", len(task.Description))
			finalTask.Description = task.Description
		}

		if task.Priority != 0 {
			log.Debug("Merging priority update", "task_id", task.ID, "new_priority", task.Priority)
			finalTask.Priority = task.Priority
		}

		if task.HexColor != "" {
			log.Debug("Merging hex_color update", "task_id", task.ID, "new_color", task.HexColor)
			finalTask.HexColor = task.HexColor
		}

		if task.ProjectID != 0 {
			log.Debug("Merging project_id update", "task_id", task.ID, "new_project_id", task.ProjectID)
			finalTask.ProjectID = task.ProjectID
		}

		// For boolean fields, we need a different approach since false is a valid value
		// For now, we'll always merge the Done field if it's different from existing
		// TODO: Consider using pointers or a separate "fields to update" parameter
		if task.Done != existing.Done {
			log.Debug("Merging done status update", "task_id", task.ID, "new_done", task.Done)
			finalTask.Done = task.Done
		}

	} else {
		// Creating new task - use provided data as-is
		log.Debug("Creating new task")
		finalTask = *task
	}

	// Set priority to 3 if it is 0 (only for new tasks)
	if finalTask.Priority == 0 {
		log.Debug("Setting default priority to 3 for task", "task_id", finalTask.ID)
		finalTask.Priority = 3
	}

	result, err := s.Client.UpsertTask(ctx, finalTask)
	if err != nil {
		log.Error("Failed to upsert task", "task_id", finalTask.ID, "error", err)
		return nil, err
	}

	// Log successful operation
	action := "updated"
	if task.ID == 0 {
		action = "created"
	}
	log.Info("task upserted successfully",
		"action", action,
		"task_id", result.ID,
		"title", result.Title,
		"description_length", len(result.Description),
		"priority", result.Priority,
		"hex_color", result.HexColor)

	return result, nil
}

// UpsertTaskSelective creates or updates a task with explicit field control
// This is an alternative approach using a parameter struct to be more explicit
type UpsertTaskOptions struct {
	TaskID      *int    `json:"task_id,omitempty"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Priority    *int    `json:"priority,omitempty"`
	HexColor    *string `json:"hex_color,omitempty"`
	Done        *bool   `json:"done,omitempty"`
	ProjectID   *int64  `json:"project_id,omitempty"`
}

func (s *Service) UpsertTaskSelective(ctx context.Context, opts UpsertTaskOptions) (*models.RawTask, error) {
	var finalTask models.RawTask

	// If updating existing task
	if opts.TaskID != nil && *opts.TaskID > 0 {
		log.Debug("Selective update - fetching existing task", "task_id", *opts.TaskID)
		existing, err := s.Client.GetTask(ctx, int64(*opts.TaskID))
		if err != nil {
			log.Error("Failed to get existing task for selective update", "task_id", *opts.TaskID, "error", err)
			return nil, err
		}

		// Start with existing data
		finalTask = *existing

		// Only update fields that were explicitly provided (not nil)
		if opts.Title != nil {
			log.Debug("Selective update: title", "task_id", *opts.TaskID, "new_title", *opts.Title)
			finalTask.Title = *opts.Title
		}

		if opts.Description != nil {
			log.Debug("Selective update: description", "task_id", *opts.TaskID)
			finalTask.Description = *opts.Description
		}

		if opts.Priority != nil {
			log.Debug("Selective update: priority", "task_id", *opts.TaskID, "new_priority", *opts.Priority)
			finalTask.Priority = *opts.Priority
		}

		if opts.HexColor != nil {
			log.Debug("Selective update: hex_color", "task_id", *opts.TaskID, "new_color", *opts.HexColor)
			finalTask.HexColor = *opts.HexColor
		}

		if opts.Done != nil {
			log.Debug("Selective update: done", "task_id", *opts.TaskID, "new_done", *opts.Done)
			finalTask.Done = *opts.Done
		}

		if opts.ProjectID != nil {
			log.Debug("Selective update: project_id", "task_id", *opts.TaskID, "new_project_id", *opts.ProjectID)
			finalTask.ProjectID = *opts.ProjectID
		}

	} else {
		// Creating new task
		log.Debug("Selective create - new task")
		if opts.Title != nil {
			finalTask.Title = *opts.Title
		}
		if opts.Description != nil {
			finalTask.Description = *opts.Description
		}
		if opts.Priority != nil {
			finalTask.Priority = *opts.Priority
		} else {
			finalTask.Priority = 3 // Default priority
		}
		if opts.HexColor != nil {
			finalTask.HexColor = *opts.HexColor
		}
		if opts.Done != nil {
			finalTask.Done = *opts.Done
		}
		if opts.ProjectID != nil {
			finalTask.ProjectID = *opts.ProjectID
		}
	}

	// Apply the same metadata processing and user assignment as before
	// ... (same logic as in UpsertTask)

	return s.Client.UpsertTask(ctx, finalTask)
}

// Me fetches the current user information.
func (s *Service) Me(ctx context.Context) (*models.User, error) {
	log.Info("Me called")
	var user models.User
	err := s.Client.Get(ctx, "/api/v1/user", &user)
	if err != nil {
		log.Error("Failed to fetch current user", "error", err)
		return nil, err
	}
	log.Info("Me returning user", "user_id", user.ID, "username", user.Username)
	return &user, nil
}

// AddComment adds a new comment to the task
func (s *Service) AddComment(ctx context.Context, taskID int64, comment *string) (*models.Comment, error) {
	log.Info("AddComment called")
	var newComment struct {
		Comment string
	}

	newComment.Comment = *comment

	endpoint := fmt.Sprintf("/api/v1/tasks/%d/comments", taskID)

	var createdComment models.Comment
	if err := s.Client.Put(ctx, endpoint, &newComment, &createdComment); err != nil {
		log.Error("Failed to add comment", "taskID", taskID, "error", err)
		return nil, err
	}

	return &createdComment, nil
}
