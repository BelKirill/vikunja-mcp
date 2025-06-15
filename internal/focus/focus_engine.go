package focus

import (
	"github.com/BelKirill/vikunja-mcp/models"
	"github.com/charmbracelet/log"
)

// basicTaskFocus provides simple task filtering by focuswhen AI is unavailable
func (s *Service) basicTaskFocus(tasks []models.Task, opts *models.FocusOptions) []models.Task {
	log.Info("Using basic task filtering fallback")

	candidateTasks := []models.Task{}

	for _, t := range tasks {
		if t.Metadata == nil {
			log.Debug("Skipping task with missing metadata", "task_id", t.RawTask.ID)
			continue
		}
		if s.isTaskSuitable(&t, opts) {
			log.Debug("Task is suitable for focus", "task_id", t.RawTask.ID, "metadata", t.Metadata)
			candidateTasks = append(candidateTasks, t)
		}
	}

	log.Debug("Sorting candidate tasks by focus score", "count", len(candidateTasks))
	s.sortTasksByFocusScore(candidateTasks, opts)

	if opts.MaxTasks > 0 && len(candidateTasks) > opts.MaxTasks {
		log.Info("Limiting candidate tasks to MaxTasks", "max_tasks", opts.MaxTasks)
		candidateTasks = candidateTasks[:opts.MaxTasks]
	}

	log.Info("Basic filtering returning", "count", len(candidateTasks))
	return candidateTasks
}

// isTaskSuitable checks if a task is suitable for the current focus session
func (s *Service) isTaskSuitable(task *models.Task, opts *models.FocusOptions) bool {
	if task.Metadata == nil {
		return false
	}

	// Check energy level compatibility
	if !s.isEnergyCompatible(task.Metadata.Energy, opts.Energy) {
		return false
	}

	// Check mode compatibility
	if !s.isModeCompatible(task.Metadata.Mode, opts.Mode) {
		return false
	}

	// Check time constraints
	if !s.isTimeCompatible(task.Metadata, opts.MaxMinutes) {
		return false
	}

	return true
}

// isEnergyCompatible checks if task energy matches user energy
func (s *Service) isEnergyCompatible(taskEnergy, userEnergy string) bool {
	// Allow flexible matching - user with high energy can do low energy tasks
	energyLevels := map[string]int{
		"low": 1, "medium": 2, "high": 3, "social": 2,
	}

	taskLevel, taskOk := energyLevels[taskEnergy]
	userLevel, userOk := energyLevels[userEnergy]

	if !taskOk || !userOk {
		return true // Default to compatible if unknown
	}

	return userLevel >= taskLevel
}

// isModeCompatible checks if task mode matches user mode
func (s *Service) isModeCompatible(taskMode, userMode string) bool {
	// Exact match or flexible compatibility
	if taskMode == userMode {
		return true
	}

	// Quick mode can handle admin tasks
	if userMode == "quick" && taskMode == "admin" {
		return true
	}

	return false
}

// isTimeCompatible checks if task fits within available time
func (s *Service) isTimeCompatible(metadata *models.HyperFocusMetadata, maxMinutes int) bool {
	// Check if base time fits
	if metadata.Minutes <= maxMinutes {
		return true
	}

	// Check if extended time is worth it (at least 70% utilization)
	if metadata.Extend {
		return float64(maxMinutes)/float64(metadata.Minutes) >= 0.7
	}

	return false
}

// sortTasksByFocusScore sorts tasks by calculated focus suitability score
func (s *Service) sortTasksByFocusScore(tasks []models.Task, opts *models.FocusOptions) {
	// Calculate focus scores for each task
	for i := range tasks {
		tasks[i].FocusScore = s.calculateFocusScore(&tasks[i], opts)
	}

	// Sort by focus score (highest first)
	for i := 0; i < len(tasks)-1; i++ {
		for j := i + 1; j < len(tasks); j++ {
			if tasks[i].FocusScore < tasks[j].FocusScore {
				tasks[i], tasks[j] = tasks[j], tasks[i]
			}
		}
	}
}

// calculateFocusScore computes a suitability score for a task
func (s *Service) calculateFocusScore(task *models.Task, opts *models.FocusOptions) float64 {
	if task.Metadata == nil {
		return 0.1 // Very low score for tasks without metadata
	}

	score := 0.0

	// Priority weight (25%)
	priorityScore := float64(task.RawTask.Priority) / 5.0
	score += 0.25 * priorityScore

	// Energy match (25%)
	energyScore := s.getEnergyMatchScore(task.Metadata.Energy, opts.Energy)
	score += 0.25 * energyScore

	// Mode match (25%)
	modeScore := s.getModeMatchScore(task.Metadata.Mode, opts.Mode)
	score += 0.25 * modeScore

	// Time fit (25%)
	timeScore := s.getTimeFitScore(task.Metadata, opts.MaxMinutes)
	score += 0.25 * timeScore

	return score
}

// getEnergyMatchScore calculates energy compatibility score
func (s *Service) getEnergyMatchScore(taskEnergy, userEnergy string) float64 {
	if taskEnergy == userEnergy {
		return 1.0
	}

	// Partial compatibility scoring
	energyLevels := map[string]int{
		"low": 1, "medium": 2, "high": 3, "social": 2,
	}

	taskLevel, taskOk := energyLevels[taskEnergy]
	userLevel, userOk := energyLevels[userEnergy]

	if !taskOk || !userOk {
		return 0.5 // Neutral score for unknown
	}

	if userLevel >= taskLevel {
		return 0.8 // Good compatibility
	}

	return 0.3 // Poor compatibility
}

// getModeMatchScore calculates mode compatibility score
func (s *Service) getModeMatchScore(taskMode, userMode string) float64 {
	if taskMode == userMode {
		return 1.0
	}

	// Flexible compatibility
	if userMode == "quick" && taskMode == "admin" {
		return 0.8
	}

	if userMode == "deep" && taskMode == "admin" {
		return 0.6
	}

	return 0.2
}

// getTimeFitScore calculates how well task fits available time
func (s *Service) getTimeFitScore(metadata *models.HyperFocusMetadata, maxMinutes int) float64 {
	baseMinutes := metadata.Minutes

	if baseMinutes <= maxMinutes {
		// Perfect fit or room for extension
		if metadata.Extend && metadata.Estimate <= maxMinutes {
			return 1.0 // Can do full estimated work
		}
		return 0.9 // Good fit for base work
	}

	// Task requires more time than available
	ratio := float64(maxMinutes) / float64(baseMinutes)
	if ratio >= 0.7 {
		return 0.6 // Acceptable partial completion
	}

	return 0.2 // Poor time fit
}
