package focus

import (
	"context"
	"sort"

	vikunja "github.com/BelKirill/vikunja-mcp/internal/vikunja"
	"github.com/BelKirill/vikunja-mcp/models"
)

// Service provides focus-specific business logic for task management
type Service struct {
	Vikunja *vikunja.Service
}

// NewService creates a new Focus Service with Vikunja dependency
func NewService() (*Service, error) {
	vikunjaSvc, err := vikunja.NewService()
	if err != nil {
		return nil, err
	}
	return &Service{Vikunja: vikunjaSvc}, nil
}

// GetFocusTasks returns tasks suitable for focus sessions based on energy and mode
func (s *Service) GetFocusTasks(ctx context.Context, opts models.FocusOptions) ([]models.FocusResult, error) {
	// Get enriched tasks from Vikunja service
	tasks, err := s.Vikunja.GetIncompleteTasks(ctx)
	if err != nil {
		return nil, err
	}

	// Convert to FocusResult and apply filtering
	var candidateTasks []models.FocusResult
	for _, t := range tasks {
		// Skip if metadata is missing
		if t.Metadata == nil {
			continue
		}

		// Apply energy and mode filtering
		if s.isTaskSuitable(t, opts) {
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

	// Sort by focus suitability score
	s.sortTasksByFocusScore(candidateTasks, opts)

	// Limit results if specified
	if opts.MaxTasks > 0 && len(candidateTasks) > opts.MaxTasks {
		candidateTasks = candidateTasks[:opts.MaxTasks]
	}

	return candidateTasks, nil
}

// isTaskSuitable determines if a task matches the focus criteria
func (s *Service) isTaskSuitable(task models.Task, opts models.FocusOptions) bool {
	meta := task.Metadata

	// Check hyperfocus compatibility threshold
	if meta.HyperFocusCompatability < 3 {
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

	return energyMatch && modeMatch && timeMatch
}

// isEnergyCompatible checks if task energy level matches user's current energy
func (s *Service) isEnergyCompatible(taskEnergy, userEnergy string) bool {
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
	sort.Slice(tasks, func(i, j int) bool {
		scoreI := s.calculateFocusScore(tasks[i], opts)
		scoreJ := s.calculateFocusScore(tasks[j], opts)
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
			score += 1.0 // Good time fit
		}
	}

	return score
}

// GetTaskRecommendation returns the single best task for a focus session
func (s *Service) GetTaskRecommendation(ctx context.Context, opts models.FocusOptions) (*models.FocusResult, error) {
	// Limit to 1 task
	opts.MaxTasks = 1

	tasks, err := s.GetFocusTasks(ctx, opts)
	if err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return nil, nil // No suitable tasks found
	}

	return &tasks[0], nil
}

// EstimateSessionLength calculates optimal session length for a task
func (s *Service) EstimateSessionLength(task models.FocusResult, userMaxMinutes int) int {
	estimate := task.Metadata.Estimate
	minutes := task.Metadata.Minutes

	// If task can extend and user has time, suggest longer session
	if task.Metadata.Extend && userMaxMinutes >= estimate {
		return estimate
	}

	// Otherwise use base pomodoro length
	if userMaxMinutes >= minutes {
		return minutes
	}

	// User doesn't have enough time for even base pomodoro
	return userMaxMinutes
}

// UpsertTask creates or updates a task through the Vikunja service
func (s *Service) UpsertTask(ctx context.Context, task models.MinimalTask) (*models.MinimalTask, error) {
	return s.Vikunja.UpsertTask(ctx, task)
}

// GetTaskMetadata retrieves detailed metadata for a specific task
func (s *Service) GetTaskMetadata(ctx context.Context, taskID int64) (*models.Task, error) {
	// Get the minimal task first
	minTask, err := s.Vikunja.GetTaskByID(ctx, taskID)
	if err != nil {
		return nil, err
	}

	// Convert to enriched task format
	// This might need adjustment based on your actual models
	enrichedTask := &models.Task{
		RawTask: &models.RawTask{
			ID:          minTask.TaskID,
			Title:       minTask.Title,
			Description: minTask.Description,
			Priority:    minTask.Priority,
			Done:        minTask.Done,
			ProjectID:   minTask.Project,
		},
		Metadata:         minTask.Metadata,
		CleanDescription: minTask.Description, // This would be cleaned by Vikunja service
	}

	return enrichedTask, nil
}
