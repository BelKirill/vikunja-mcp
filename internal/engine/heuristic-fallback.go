package engine

import (
	"context"
	"fmt"
	"sort"

	"github.com/BelKirill/vikunja-mcp/models"
	"github.com/charmbracelet/log"
)

// HeuristicFallback provides deterministic task selection using business rules
type HeuristicFallback struct{}

func NewHeuristicFallback() HeuristicFallback {
	return HeuristicFallback{}
}

func (h HeuristicFallback) SelectTasks(ctx context.Context, tasks []models.Task, opts *models.FocusOptions) []models.RankedTask {
	log.Info("HeuristicFallback.SelectTasks called", "task_count", len(tasks), "opts", opts)
	var ranked []models.RankedTask

	for _, task := range tasks {
		score := h.calculateHeuristicScore(task, opts)
		log.Debug("Heuristic score calculated", "task_id", task.RawTask.ID, "score", score)
		ranked = append(ranked, models.RankedTask{
			Task:             task,
			Score:            score,
			ReasoningFactors: h.getReasoningFactors(task, opts),
			EstimatedLength:  h.estimateSessionLength(task, opts.MaxMinutes),
			CanExtend:        task.Metadata != nil && task.Metadata.Extend,
		})
	}

	log.Debug("Sorting ranked tasks by score (heuristic)")
	// Sort by score
	sort.Slice(ranked, func(i, j int) bool {
		return ranked[i].Score > ranked[j].Score
	})

	if opts.MaxTasks > 0 && len(ranked) > opts.MaxTasks {
		log.Debug("Trimming ranked tasks to max_tasks (heuristic)", "before", len(ranked), "after", opts.MaxTasks)
		ranked = ranked[:opts.MaxTasks]
	}

	return ranked
}

func (h HeuristicFallback) GetRecommendation(ctx context.Context, tasks []models.Task, opts *models.FocusOptions) *models.TaskRecommendation {
	log.Info("HeuristicFallback.GetRecommendation called", "task_count", len(tasks), "opts", opts)
	ranked := h.SelectTasks(ctx, tasks, opts)

	if len(ranked) == 0 {
		log.Info("No ranked tasks available for recommendation")
		return nil
	}

	best := ranked[0]
	var alternatives []models.RankedTask
	if len(ranked) > 1 {
		alternatives = ranked[1:min(4, len(ranked))] // Up to 3 alternatives
	}

	log.Info("Returning best task recommendation", "task_id", best.Task.RawTask.ID, "alternatives_count", len(alternatives))
	return &models.TaskRecommendation{
		Task:              &best,
		RecommendedLength: best.EstimatedLength,
		SessionStrategy:   h.getSessionStrategy(best.Task),
		Reasoning:         fmt.Sprintf("Selected based on heuristic scoring: %s", joinStrings(best.ReasoningFactors, ", ")),
		Alternatives:      alternatives,
	}
}

func (h HeuristicFallback) calculateHeuristicScore(task models.Task, opts *models.FocusOptions) float64 {
	score := 0.0

	if task.Metadata == nil {
		return 0.1 // Very low score for tasks without metadata
	}

	// Priority weight (25%)
	priorityScore := float64(task.RawTask.Priority) / 5.0
	score += 0.25 * priorityScore

	// Energy match (25%)
	energyScore := h.getEnergyMatchScore(task.Metadata.Energy, opts.Energy)
	score += 0.25 * energyScore

	// Mode match (25%)
	modeScore := h.getModeMatchScore(task.Metadata.Mode, opts.Mode)
	score += 0.25 * modeScore

	// Time fit (25%)
	timeScore := h.getTimeFitScore(task.Metadata, opts.MaxMinutes)
	score += 0.25 * timeScore

	return score
}

func (h HeuristicFallback) getReasoningFactors(task models.Task, opts *models.FocusOptions) []string {
	var factors []string

	if task.RawTask.Priority >= 4 {
		factors = append(factors, "high priority")
	}

	if task.Metadata != nil {
		if task.Metadata.Energy == opts.Energy {
			factors = append(factors, "perfect energy match")
		}

		if task.Metadata.Mode == opts.Mode {
			factors = append(factors, "mode alignment")
		}

		if task.Metadata.Minutes <= opts.MaxMinutes {
			factors = append(factors, "good time fit")
		}
	}

	return factors
}

func (h HeuristicFallback) estimateSessionLength(task models.Task, maxMinutes int) int {
	if task.Metadata == nil {
		return min(25, maxMinutes) // Default pomodoro
	}

	baseMinutes := task.Metadata.Minutes
	if task.Metadata.Extend && maxMinutes >= task.Metadata.Estimate {
		return min(task.Metadata.Estimate, maxMinutes)
	}

	return min(baseMinutes, maxMinutes)
}

func (h HeuristicFallback) getSessionStrategy(task models.Task) string {
	if task.Metadata == nil {
		return "pomodoro"
	}

	if task.Metadata.Extend && task.Metadata.Estimate > 45 {
		return "hyperfocus"
	}

	if task.Metadata.Minutes <= 15 {
		return "quick"
	}

	return "pomodoro"
}

func (h HeuristicFallback) getEnergyMatchScore(taskEnergy, userEnergy string) float64 {
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

func (h HeuristicFallback) getModeMatchScore(taskMode, userMode string) float64 {
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

func (h HeuristicFallback) getTimeFitScore(metadata *models.HyperFocusMetadata, maxMinutes int) float64 {
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

// =============================================================================
// Utility Functions
// =============================================================================

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func joinStrings(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}

	result := strs[0]
	for i := 1; i < len(strs); i++ {
		result += sep + strs[i]
	}
	return result
}
