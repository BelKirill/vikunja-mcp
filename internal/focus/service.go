package focus

import (
	"context"
	"os"

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
		APIKey: os.Getenv("OPENAI_API_KEY"),
		Model:  getEnvOrDefault("OPENAI_MODEL", "gpt-4"),
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

func (s Service) GetFilteredTasks(ctx context.Context, filter string) ([]models.Task, error) {
	log.Info("GetFilteredTasks called", "filter", filter)

	// Use AI-powered filter engine for intelligent task filtering
	newFilter, err := s.FocusEngine.SuggestFilter(ctx, &filter)
	if err != nil {
		log.Warn("Filter engine failed", "error", err)

		tasks, err := s.Vikunja.GetFilteredTasks(ctx, &filter)
		if err != nil {
			log.Error("Basic filtering failed", "filter", filter, "error", err)
			return nil, err
		}

		return tasks, nil
	}

	tasks, err := s.Vikunja.GetFilteredTasks(ctx, &newFilter.Filter)
	if err != nil {
		log.Error("AI suggested filtering failed", "filter", newFilter, "error", err)
		return nil, err
	}

	return tasks, nil
}

// GetFocusTasks returns AI-ranked tasks suitable for focus sessions
func (s *Service) GetFocusTasks(ctx context.Context, opts *models.FocusOptions) ([]models.Task, error) {
	log.Info("GetFocusTasks called with AI engine", "opts", opts)

	// Get all incomplete tasks from Vikunja
	tasks, err := s.Vikunja.GetIncompleteTasks(ctx)
	if err != nil {
		log.Error("Failed to get incomplete tasks from Vikunja", "error", err)
		return nil, err
	}
	log.Info("Fetched incomplete tasks", "count", len(tasks))

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

// GetTaskRecommendation returns the single best task with AI-powered reasoning
func (s *Service) GetTaskRecommendation(ctx context.Context, opts *models.FocusOptions) (*models.Task, error) {
	log.Info("GetTaskRecommendation called with AI engine", "opts", opts)

	// Get all incomplete tasks
	tasks, err := s.Vikunja.GetIncompleteTasks(ctx)
	if err != nil {
		log.Error("Failed to get incomplete tasks", "error", err)
		return nil, err
	}

	// Use AI-powered recommendation
	recommendation, err := s.FocusEngine.GetTaskRecommendation(ctx, tasks, opts)
	if err != nil {
		log.Error("AI recommendation failed", "error", err)
		// Fallback to basic recommendation
		return s.basicTaskRecommendation(tasks, opts), nil
	}

	if recommendation == nil || recommendation.Task == nil {
		log.Info("No suitable tasks found for recommendation")
		return nil, nil
	}

	log.Info("Returning AI-powered task recommendation",
		"task_id", recommendation.Task.Task.RawTask.ID,
		"strategy", recommendation.SessionStrategy,
		"length", recommendation.RecommendedLength,
		"reasoning", recommendation.Reasoning)

	return &recommendation.Task.Task, nil
}

// GetEnhancedTaskRecommendation returns the full AI recommendation with reasoning
func (s *Service) GetEnhancedTaskRecommendation(ctx context.Context, opts *models.FocusOptions) (*models.TaskRecommendation, error) {
	log.Info("GetEnhancedTaskRecommendation called")

	tasks, err := s.Vikunja.GetIncompleteTasks(ctx)
	if err != nil {
		log.Error("Failed to get incomplete tasks", "error", err)
		return nil, err
	}

	recommendation, err := s.FocusEngine.GetTaskRecommendation(ctx, tasks, opts)
	if err != nil {
		log.Error("Enhanced recommendation failed", "error", err)
		return nil, err
	}

	return recommendation, nil
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
	return s.Vikunja.UpsertTask(ctx, *task)
}

// GetTaskMetadata retrieves detailed metadata for a specific task
func (s *Service) GetTaskMetadata(ctx context.Context, taskID int64) (*models.Task, error) {
	log.Info("GetTaskMetadata called", "task_id", taskID)
	log.Debug("Fetching task metadata from Vikunja", "task_id", taskID)
	task, err := s.Vikunja.GetTaskByID(ctx, taskID)
	if err != nil {
		log.Error("Failed to get task by ID in GetTaskMetadata", "task_id", taskID, "error", err)
		return nil, err
	}
	log.Info("Returning task metadata", "task_id", taskID)
	log.Debug("Task metadata result", "task", task)
	return task, nil
}

// =============================================================================
// Fallback Methods (for when AI is unavailable)
// =============================================================================

// =============================================================================
// Utility Functions
// =============================================================================

// getEnvOrDefault returns environment variable value or default
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// min returns the smaller of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
