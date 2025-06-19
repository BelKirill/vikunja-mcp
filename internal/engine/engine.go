// Package engine provides AI-powered task selection for focus sessions
package engine

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/BelKirill/vikunja-mcp/models"
	"github.com/charmbracelet/log"
)

// FocusEngine orchestrates intelligent task selection with multiple strategies
type FocusEngine struct {
	decisionEngine    models.DecisionEngine
	fallbackStrategy  models.FallbackStrategy
	contextualFilters []models.ContextualFilter
	learningEnabled   bool
}

// NewFocusEngine creates a new focus engine with pluggable components
func NewFocusEngine(decisionEngine models.DecisionEngine, options ...EngineOption) *FocusEngine {
	engine := &FocusEngine{
		decisionEngine:   decisionEngine,
		fallbackStrategy: NewHeuristicFallback(),
		contextualFilters: []models.ContextualFilter{
			NewTimeConstraintFilter(),
			NewEnergyModeFilter(),
			NewDependencyFilter(),
		},
		learningEnabled: true,
	}

	for _, opt := range options {
		opt(engine)
	}

	return engine
}

// EngineOption allows configuring the focus engine
type EngineOption func(*FocusEngine)

// WithFallbackStrategy sets a custom fallback strategy
func WithFallbackStrategy(strategy *models.FallbackStrategy) EngineOption {
	return func(e *FocusEngine) {
		e.fallbackStrategy = *strategy
	}
}

// WithContextualFilters sets custom contextual filters
func WithContextualFilters(filters ...models.ContextualFilter) EngineOption {
	return func(e *FocusEngine) {
		e.contextualFilters = filters
	}
}

// WithLearning enables/disables learning from past sessions
func WithLearning(enabled bool) EngineOption {
	return func(e *FocusEngine) {
		e.learningEnabled = enabled
	}
}

// SuggestFilter return AI filtered tasks
func (e *FocusEngine) SuggestFilter(ctx context.Context, filter *string) (*models.FilterSuggestionResponse, error) {
	log.Info("FocusEngine.SuggestFilter called", "filter", &filter)

	response, err := e.decisionEngine.SuggestFilter(ctx, filter)
	if err != nil {
		log.Error("AI decision engine failed, filter directly", "error", err)
		return nil, err
	}

	log.Info("FilterEngine returning AI-suggested filter", "filter", response.Filter)
	return response, nil
}

// GetFocusTasks returns AI-ranked tasks for focus session
func (e *FocusEngine) GetFocusTasks(ctx context.Context, tasks []models.Task, opts *models.FocusOptions) (*models.DecisionResponse, error) {
	log.Info("FocusEngine.GetFocusTasks called", "task_count", len(tasks), "opts", opts)

	// Step 1: Apply contextual filters
	filteredTasks := e.applyContextualFilters(ctx, tasks, opts)
	if len(filteredTasks) == 0 {
		log.Warn("No tasks remaining after contextual filtering")
		return &models.DecisionResponse{
			RankedTasks: []models.RankedTask{},
			Reasoning:   "No tasks match current context and constraints",
			Confidence:  0.0,
			Strategy:    "contextual_filter",
			Fallback:    false,
		}, nil
	}

	// Step 2: Build decision request with rich context
	request := e.buildDecisionRequest(ctx, filteredTasks, opts)

	// Step 3: Try AI-powered ranking
	response, err := e.decisionEngine.RankTasks(ctx, request)
	if err != nil {
		log.Warn("AI decision engine failed, using fallback", "error", err)
		return e.fallbackToHeuristic(ctx, filteredTasks, opts)
	}

	// Step 4: Post-process and validate results
	e.validateAndEnrichResponse(response, opts)

	log.Info("FocusEngine returning AI-ranked tasks", "count", len(response.RankedTasks))
	return response, nil
}

