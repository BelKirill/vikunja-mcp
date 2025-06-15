// Package openai provides OpenAI-powered decision engine for task selection
package openai

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/BelKirill/vikunja-mcp/models"
	"github.com/charmbracelet/log"
)

// GetRecommendation uses OpenAI to get a single task recommendation with detailed reasoning
func (e *OpenAIDecisionEngine) GetRecommendation(ctx context.Context, request *models.DecisionRequest) (*models.TaskRecommendation, error) {
	log.Info("OpenAI: Getting task recommendation")
	log.Debug("DecisionRequest", "energy", request.Energy, "mode", request.Mode, "max_minutes", request.MaxMinutes, "tasks", request.CandidateTasks)

	prompt := e.buildRecommendationPrompt(request)
	log.Debug("Generated recommendation prompt", "prompt", prompt)

	response, err := e.callOpenAI(ctx, prompt)
	if err != nil {
		log.Error("OpenAI API call failed", "error", err)
		return nil, fmt.Errorf("OpenAI API call failed: %w", err)
	}

	log.Debug("Raw OpenAI response", "response", response)
	// Parse the AI response
	recommendation, err := e.parseRecommendationResponse(response, request.CandidateTasks)
	if err != nil {
		log.Error("Failed to parse AI response", "error", err)
		return nil, fmt.Errorf("failed to parse AI response: %w", err)
	}

	log.Info("OpenAI: Successfully generated recommendation")
	return recommendation, nil
}

// parseRecommendationResponse parses OpenAI's recommendation response
func (e *OpenAIDecisionEngine) parseRecommendationResponse(response string, tasks []models.Task) (*models.TaskRecommendation, error) {
	// Create task lookup map
	taskMap := make(map[int64]models.Task)
	for _, task := range tasks {
		taskMap[task.RawTask.ID] = task
	}

	// Parse JSON response
	var aiResponse struct {
		Task struct {
			TaskID           int64    `json:"task_id"`
			Score            float64  `json:"score"`
			ReasoningFactors []string `json:"reasoning_factors"`
			EstimatedLength  int      `json:"estimated_length"`
			CanExtend        bool     `json:"can_extend"`
		} `json:"task"`
		RecommendedLength int    `json:"recommended_length"`
		SessionStrategy   string `json:"session_strategy"`
		Reasoning         string `json:"reasoning"`
		Alternatives      []struct {
			TaskID           int64    `json:"task_id"`
			Score            float64  `json:"score"`
			ReasoningFactors []string `json:"reasoning_factors"`
		} `json:"alternatives"`
	}

	if err := json.Unmarshal([]byte(response), &aiResponse); err != nil {
		return nil, fmt.Errorf("failed to parse AI response JSON: %w", err)
	}

	// Get main task
	mainTask, exists := taskMap[aiResponse.Task.TaskID]
	if !exists {
		return nil, fmt.Errorf("AI recommended unknown task ID: %d", aiResponse.Task.TaskID)
	}

	rankedTask := models.RankedTask{
		Task:             mainTask,
		Score:            aiResponse.Task.Score,
		ReasoningFactors: aiResponse.Task.ReasoningFactors,
		EstimatedLength:  aiResponse.Task.EstimatedLength,
		CanExtend:        aiResponse.Task.CanExtend,
	}

	// Parse alternatives
	var alternatives []models.RankedTask
	for _, alt := range aiResponse.Alternatives {
		if altTask, exists := taskMap[alt.TaskID]; exists {
			alternatives = append(alternatives, models.RankedTask{
				Task:             altTask,
				Score:            alt.Score,
				ReasoningFactors: alt.ReasoningFactors,
			})
		}
	}

	return &models.TaskRecommendation{
		Task:              &rankedTask,
		RecommendedLength: aiResponse.RecommendedLength,
		SessionStrategy:   aiResponse.SessionStrategy,
		Reasoning:         aiResponse.Reasoning,
		Alternatives:      alternatives,
	}, nil
}

// buildRecommendationPrompt creates a prompt for single task recommendation
func (e *OpenAIDecisionEngine) buildRecommendationPrompt(request *models.DecisionRequest) string {
	taskSummaries := make([]map[string]interface{}, len(request.CandidateTasks))
	for i, task := range request.CandidateTasks {
		summary := map[string]interface{}{
			"id":          task.RawTask.ID,
			"title":       task.RawTask.Title,
			"description": task.CleanDescription,
			"priority":    task.RawTask.Priority,
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

	return fmt.Sprintf(`You are an AI assistant specialized in ADHD-friendly task management. 

Select the ONE BEST task for a focus session and provide detailed recommendations.

## Current Context:
- User Energy Level: %s
- Focus Mode: %s
- Available Time: %d minutes
- Time of Day: %s
- Instructions: %s

## Available Tasks:
%s

## Instructions:
1. Select the single best task considering:
   - Perfect energy/mode alignment
   - Optimal time utilization
   - Motivation and interest potential
   - ADHD-friendly characteristics
   - Priority and impact

2. Determine optimal session strategy:
   - "pomodoro" (25 min focused work)
   - "hyperfocus" (extended deep work)
   - "quick" (<15 min rapid completion)

3. Suggest 2-3 alternatives in case the top choice doesn't feel right

## Required Response Format (JSON):
{
  "task": {
    "task_id": 123,
    "score": 0.95,
    "reasoning_factors": ["perfect energy match", "high dopamine potential"],
    "estimated_length": 25,
    "can_extend": true
  },
  "recommended_length": 25,
  "session_strategy": "pomodoro",
  "reasoning": "Detailed explanation of why this is the optimal choice",
  "alternatives": [
    {
      "task_id": 456,
      "score": 0.87,
      "reasoning_factors": ["good priority", "clear scope"]
    }
  ]
}

Respond with valid JSON only.`,
		request.Energy,
		request.Mode,
		request.MaxMinutes,
		request.TimeOfDay,
		request.Instructions,
		string(tasksJSON))
}
