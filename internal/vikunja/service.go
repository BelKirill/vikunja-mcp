package vikunja

import (
	"context"
	"encoding/json"
	"strings"

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
	vikClient, err := client.NewClient()
	if err != nil {
		log.Error("Failed to create Vikunja client", "error", err)
		return nil, err
	}
	log.Info("Vikunja client created successfully")
	return &Service{Client: vikClient}, nil
}

// GetUserTasks fetches all tasks for the current user and enriches them with metadata.
func (s *Service) GetUserTasks(ctx context.Context) ([]models.Task, error) {
	log.Info("GetUserTasks called")
	rawTasks, err := s.Client.GetAllTasks(ctx)
	if err != nil {
		log.Error("Failed to get all tasks", "error", err)
		return nil, err
	}
	log.Debug("Fetched raw tasks", "count", len(rawTasks))
	var rawTasksConverted []models.RawTask
	for _, task := range rawTasks {
		log.Debug("Converting MinimalTask to RawTask", "task_id", task.TaskID)
		rawTask := models.RawTask{
			ID:          task.TaskID,
			Title:       task.Title,
			Description: task.Description,
			Priority:    task.Priority,
			Done:        task.Done,
			ProjectID:   task.Project,
		}
		rawTasksConverted = append(rawTasksConverted, rawTask)
	}
	enrichedTasks, err := enrichTasks(rawTasksConverted)
	if err != nil {
		log.Error("Failed to enrich tasks", "error", err)
		return nil, err
	}
	log.Info("GetUserTasks returning", "enriched_count", len(enrichedTasks))
	return enrichedTasks, nil
}

// GetTaskByID fetches a single task by its ID.
func (s *Service) GetTaskByID(ctx context.Context, id int64) (*models.Task, error) {
	log.Info("GetTaskByID called", "id", id)
	task, err := s.Client.GetTask(ctx, id)
	if err != nil {
		log.Error("Failed to get task by ID", "id", id, "error", err)
		return nil, err
	}
	log.Debug("Fetched task", "task", task)
	result, err := enrichTask(task)
	if err != nil {
		log.Error("Failed to enrich task by ID", "id", id, "error", err)
		return nil, err
	}
	log.Info("GetTaskByID returning", "id", id, "has_metadata", result.Metadata != nil)
	return result, nil
}

// GetIncompleteTasks returns all tasks that are not marked as done.
func (s *Service) GetIncompleteTasks(ctx context.Context) ([]models.Task, error) {
	log.Info("GetIncompleteTasks called")
	tasks, err := s.GetUserTasks(ctx)
	if err != nil {
		log.Error("Failed to get user tasks", "error", err)
		return nil, err
	}
	var result []models.Task
	for _, t := range tasks {
		if !t.RawTask.Done {
			log.Debug("Task is incomplete", "task_id", t.RawTask.ID)
			result = append(result, t)
		}
	}
	log.Info("GetIncompleteTasks returning", "count", len(result))
	return result, nil
}

// UpsertTask creates or updates a task with proper metadata embedding
func (s *Service) UpsertTask(ctx context.Context, task models.MinimalTask) (*models.MinimalTask, error) {
	log.Info("UpsertTask called", "task_id", task.TaskID)
	enriched := enrichMinimalTask(&task)

	// CRITICAL: If metadata is provided via description field (from MCP),
	// treat the description AS the metadata JSON and embed it properly
	if enriched.Description != "" {
		// Check if description contains JSON metadata
		if strings.Contains(enriched.Description, "{") && strings.Contains(enriched.Description, "energy") {
			// Parse the JSON metadata from description
			var metadata models.HyperFocusMetadata
			if err := json.Unmarshal([]byte(enriched.Description), &metadata); err == nil {
				log.Debug("Parsed metadata from description JSON", "task_id", task.TaskID)
				enriched.Metadata = &metadata
				// Clear the description since it was just JSON metadata
				enriched.Description = ""
			}
		}
	}

	// If we have metadata, embed it properly in the description
	if enriched.Metadata != nil {
		log.Debug("Embedding metadata in description", "task_id", task.TaskID)
		enriched.Description = embedMetadataInDescription(enriched.Description, enriched.Metadata)
	}

	result, err := s.Client.UpsertTask(ctx, *enriched)
	if err != nil {
		log.Error("Failed to upsert task", "task_id", task.TaskID, "error", err)
		return nil, err
	}

	// Log successful operation
	action := "updated"
	if task.TaskID == 0 {
		action = "created"
	}
	log.Info("task upserted successfully",
		"action", action,
		"task_id", result.TaskID,
		"title", result.Title,
		"description_length", len(result.Description),
		"has_metadata", result.Metadata != nil)

	return result, nil
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