// applyContextualFilters runs all contextual filters in sequence
func (e *FocusEngine) applyContextualFilters(ctx context.Context, tasks []models.Task, opts *models.FocusOptions) []models.Task {
	filtered := tasks

	for _, filter := range e.contextualFilters {
		before := len(filtered)
		filtered = filter.Filter(ctx, filtered, opts)
		log.Debug("Applied contextual filter",
			"filter", filter.Name(),
			"before", before,
			"after", len(filtered))
	}

	return filtered
}

// buildDecisionRequest creates rich context for AI decision making
func (e *FocusEngine) buildDecisionRequest(ctx context.Context, tasks []models.Task, opts *models.FocusOptions) *models.DecisionRequest {
	// Use context to get current time if available, otherwise fallback to time.Now()
	var now time.Time
	if ctx != nil {
		if deadline, ok := ctx.Deadline(); ok {
			now = deadline
		} else {
			now = time.Now()
		}
	} else {
		now = time.Now()
	}

	request := &models.DecisionRequest{
		Energy:         opts.Energy,
		Mode:           opts.Mode,
		MaxMinutes:     opts.MaxMinutes,
		Date:           now,
		TimeOfDay:      getTimeOfDay(now),
		CandidateTasks: tasks,
		MaxTasks:       opts.MaxTasks,
		Instructions:   opts.Instructions,
	}

	// // Add learning context if enabled
	// if e.learningEnabled {
	// 	// These would be loaded from your learning system
	// 	// request.RecentSessions = e.loadRecentSessions(ctx)
	// 	// request.EnergyPatterns = e.loadEnergyPatterns(ctx)
	// 	// request.TaskPerformance = e.loadTaskPerformance(ctx)
	// 	// No-op for now
	// }

	return request
}

// fallbackToHeuristic provides deterministic task selection when AI fails
func (e *FocusEngine) fallbackToHeuristic(ctx context.Context, tasks []models.Task, opts *models.FocusOptions) (*models.DecisionResponse, error) {
	log.Info("Using heuristic fallback for task selection")

	rankedTasks := e.fallbackStrategy.SelectTasks(ctx, tasks, opts)

	return &models.DecisionResponse{
		RankedTasks: rankedTasks,
		Reasoning:   "Used heuristic fallback due to AI service unavailability",
		Confidence:  0.7, // Moderate confidence in heuristic
		Strategy:    "heuristic_fallback",
		Fallback:    true,
	}, nil
}

// validateAndEnrichResponse ensures response quality and adds missing data
func (e *FocusEngine) validateAndEnrichResponse(response *models.DecisionResponse, opts *models.FocusOptions) {
	log.Debug("Validating and enriching response", "ranked_tasks", len(response.RankedTasks), "max_tasks", opts.MaxTasks)
	// Limit to requested max tasks
	if opts.MaxTasks > 0 && len(response.RankedTasks) > opts.MaxTasks {
		log.Debug("Trimming ranked tasks to max_tasks", "before", len(response.RankedTasks), "after", opts.MaxTasks)
		response.RankedTasks = response.RankedTasks[:opts.MaxTasks]
	}

	// Ensure scores are valid
	for i := range response.RankedTasks {
		if response.RankedTasks[i].Score < 0.0 {
			log.Debug("Score below 0.0, correcting", "index", i, "score", response.RankedTasks[i].Score)
			response.RankedTasks[i].Score = 0.0
		}
		if response.RankedTasks[i].Score > 1.0 {
			log.Debug("Score above 1.0, correcting", "index", i, "score", response.RankedTasks[i].Score)
			response.RankedTasks[i].Score = 1.0
		}
	}

	// Sort by score (should already be sorted, but ensure it)
	log.Debug("Sorting ranked tasks by score")
	sort.Slice(response.RankedTasks, func(i, j int) bool {
		return response.RankedTasks[i].Score > response.RankedTasks[j].Score
	})
}

// getTimeOfDay determines time of day category
func getTimeOfDay(t time.Time) string {
	hour := t.Hour()
	switch {
	case hour < 12:
		return "morning"
	case hour < 17:
		return "afternoon"
	default:
		return "evening"
	}
}

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

// =============================================================================
// Heuristic Fallback Strategy
// =============================================================================

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
