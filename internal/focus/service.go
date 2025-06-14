package focus

import (
	"context"
	"sort"

	vikunja "github.com/BelKirill/vikunja-mcp/internal/vikunja"
	"github.com/BelKirill/vikunja-mcp/models"
	"github.com/charmbracelet/log"
)

// Service provides focus-specific business logic for task management
type Service struct {
	Vikunja *vikunja.Service
}

// NewService creates a new Focus Service with Vikunja dependency
func NewService() (*Service, error) {
	log.Info("NewService (focus) called")
	vikunjaSvc, err := vikunja.NewService()
	if err != nil {
		log.Error("Failed to create Vikunja service", "error", err)
		return nil, err
	}
	log.Info("Vikunja service created successfully for focus.Service")
	return &Service{Vikunja: vikunjaSvc}, nil
}

// GetFocusTasks returns tasks suitable for focus sessions based on energy and mode
func (s *Service) GetFocusTasks(ctx context.Context, opts models.FocusOptions) ([]models.FocusResult, error) {
	log.Info("GetFocusTasks called", "opts", opts)
	tasks, err := s.Vikunja.GetIncompleteTasks(ctx)
	if err != nil {
		log.Error("Failed to get incomplete tasks from Vikunja", "error", err)
		return nil, err
	}
	log.Info("Fetched incomplete tasks", "count", len(tasks))
	var candidateTasks []models.FocusResult
	for _, t := range tasks {
		if t.Metadata == nil {
			log.Debug("Skipping task with missing metadata", "task_id", t.RawTask.ID)
			continue
		}
		if s.isTaskSuitable(t, opts) {
			log.Debug("Task is suitable for focus", "task_id", t.RawTask.ID, "metadata", t.Metadata)
			result := models.FocusResult{
				TaskID:      t.RawTask.ID,
				Project:     t.RawTask.ProjectID,
				Metadata:    t.Metadata,
				Priority:    t.RawTask.Priority,
				Title:       t.RawTask.Title,
				Done:        t.RawTask.Done,
				Description: t.CleanDescription,
				Estimate:    t.Metadata.Estimate,
			}
			candidateTasks = append(candidateTasks, result)
		}
	}
	log.Debug("Sorting candidate tasks by focus score", "count", len(candidateTasks))
	s.sortTasksByFocusScore(candidateTasks, opts)
	if opts.MaxTasks > 0 && len(candidateTasks) > opts.MaxTasks {
		log.Info("Limiting candidate tasks to MaxTasks", "max_tasks", opts.MaxTasks)
		candidateTasks = candidateTasks[:opts.MaxTasks]
	}
	log.Info("GetFocusTasks returning", "count", len(candidateTasks))
	return candidateTasks, nil
}

// isTaskSuitable determines if a task matches the focus criteria
func (s *Service) isTaskSuitable(task models.Task, opts models.FocusOptions) bool {
	meta := task.Metadata

	// Check hyperfocus compatibility threshold
	if meta.HyperFocusCompatability < 3 {
		log.Debug("Task below hyperfocus compatibility threshold", "task_id", task.RawTask.ID)
		return false
	}

	// Energy level matching with flexibility
	energyMatch := s.isEnergyCompatible(meta.Energy, opts.Energy)

	// Mode matching (exact match preferred)
	modeMatch := meta.Mode == opts.Mode

	// Time constraint matching (if specified)
	timeMatch := true
	if opts.MaxMinutes > 0 {
		timeMatch = meta.Estimate <= opts.MaxMinutes
	}

	if !energyMatch || !modeMatch || !timeMatch {
		log.Debug("Task not suitable for focus", "task_id", task.RawTask.ID, "energyMatch", energyMatch, "modeMatch", modeMatch, "timeMatch", timeMatch)
	}
	return energyMatch && modeMatch && timeMatch
}

