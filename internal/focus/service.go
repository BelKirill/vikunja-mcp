package focus

import (
	"context"
	"sync"

	"github.com/BelKirill/vikunja-mcp/internal/config"
	"github.com/BelKirill/vikunja-mcp/internal/engine"
	"github.com/BelKirill/vikunja-mcp/internal/openai"
	vikunja "github.com/BelKirill/vikunja-mcp/internal/vikunja"
	"github.com/BelKirill/vikunja-mcp/models"
	"github.com/charmbracelet/log"
)

// Service provides focus-specific business logic for task management
type Service struct {
	Vikunja     *vikunja.Service
	FocusEngine *engine.FocusEngine
}

// NewService creates a new Focus Service with Vikunja dependency and AI engine
func NewService() (*Service, error) {
	log.Info("NewService (focus) called")

	// Initialize Vikunja service
	vikunjaSvc, err := vikunja.NewService()
	if err != nil {
		log.Error("Failed to create Vikunja service", "error", err)
		return nil, err
	}
	log.Info("Vikunja service created successfully for focus.Service")

	// Initialize AI-powered focus engine
	focusEngine, err := initializeFocusEngine()
	if err != nil {
		log.Error("Failed to create focus engine", "error", err)
		return nil, err
	}
	log.Info("Focus engine created successfully")

	return &Service{
		Vikunja:     vikunjaSvc,
		FocusEngine: focusEngine,
	}, nil
}

// initializeFocusEngine creates and configures the focus engine with AI decision making
func initializeFocusEngine() (*engine.FocusEngine, error) {
	// Configure OpenAI decision engine
	openaiConfig := openai.OpenAIConfig{
		APIKey: config.GetOpenAI().APIKey,
		Model:  config.GetOpenAI().Model,
	}

	if openaiConfig.APIKey == "" {
		log.Warn("OPENAI_API_KEY not set, focus engine will use heuristic fallback only")
	}

	decisionEngine := openai.NewOpenAIDecisionEngine(openaiConfig)

	// Create focus engine with AI decision making and heuristic fallback
	focusEngine := engine.NewFocusEngine(
		decisionEngine,
		engine.WithLearning(true), // Enable learning from past sessions
	)

	return focusEngine, nil
}

func (s Service) GetFilteredTasks(ctx context.Context, filter string, useAI bool) ([]models.RawTask, error) {
	log.Info("GetFilteredTasks called", "filter", filter, "useAI", useAI)

	finalFilter := filter

	// Only use AI if explicitly requested
	if useAI {
		log.Debug("Using AI to enhance filter", "original_filter", filter)
		newFilter, err := s.FocusEngine.SuggestFilter(ctx, &filter)
		if err != nil {
			log.Warn("Filter engine failed, using original filter", "error", err, "original_filter", filter)
		} else {
			finalFilter = newFilter.Filter
			log.Debug("AI enhanced filter", "original", filter, "enhanced", finalFilter)
		}
	} else {
		log.Debug("Using filter expression directly", "filter", filter)
	}

	tasks, err := s.Vikunja.GetFilteredTasks(ctx, &finalFilter)
	if err != nil {
		log.Error("Filtering failed", "filter", finalFilter, "error", err)
		return nil, err
	}

	log.Info("Successfully filtered tasks", "count", len(tasks), "filter", finalFilter)
	return tasks, nil
}

