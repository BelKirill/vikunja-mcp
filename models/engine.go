// Package engine provides AI-powered task selection for focus sessions
package models

import (
	"time"
)

// DecisionRequest contains all context needed for AI-powered task selection
type DecisionRequest struct {
	// User context
	Energy     string    `json:"energy"` // "low"|"medium"|"high"|"social"
	Mode       string    `json:"mode"`   // "deep"|"quick"|"admin"
	MaxMinutes int       `json:"max_minutes"`
	Date       time.Time `json:"date"`
	TimeOfDay  string    `json:"time_of_day"` // "morning"|"afternoon"|"evening"

	// Available tasks with rich metadata
	CandidateTasks []Task `json:"candidate_tasks"`

	// Historical context (for learning)
	RecentSessions  []FocusSession    `json:"recent_sessions,omitempty"`
	EnergyPatterns  []EnergyPattern   `json:"energy_patterns,omitempty"`
	TaskPerformance []TaskPerformance `json:"task_performance,omitempty"`

	// Constraints
	MaxTasks        int      `json:"max_tasks"`
	ExcludeTaskIDs  []int64  `json:"exclude_task_ids,omitempty"`
	RequiredLabels  []string `json:"required_labels,omitempty"`
	BlockedProjects []int64  `json:"blocked_projects,omitempty"`
}

// DecisionResponse contains AI-ranked tasks with reasoning
type DecisionResponse struct {
	RankedTasks []RankedTask `json:"ranked_tasks"`
	Reasoning   string       `json:"reasoning"`
	Confidence  float64      `json:"confidence"` // 0.0-1.0
	Strategy    string       `json:"strategy"`   // Which strategy was used
	Fallback    bool         `json:"fallback"`   // Whether fallback was used
}

// RankedTask represents a task with AI-calculated suitability
type RankedTask struct {
	Task             Task     `json:"task"`
	Score            float64  `json:"score"`             // 0.0-1.0 suitability score
	ReasoningFactors []string `json:"reasoning_factors"` // Why this score
	EstimatedLength  int      `json:"estimated_length"`  // Recommended session length
	CanExtend        bool     `json:"can_extend"`
}

// TaskRecommendation is the single best task recommendation
type TaskRecommendation struct {
	Task              *RankedTask  `json:"task"`
	RecommendedLength int          `json:"recommended_length"`
	SessionStrategy   string       `json:"session_strategy"` // "pomodoro"|"hyperfocus"|"quick"
	Reasoning         string       `json:"reasoning"`
	Alternatives      []RankedTask `json:"alternatives"` // Other good options
}
