// Package engine provides AI-powered task selection for focus sessions
package engine

import (
	"context"

	"github.com/BelKirill/vikunja-mcp/models"
)

// =============================================================================
// Contextual Filters
// =============================================================================

// TimeConstraintFilter filters tasks based on available time
type TimeConstraintFilter struct{}

func NewTimeConstraintFilter() *TimeConstraintFilter {
	return &TimeConstraintFilter{}
}

func (f *TimeConstraintFilter) Name() string {
	return "time_constraint"
}

func (f *TimeConstraintFilter) Filter(ctx context.Context, tasks []models.Task, opts *models.FocusOptions) []models.Task {
	var filtered []models.Task

	for _, task := range tasks {
		if task.Metadata == nil {
			continue
		}

		// Skip tasks that require more time than available
		minTime := task.Metadata.Minutes
		if task.Metadata.Extend && task.Metadata.Estimate > minTime {
			minTime = task.Metadata.Estimate
		}

		if minTime <= opts.MaxMinutes {
			filtered = append(filtered, task)
		}
	}

	return filtered
}

// PriorityConstraintFilter filters tasks based on priority
type PriorityConstraintFilter struct{}

func NewPriorityConstraintFilter() *PriorityConstraintFilter {
	return &PriorityConstraintFilter{}
}

func (f *PriorityConstraintFilter) Name() string {
	return "priority_constraint"
}

func (f *PriorityConstraintFilter) Filter(ctx context.Context, tasks []models.Task, opts *models.FocusOptions) []models.Task {
	if len(tasks) == 0 {
		return tasks
	}

	// Group tasks by priority level
	priorityGroups := make(map[int][]models.Task)
	for _, task := range tasks {
		priority := task.RawTask.Priority
		if priority == 0 {
			priority = 3 // Default to medium priority if not set
		}
		priorityGroups[priority] = append(priorityGroups[priority], task)
	}

	var selected []models.Task
	remainingMinutes := opts.MaxMinutes

	// Priority cascade: 5 → 4 → 3 → 2 → 1
	for priority := 5; priority >= 1 && remainingMinutes > 0; priority-- {
		priorityTasks, exists := priorityGroups[priority]
		if !exists {
			continue
		}

		// Add tasks from this priority level until time budget exhausted
		for _, task := range priorityTasks {
			estimatedMinutes := f.getTaskEstimate(task)

			// Always include high-priority tasks even if they exceed remaining time
			if priority >= 4 || estimatedMinutes <= remainingMinutes {
				selected = append(selected, task)
				remainingMinutes -= estimatedMinutes

				// Prevent negative remaining time for display purposes
				if remainingMinutes < 0 {
					remainingMinutes = 0
				}
			}
		}
	}

	return selected
}

// getTaskEstimate returns the estimated time for a task in minutes
func (f *PriorityConstraintFilter) getTaskEstimate(task models.Task) int {
	if task.Metadata == nil {
		return 25 // Default Pomodoro
	}

	// Use explicit estimate if available
	if task.Metadata.Estimate > 0 {
		return task.Metadata.Estimate
	}

	// Fall back to AI-suggested minutes
	if task.Metadata.Minutes > 0 {
		return task.Metadata.Minutes
	}

	// Default to one Pomodoro
	return 25
}

// EnergyModeFilter filters tasks based on energy level and mode compatibility
type EnergyModeFilter struct{}

func NewEnergyModeFilter() *EnergyModeFilter {
	return &EnergyModeFilter{}
}

func (f *EnergyModeFilter) Name() string {
	return "energy_mode"
}

func (f *EnergyModeFilter) Filter(ctx context.Context, tasks []models.Task, opts *models.FocusOptions) []models.Task {
	var filtered []models.Task

	for _, task := range tasks {
		if task.Metadata == nil {
			continue
		}

		// Filter by energy compatibility
		if f.isEnergyCompatible(task.Metadata.Energy, opts.Energy) &&
			f.isModeCompatible(task.Metadata.Mode, opts.Mode) {
			filtered = append(filtered, task)
		}
	}

	return filtered
}

func (f *EnergyModeFilter) isEnergyCompatible(taskEnergy, userEnergy string) bool {
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

func (f *EnergyModeFilter) isModeCompatible(taskMode, userMode string) bool {
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

// DependencyFilter filters out blocked tasks
type DependencyFilter struct{}

func NewDependencyFilter() *DependencyFilter {
	return &DependencyFilter{}
}

func (f *DependencyFilter) Name() string {
	return "dependency"
}

func (f *DependencyFilter) Filter(ctx context.Context, tasks []models.Task, opts *models.FocusOptions) []models.Task {
	// For now, just return all tasks
	// In the future, this would check for:
	// - Blocked by other incomplete tasks
	// - Dependencies on external factors
	// - Project status
	return tasks
}
