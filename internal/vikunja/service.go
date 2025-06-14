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
	vikClient, err := client.NewClient()
	if err != nil {
		return nil, err
	}
	return &Service{Client: vikClient}, nil
}

// GetUserTasks fetches all tasks for the current user and enriches them with metadata.
func (s *Service) GetUserTasks(ctx context.Context) ([]models.Task, error) {
	rawTasks, err := s.Client.GetAllTasks(ctx)
	if err != nil {
		return nil, err
	}

	// Convert MinimalTask to RawTask - you may need to adjust this based on your actual types
	// For now assuming they're compatible or you have a conversion method
	var rawTasksConverted []models.RawTask
	for _, task := range rawTasks {
		// Convert MinimalTask to RawTask format
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
		return nil, err
	}

	return enrichedTasks, nil
}

// GetTaskByID fetches a single task by its ID.
func (s *Service) GetTaskByID(ctx context.Context, id int64) (*models.MinimalTask, error) {
	task, err := s.Client.GetTask(ctx, id)
	if err != nil {
		return nil, err
	}
	return enrichMinimalTask(task), nil
}

// GetIncompleteTasks returns all tasks that are not marked as done.
func (s *Service) GetIncompleteTasks(ctx context.Context) ([]models.Task, error) {
	tasks, err := s.GetUserTasks(ctx)
	if err != nil {
		return nil, err
	}
	var result []models.Task
	for _, t := range tasks {
		if !t.RawTask.Done {
			result = append(result, t)
		}
	}
	return result, nil
}

// UpsertTask creates or updates a task with proper metadata embedding
func (s *Service) UpsertTask(ctx context.Context, task models.MinimalTask) (*models.MinimalTask, error) {
	enriched := enrichMinimalTask(&task)

	// CRITICAL: If metadata is provided via description field (from MCP),
	// treat the description AS the metadata JSON and embed it properly
	if enriched.Description != "" {
		// Check if description contains JSON metadata
		if strings.Contains(enriched.Description, "{") && strings.Contains(enriched.Description, "energy") {
			// Parse the JSON metadata from description
			var metadata models.HyperFocusMetadata
			if err := json.Unmarshal([]byte(enriched.Description), &metadata); err == nil {
				// Successfully parsed - use this metadata
				enriched.Metadata = &metadata
				// Clear the description since it was just JSON metadata
				enriched.Description = ""
			}
		}
	}

	// If we have metadata, embed it properly in the description
	if enriched.Metadata != nil {
		enriched.Description = embedMetadataInDescription(enriched.Description, enriched.Metadata)
	}

	result, err := s.Client.UpsertTask(ctx, *enriched)
	if err != nil {
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