// GetFocusTasks returns AI-ranked tasks suitable for focus sessions
func (s *Service) GetFocusTasks(ctx context.Context, opts *models.FocusOptions) ([]models.Task, error) {
	log.Info("GetFocusTasks called with AI engine", "opts", opts)

	// Get all incomplete tasks from Vikunja
	rawTasks, err := s.Vikunja.GetIncompleteTasks(ctx)
	if err != nil {
		log.Error("Failed to get incomplete tasks from Vikunja", "error", err)
		return nil, err
	}
	log.Info("Fetched incomplete tasks", "count", len(rawTasks))

	log.Info("Sending to project filter", "exclude project", opts.ExcludeProjects, "only_projects", opts.OnlyProjects)
	rawTasksFiltered := filterTasksByProjects(rawTasks, opts.ExcludeProjects, opts.OnlyProjects)

	tasks := s.EnrichTasksParallel(ctx, rawTasksFiltered)

	// Use AI-powered focus engine for intelligent task selection
	decision, err := s.FocusEngine.GetFocusTasks(ctx, tasks, opts)
	if err != nil {
		log.Error("Focus engine failed", "error", err)
		// Fallback to basic filtering if AI fails
		return s.basicTaskFocus(tasks, opts), nil
	}

	// Convert ranked tasks back to models.Task slice
	result := make([]models.Task, len(decision.RankedTasks))
	for i, rankedTask := range decision.RankedTasks {
		result[i] = rankedTask.Task
	}

	log.Info("GetFocusTasks returning AI-ranked tasks",
		"count", len(result),
		"strategy", decision.Strategy,
		"confidence", decision.Confidence,
		"reasoning", decision.Reasoning)

	return result, nil
}

// EstimateSessionLength calculates optimal session length using AI insights
func (s *Service) EstimateSessionLength(task *models.FocusResult, userMaxMinutes int) int {
	log.Info("EstimateSessionLength called", "task_id", task.TaskID, "userMaxMinutes", userMaxMinutes)

	// If we have AI metadata, use it
	if task.Metadata != nil {
		estimate := task.Metadata.Estimate
		minutes := task.Metadata.Minutes

		// If task can extend and user has time, suggest longer session
		if task.Metadata.Extend && userMaxMinutes >= estimate {
			log.Debug("Suggesting extended session length", "estimate", estimate)
			return estimate
		}

		// Otherwise use base pomodoro length
		if userMaxMinutes >= minutes {
			log.Debug("Suggesting base pomodoro length", "minutes", minutes)
			return minutes
		}
	}

	// Default fallback
	log.Debug("Using default session length", "userMaxMinutes", userMaxMinutes)
	return min(25, userMaxMinutes) // Default 25-minute pomodoro
}

// UpsertTask creates or updates a task through the Vikunja service
func (s *Service) UpsertTask(ctx context.Context, task *models.RawTask) (*models.RawTask, error) {
	log.Info("UpsertTask called in focus.Service", "task_id", task.ID)
	log.Debug("UpsertTask details", "task", task)
	return s.Vikunja.UpsertTask(ctx, task)
}

// GetFullTaskData retrieves detailed data for a specific task
func (s *Service) GetFullTaskData(ctx context.Context, taskID int64) (*models.Task, []models.Comment, error) {
	log.Info("GetFullTaskData called", "task_id", taskID)
	availableLabels := []models.PartialLabel{}
	vikunjaLabels, err := s.Vikunja.GetAllLabels(ctx)
	if err != nil {
		log.Warn("Error in bringing available labels", "error", err)
	} else {
		availableLabels = vikunjaLabels
	}

	log.Debug("Fetching task data from Vikunja", "task_id", taskID)
	rawTask, err := s.Vikunja.GetTaskByID(ctx, taskID)
	if err != nil {
		log.Error("Failed to get task by ID in GetFullTaskData", "task_id", taskID, "error", err)
		return nil, nil, err
	}

	task, enriched, err := s.FocusEngine.EnrichTask(ctx, rawTask, availableLabels)
	if err != nil {
		log.Warn("Failed to enrich task in GetFullTaskData", "task_id", taskID, "error", err)
		task = &models.Task{
			Identifier:       rawTask.Identifier,
			RawTask:          rawTask,
			CleanDescription: rawTask.Description,
		}
	}

	if enriched {
		updated, err := s.Vikunja.UpsertTask(ctx, task.RawTask)
		if err != nil {
			log.Warn("Could not save enriched data", "description", task.RawTask.Description, "err", err)
		} else {
			task.RawTask.Description = updated.Description
		}
	}
	if len(task.RawTask.Labels) > 0 {
		createdLabels, err := s.Vikunja.AddLabels(ctx, task.RawTask.ID, task.RawTask.Created, task.RawTask.Labels)
		if err != nil {
			log.Warn("Could not save label data", "labels", task.RawTask.Labels, "err", err)
		} else {
			task.RawTask.Labels = createdLabels
		}
	}

	log.Debug("Fetching comments from Vikunja", "task_id", taskID)
	comments, err := s.Vikunja.GetTaskCommentsByID(ctx, taskID)
	if err != nil {
		log.Error("Failed to get task comments by ID in GetTaskCommentsByID", "task_id", taskID, "error", err)
		return task, nil, err
	}

	log.Info("Returning task metadata", "task_id", taskID)
	log.Debug("Task data result", "task", task, "comments", comments)
	return task, comments, nil
}

