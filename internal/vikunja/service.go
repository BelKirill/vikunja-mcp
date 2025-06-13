package vikunja

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/BelKirill/vikunja-mcp/models"
	"github.com/BelKirill/vikunja-mcp/pkg/vikunja/client"
)

// extractJSON extracts the first valid JSON object or array from a string.
func extractJSON(input string) (string, error) {
	input = strings.TrimSpace(input)
	startObj := strings.Index(input, "{")
	startArr := strings.Index(input, "[")
	var start int
	if startObj == -1 && startArr == -1 {
		return "", nil // No JSON found
	} else if startObj == -1 {
		start = startArr
	} else if startArr == -1 {
		start = startObj
	} else if startObj < startArr {
		start = startObj
	} else {
		start = startArr
	}
	for end := len(input); end > start; end-- {
		candidate := input[start:end]
		var js interface{}
		if json.Unmarshal([]byte(candidate), &js) == nil {
			return candidate, nil
		}
	}
	return "", nil // No valid JSON found
}

func enrichTask(task *models.RawTask) (*models.Task, error) {
	meta, err := extractJSON(task.Description)
	if err != nil {
		return nil, err
	}

	enrichedTask := &models.Task{
		RawTask:          task,             // Embed the raw task
		CleanDescription: task.Description, // Will be cleaned below
	}

	if meta == "" {
		// No JSON metadata found - use defaults
		enrichedTask.Metadata = &models.HyperFocusMetadata{
			Energy:                  "medium", // Default energy level
			Mode:                    "quick",  // Default mode
			Extend:                  false,    // Default no extension
			Minutes:                 25,       // Default pomodoro
			Estimate:                25,       // Default estimate same as minutes
			HyperFocusCompatability: 3,        // Default middle compatibility
		}
		// Description stays as-is since no JSON to remove
	} else {
		// Parse the JSON metadata
		var hyperfocusData models.HyperFocusMetadata
		if err := json.Unmarshal([]byte(meta), &hyperfocusData); err != nil {
			// JSON exists but invalid - use defaults
			enrichedTask.Metadata = &models.HyperFocusMetadata{
				Energy:                  "medium",
				Mode:                    "quick",
				Extend:                  false,
				Minutes:                 25,
				Estimate:                25,
				HyperFocusCompatability: 3,
			}
		} else {
			// Valid JSON metadata found
			enrichedTask.Metadata = &hyperfocusData

			// Validate and set defaults for missing fields
			if hyperfocusData.Energy == "" {
				enrichedTask.Metadata.Energy = "medium"
			}
			if hyperfocusData.Mode == "" {
				enrichedTask.Metadata.Mode = "quick"
			}
			if hyperfocusData.Minutes == 0 {
				enrichedTask.Metadata.Minutes = 25
			}
			if hyperfocusData.Estimate == 0 {
				// If no estimate provided, use minutes as estimate
				enrichedTask.Metadata.Estimate = enrichedTask.Metadata.Minutes
			}
			if hyperfocusData.HyperFocusCompatability == 0 {
				enrichedTask.Metadata.HyperFocusCompatability = 3 // Default middle
			}
		}

		// Clean the description by removing the JSON metadata
		enrichedTask.CleanDescription = strings.Replace(task.Description, meta, "", 1)
		enrichedTask.CleanDescription = strings.TrimSpace(enrichedTask.CleanDescription)
	}

	return enrichedTask, nil
}

func enrichTasks(tasks []models.RawTask) ([]models.Task, error) {
	enrichedTasks := make([]models.Task, 0, len(tasks))

	for _, task := range tasks {
		enriched, err := enrichTask(&task)
		if err != nil {
			// Log error but continue with others
			continue
		}
		enrichedTasks = append(enrichedTasks, *enriched)
	}

	return enrichedTasks, nil
}

// enrichMinimalTask enriches a MinimalTask with additional metadata or computed fields.
func enrichMinimalTask(task *models.MinimalTask) *models.MinimalTask {
	if task == nil {
		return nil
	}
	// Example: set a default priority if not set
	if task.Priority == 0 {
		task.Priority = 3 // default priority
	}
	// Example: add a stub metadata if missing
	if task.Metadata == nil {
		task.Metadata = &models.HyperFocusMetadata{
			Energy:                  "medium",
			Mode:                    "quick",
			Extend:                  false,
			Minutes:                 25,
			Estimate:                25,
			HyperFocusCompatability: 3,
		}
	}
	// Example: mark as high priority if title contains "urgent"
	if task.Title != "" && (containsIgnoreCase(task.Title, "urgent") || containsIgnoreCase(task.Description, "urgent")) {
		task.Priority = 5
	}
	return task
}

// containsIgnoreCase checks if substr is in s, case-insensitive
func containsIgnoreCase(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

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

// UpsertTask creates or updates a task.
func (s *Service) UpsertTask(ctx context.Context, task models.MinimalTask) (*models.MinimalTask, error) {
	enriched := enrichMinimalTask(&task)
	return s.Client.UpsertTask(ctx, *enriched)
}

// Me fetches the current authenticated user.
func (s *Service) Me(ctx context.Context) (*models.User, error) {
	var user models.User
	err := s.Client.Get(ctx, "/api/v1/user", &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
