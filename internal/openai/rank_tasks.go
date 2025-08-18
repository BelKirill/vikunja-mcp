// Package openai provides OpenAI-powered decision engine for task selection
package openai

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/BelKirill/vikunja-mcp/models"
	"github.com/charmbracelet/log"
)

// parseTaskRankingResponse parses OpenAI's task ranking response
func (e *OpenAIDecisionEngine) parseTaskRankingResponse(response string, tasks []models.Task) (*models.DecisionResponse, error) {
	// Create task lookup map
	taskMap := make(map[int64]models.Task)
	for _, task := range tasks {
		taskMap[task.RawTask.ID] = task
	}

	// Parse JSON response
	var aiResponse struct {
		RankedTasks []struct {
			TaskID           int64    `json:"task_id"`
			Score            float64  `json:"score"`
			ReasoningFactors []string `json:"reasoning_factors"`
			EstimatedLength  int      `json:"estimated_length"`
			CanExtend        bool     `json:"can_extend"`
		} `json:"ranked_tasks"`
		Reasoning  string  `json:"reasoning"`
		Confidence float64 `json:"confidence"`
		Strategy   string  `json:"strategy"`
	}

	cleanJSON, err := e.sanitizeResponse(response)
	if err != nil {
		log.Error("Error in sanitizing response!", "response", response)
		return nil, fmt.Errorf("failed to parse enrichment response: %w", err)
	}

	if err := json.Unmarshal([]byte(cleanJSON), &aiResponse); err != nil {
		return nil, fmt.Errorf("failed to parse AI response JSON: %w", err)
	}

	// Convert to internal format
	var rankedTasks []models.RankedTask
	for _, aiTask := range aiResponse.RankedTasks {
		task, exists := taskMap[aiTask.TaskID]
		if !exists {
			log.Warn("AI referenced unknown task ID", "task_id", aiTask.TaskID)
			continue
		}

		rankedTasks = append(rankedTasks, models.RankedTask{
			Task:             task,
			Score:            aiTask.Score,
			ReasoningFactors: aiTask.ReasoningFactors,
			EstimatedLength:  aiTask.EstimatedLength,
			CanExtend:        aiTask.CanExtend,
		})
	}

	return &models.DecisionResponse{
		RankedTasks: rankedTasks,
		Reasoning:   aiResponse.Reasoning,
		Confidence:  aiResponse.Confidence,
		Strategy:    aiResponse.Strategy,
		Fallback:    false,
	}, nil
}

// openai/ranking.go - Add these methods to match your labeling pattern

// buildTaskRankingSystemPrompt creates the system prompt for ranking
func (e *OpenAIDecisionEngine) buildTaskRankingSystemPrompt(request *models.DecisionRequest) string {
	projectsJSON, _ := json.MarshalIndent(request.Projects, "", "  ")

	return fmt.Sprintf(`You are an AI assistant specialized in ADHD-friendly task management and focus optimization.

Your job is to rank tasks for a focus session based on the user's current state and task characteristics.

## Analysis Framework:
1. Analyze each task considering:
   - Energy level compatibility (low/medium/high/social)
   - Mode alignment (deep/quick/admin)
   - Time constraints and task estimates
   - Priority and urgency
   - Task complexity vs current mental state
   - Potential for hyperfocus if user has energy
2. Rank ALL tasks by suitability score (0.0-1.0)
3. Consider ADHD-specific factors:
	 - Switching context between projects has a cost
   - Interest/motivation potential
   - Clear vs ambiguous tasks
   - Dopamine reward potential
   - Cognitive load requirements
4. Consider the different projects:
   %s

## Required Response Format (JSON):
{
  "ranked_tasks": [
    {
      "task_id": 123,
      "score": 0.95,
      "reasoning_factors": ["high priority", "energy match", "clear scope"],
      "estimated_length": 25,
      "can_extend": true
    }
  ],
  "reasoning": "Overall strategy explanation",
  "confidence": 0.85,
  "strategy": "energy_optimized"
}

Respond with valid JSON only.`, string(projectsJSON))
}

// buildTaskRankingUserPrompt creates the user prompt with context and tasks
func (e *OpenAIDecisionEngine) buildTaskRankingUserPrompt(request *models.DecisionRequest) string {
	// Create simplified task representations for the prompt
	taskSummaries := make([]map[string]interface{}, len(request.CandidateTasks))
	for i, task := range request.CandidateTasks {
		summary := map[string]interface{}{
			"id":          task.RawTask.ID,
			"title":       task.RawTask.Title,
			"description": task.CleanDescription,
			"priority":    task.RawTask.Priority,
			"project_id":  task.RawTask.ProjectID,
			"labels":      task.RawTask.Labels,
		}

		if task.Metadata != nil {
			summary["energy"] = task.Metadata.Energy
			summary["mode"] = task.Metadata.Mode
			summary["minutes"] = task.Metadata.Minutes
			summary["estimate"] = task.Metadata.Estimate
			summary["extend"] = task.Metadata.Extend
		}

		taskSummaries[i] = summary
	}

	tasksJSON, _ := json.MarshalIndent(taskSummaries, "", "  ")

	return fmt.Sprintf(`## Current Context:
- User Energy Level: %s
- Focus Mode: %s  
- Available Time: %d minutes
- Time of Day: %s
- Date: %s
- Instructions: %s

## Available Tasks:
%s

Please rank these tasks for the current focus session.`,
		request.Energy,
		request.Mode,
		request.MaxMinutes,
		request.TimeOfDay,
		request.Date.Format("2006-01-02"),
		request.Instructions,
		string(tasksJSON))
}

// Updated RankTasks method to use separated prompts
func (e *OpenAIDecisionEngine) RankTasks(ctx context.Context, request *models.DecisionRequest) (*models.DecisionResponse, error) {
	log.Info("OpenAI: Ranking tasks for focus session", "task_count", len(request.CandidateTasks))

	systemPrompt := e.buildTaskRankingSystemPrompt(request)
	userPrompt := e.buildTaskRankingUserPrompt(request)

	log.Debug("Generated prompts", "system_prompt", systemPrompt, "user_prompt", userPrompt)

	// Use your existing callOpenAIWithSystem method (like in labeling)
	response, err := e.callOpenAIWithSystem(ctx, systemPrompt, userPrompt)
	if err != nil {
		log.Error("OpenAI API call failed", "error", err)
		return nil, fmt.Errorf("OpenAI API call failed: %w", err)
	}

	log.Debug("Raw OpenAI response", "response", response)

	decision, err := e.parseTaskRankingResponse(response, request.CandidateTasks)
	if err != nil {
		log.Error("Failed to parse AI response", "error", err)
		return nil, fmt.Errorf("failed to parse AI response: %w", err)
	}

	log.Info("OpenAI: Successfully ranked tasks", "count", len(decision.RankedTasks))
	return decision, nil
}