func (s *Service) EnrichTasksParallel(ctx context.Context, tasks []models.RawTask) []models.Task {
	enriched := make([]models.Task, len(tasks))
	availableLabels := []models.PartialLabel{}
	vikunjaLabels, err := s.Vikunja.GetAllLabels(ctx)
	if err != nil {
		log.Warn("Error in bringing available labels", "error", err)
	} else {
		availableLabels = vikunjaLabels
	}

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 10) // Limit concurrent calls

	for i, task := range tasks {
		wg.Add(1)
		go func(index int, t models.RawTask) {
			defer wg.Done()
			semaphore <- struct{}{}        // Acquire
			defer func() { <-semaphore }() // Release

			enrichedTask, upsert, err := s.FocusEngine.EnrichTask(ctx, &t, availableLabels)
			if err != nil {
				log.Warn("Enrich task failed", "taskID", t.ID)
				structuredTask := models.Task{
					Identifier:       task.Identifier,
					CleanDescription: task.Description,
					RawTask:          &task,
				}
				enriched[index] = structuredTask
			} else {
				if upsert {
					updated, err := s.Vikunja.UpsertTask(ctx, enrichedTask.RawTask)
					if err != nil {
						log.Error("Failed to upsert enriched task", "error", err, "task_id", enrichedTask.RawTask.ID)
					} else {
						enrichedTask.RawTask.Description = updated.Description
					}
					newLabels, err := s.Vikunja.AddLabels(ctx, enrichedTask.RawTask.ID, enrichedTask.RawTask.Created, enrichedTask.RawTask.Labels)
					if err != nil {
						log.Error("Failed to add labels", "error", err)
					} else {
						enrichedTask.RawTask.Labels = newLabels
					}
				}
				enriched[index] = *enrichedTask
			}
		}(i, task)
	}

	wg.Wait()
	return enriched
}

// AddComment adds a comment to a specific task
func (s *Service) AddComment(ctx context.Context, taskID int64, comment *string) (*models.Comment, error) {
	newComment, err := s.Vikunja.AddComment(ctx, taskID, comment)
	if err != nil {
		return nil, err
	}
	return newComment, nil
}

// =============================================================================
// Fallback Methods (for when AI is unavailable)
// =============================================================================

// =============================================================================
// Utility Functions
// =============================================================================

// min returns the smaller of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func filterTasksByProjects(tasks []models.RawTask, excludeProjects []int64, onlyProjects []int64) []models.RawTask {
	log.Info("filterTasksByProjects called", "total_tasks", len(tasks), "excludeProjects", excludeProjects, "onlyProjects", onlyProjects)
	filtered := make([]models.RawTask, 0, len(tasks))

	for _, task := range tasks {
		projectID := task.ProjectID
		// Skip excluded projects
		if contains(excludeProjects, projectID) {
			log.Info("Excluding task due to excludeProjects", "task_id", task.ID, "project_id", projectID)
			continue
		}

		// If onlyProjects specified, task must be in that list
		if len(onlyProjects) > 0 && !contains(onlyProjects, projectID) {
			log.Info("Excluding task not in onlyProjects", "task_id", task.ID, "project_id", projectID)
			continue
		}

		log.Info("Including task", "task_id", task.ID, "project_id", projectID)
		filtered = append(filtered, task)
	}

	log.Info("filterTasksByProjects result", "filtered_count", len(filtered))
	return filtered
}

func contains(slice []int64, item int64) bool {
	log.Debug("contains called", "slice", slice, "item", item)
	for _, v := range slice {
		if v == item {
			log.Debug("Item found in slice", "item", item)
			return true
		}
	}
	log.Debug("Item not found in slice", "item", item)
	return false
}
