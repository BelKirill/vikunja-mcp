// Package engine provides AI-powered task selection for focus sessions
package engine

import (
	"context"
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
			NewPriorityConstraintFilter(),
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