// isEnergyCompatible checks if task energy level matches user's current energy
func (s *Service) isEnergyCompatible(taskEnergy, userEnergy string) bool {
	log.Debug("isEnergyCompatible called", "taskEnergy", taskEnergy, "userEnergy", userEnergy)

	// Exact match is always good
	if taskEnergy == userEnergy {
		return true
	}

	// Medium energy can handle low or high energy tasks (flexibility)
	if userEnergy == "medium" && (taskEnergy == "low" || taskEnergy == "high") {
		return true
	}

	// High energy can handle medium tasks
	if userEnergy == "high" && taskEnergy == "medium" {
		return true
	}

	// Social energy is separate - no cross-compatibility
	return false
}

// sortTasksByFocusScore sorts tasks by their suitability for focus sessions
func (s *Service) sortTasksByFocusScore(tasks []models.FocusResult, opts models.FocusOptions) {
	log.Debug("Sorting tasks by focus score", "count", len(tasks))
	sort.Slice(tasks, func(i, j int) bool {
		scoreI := s.calculateFocusScore(tasks[i], opts)
		scoreJ := s.calculateFocusScore(tasks[j], opts)
		log.Debug("Comparing focus scores", "i", i, "scoreI", scoreI, "j", j, "scoreJ", scoreJ)
		return scoreI > scoreJ // Higher score first
	})
}

// calculateFocusScore computes a focus suitability score for a task
func (s *Service) calculateFocusScore(task models.FocusResult, opts models.FocusOptions) float64 {
	var score float64

	// Base score from hyperfocus compatibility (1-5)
	score = float64(task.Metadata.HyperFocusCompatability)

	// Bonus for exact energy match
	if task.Metadata.Energy == opts.Energy {
		score += 2.0
	}

	// Bonus for exact mode match
	if task.Metadata.Mode == opts.Mode {
		score += 2.0
	}

	// Bonus for extendable tasks if user has time
	if task.Metadata.Extend && opts.MaxMinutes >= 50 {
		score += 1.0
	}

	// Priority weighting (Vikunja priority 1-5)
	score += float64(task.Priority) * 0.5

	// Time preference: prefer tasks that fit well in available time
	if opts.MaxMinutes > 0 {
		timeRatio := float64(task.Metadata.Estimate) / float64(opts.MaxMinutes)
		if timeRatio >= 0.5 && timeRatio <= 1.0 {
			log.Debug("Good time fit for task", "task_id", task.TaskID, "timeRatio", timeRatio)
			score += 1.0 // Good time fit
		}
	}
	log.Debug("Calculated focus score", "task_id", task.TaskID, "score", score)
	return score
}

// GetTaskRecommendation returns the single best task for a focus session
func (s *Service) GetTaskRecommendation(ctx context.Context, opts models.FocusOptions) (*models.FocusResult, error) {
	log.Info("GetTaskRecommendation called", "opts", opts)
	// Limit to 1 task
	opts.MaxTasks = 1

	tasks, err := s.GetFocusTasks(ctx, opts)
	if err != nil {
		log.Error("Failed to get focus tasks for recommendation", "error", err)
		return nil, err
	}

	if len(tasks) == 0 {
		log.Info("No suitable tasks found for recommendation")
		return nil, nil // No suitable tasks found
	}
	log.Info("Returning top recommended task", "task_id", tasks[0].TaskID)
	return &tasks[0], nil
}

// EstimateSessionLength calculates optimal session length for a task
func (s *Service) EstimateSessionLength(task models.FocusResult, userMaxMinutes int) int {
	log.Info("EstimateSessionLength called", "task_id", task.TaskID, "userMaxMinutes", userMaxMinutes)
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
	log.Debug("User has less time than base pomodoro", "userMaxMinutes", userMaxMinutes)
	return userMaxMinutes
}

// UpsertTask creates or updates a task through the Vikunja service
func (s *Service) UpsertTask(ctx context.Context, task models.MinimalTask) (*models.MinimalTask, error) {
	log.Info("UpsertTask called in focus.Service", "task_id", task.TaskID)
	log.Debug("UpsertTask details", "task", task)
	return s.Vikunja.UpsertTask(ctx, task)
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
