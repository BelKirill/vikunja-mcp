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

	enrichedTasks, err := enrichTasks(rawTasks)
	if err != nil {
		log.Error("Failed to enrich tasks", "error", err)
		return nil, err
	}
	log.Info("GetUserTasks returning", "enriched_count", len(enrichedTasks))
	return enrichedTasks, nil
}

// GetFilteredTask fetches tasks using a filter
// https://vikunja.io/docs/filters/
func (s *Service) GetFilteredTasks(ctx context.Context, filter *string) ([]models.Task, error) {
	log.Info("GetFilteredTask called")
	filtered, err := s.Client.GetFilteredTasks(ctx, filter)
	if err != nil {
		log.Error("Failed to get filter", "error", err)
		return nil, err
	}
	log.Debug("Fetched filtered raw tasks", "count", len(filtered))

	enrichedTasks, err := enrichTasks(filtered)
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
func (s *Service) UpsertTask(ctx context.Context, task models.RawTask) (*models.RawTask, error) {
	log.Info("UpsertTask called", "task_id", task.ID)

	// If metadata is provided via description field (from MCP),
	// treat the description AS the metadata JSON and embed it properly
	if task.Description != "" {
		// Check if description contains JSON metadata
		if strings.Contains(task.Description, "{") && strings.Contains(task.Description, "energy") {
			// Parse the JSON metadata from description
			var metadata models.HyperFocusMetadata
			if err := json.Unmarshal([]byte(task.Description), &metadata); err == nil {
				log.Debug("Parsed metadata from description JSON", "task_id", task.ID)
				// Embed metadata as JSON in description
				task.Description = embedMetadataInDescription("", &metadata)
			}
		}
	}

	// Assign the value of the vikunja service Me to assignee
	if user, err := s.Me(ctx); err == nil && user != nil {
		task.Assignees = []models.User{*user}
	} else if err != nil {
		log.Warn("Could not fetch current user for assignee", "error", err)
	}

	// Set priority to 3 if it is 0
	if task.Priority == 0 {
		log.Debug("Setting default priority to 3 for task", "task_id", task.ID)
		task.Priority = 3
	}

	result, err := s.Client.UpsertTask(ctx, task)
	if err != nil {
		log.Error("Failed to upsert task", "task_id", task.ID, "error", err)
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
		"description_length", len(result.Description))

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
