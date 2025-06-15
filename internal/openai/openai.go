// Package openai provides OpenAI-powered decision engine for task selection
package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/BelKirill/vikunja-mcp/models"
	"github.com/charmbracelet/log"
)

// OpenAIConfig holds configuration for OpenAI API
type OpenAIConfig struct {
	APIKey      string
	BaseURL     string        // Default: "https://api.openai.com/v1"
	Model       string        // Default: "gpt-4"
	MaxTokens   int           // Default: 1000
	Temperature float64       // Default: 0.3 for consistent decisions
	Timeout     time.Duration // Default: 30s
}

// OpenAIDecisionEngine implements the DecisionEngine interface using OpenAI
type OpenAIDecisionEngine struct {
	config     OpenAIConfig
	httpClient *http.Client
}

// NewOpenAIDecisionEngine creates a new OpenAI-powered decision engine
func NewOpenAIDecisionEngine(config OpenAIConfig) *OpenAIDecisionEngine {
	// Set defaults
	if config.BaseURL == "" {
		config.BaseURL = "https://api.openai.com/v1"
	}
	if config.Model == "" {
		config.Model = "gpt-4"
	}
	if config.MaxTokens == 0 {
		config.MaxTokens = 1000
	}
	if config.Temperature == 0 {
		config.Temperature = 0.3
	}
	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second
	}

	return &OpenAIDecisionEngine{
		config: config,
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
	}
}

// OpenAI API request/response structures
type openAIRequest struct {
	Model       string    `json:"model"`
	Messages    []message `json:"messages"`
	MaxTokens   int       `json:"max_tokens"`
	Temperature float64   `json:"temperature"`
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type openAIResponse struct {
	Choices []choice  `json:"choices"`
	Error   *apiError `json:"error,omitempty"`
}

type choice struct {
	Message message `json:"message"`
}

type apiError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Code    string `json:"code"`
}

// RankTasks uses OpenAI to intelligently rank tasks for focus sessions
func (e *OpenAIDecisionEngine) RankTasks(ctx context.Context, request *models.DecisionRequest) (*models.DecisionResponse, error) {
	log.Info("OpenAI: Ranking tasks for focus session", "task_count", len(request.CandidateTasks))
	log.Debug("DecisionRequest", "energy", request.Energy, "mode", request.Mode, "max_minutes", request.MaxMinutes, "tasks", request.CandidateTasks)

	prompt := e.buildTaskRankingPrompt(request)
	log.Debug("Generated task ranking prompt", "prompt", prompt)

	response, err := e.callOpenAI(ctx, prompt)
	if err != nil {
		log.Error("OpenAI API call failed", "error", err)
		return nil, fmt.Errorf("OpenAI API call failed: %w", err)
	}

	log.Debug("Raw OpenAI response", "response", response)
	// Parse the AI response
	decision, err := e.parseTaskRankingResponse(response, request.CandidateTasks)
	if err != nil {
		log.Error("Failed to parse AI response", "error", err)
		return nil, fmt.Errorf("failed to parse AI response: %w", err)
	}

	log.Info("OpenAI: Successfully ranked tasks", "count", len(decision.RankedTasks))
	return decision, nil
}

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

// buildTaskRankingPrompt creates a structured prompt for task ranking
func (e *OpenAIDecisionEngine) buildTaskRankingPrompt(request *models.DecisionRequest) string {
	// Create simplified task representations for the prompt
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

	return fmt.Sprintf(`You are an AI assistant specialized in ADHD-friendly task management and focus optimization. 

Your job is to rank tasks for a focus session based on the user's current state and task characteristics.

## Current Context:
- User Energy Level: %s
- Focus Mode: %s  
- Available Time: %d minutes
- Time of Day: %s
- Date: %s
- Instructions: %s

## Available Tasks:
%s

## Instructions:
1. Analyze each task considering:
   - Energy level compatibility (low/medium/high/social)
   - Mode alignment (deep/quick/admin)
   - Time constraints and task estimates
   - Priority and urgency
   - Task complexity vs current mental state
   - Potential for hyperfocus if user has energy

2. Rank ALL tasks by suitability score (0.0-1.0)
3. Provide reasoning for top recommendations
4. Consider ADHD-specific factors:
   - Interest/motivation potential
   - Clear vs ambiguous tasks
   - Dopamine reward potential
   - Cognitive load requirements

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

Respond with valid JSON only.`,
		request.Energy,
		request.Mode,
		request.MaxMinutes,
		request.TimeOfDay,
		request.Date.Format("2006-01-02"),
		request.Instructions,
		string(tasksJSON))
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
		string(tasksJSON))
}

// callOpenAI makes the actual API call to OpenAI
func (e *OpenAIDecisionEngine) callOpenAI(ctx context.Context, prompt string) (string, error) {
	log.Info("Calling OpenAI API", "model", e.config.Model, "max_tokens", e.config.MaxTokens)
	log.Debug("Prompt sent to OpenAI", "prompt", prompt)

	reqBody := openAIRequest{
		Model:       e.config.Model,
		MaxTokens:   e.config.MaxTokens,
		Temperature: e.config.Temperature,
		Messages:    []message{{Role: "user", Content: prompt}},
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		log.Error("Failed to marshal OpenAI request body", "error", err)
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", e.config.BaseURL+"/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Error("Failed to create OpenAI HTTP request", "error", err)
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+e.config.APIKey)

	log.Debug("Sending HTTP request to OpenAI", "url", req.URL.String())
	resp, err := e.httpClient.Do(req)
	if err != nil {
		log.Error("HTTP request to OpenAI failed", "error", err)
		return "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer func() {
		cerr := resp.Body.Close()
		if cerr != nil {
			log.Warn("warning: failed to close response body", "error", cerr)
		}
	}()

	log.Debug("Received HTTP response from OpenAI", "status", resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Failed to read OpenAI response body", "error", err)
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	log.Debug("OpenAI response body", "body", string(body))
	if resp.StatusCode != http.StatusOK {
		log.Error("OpenAI API returned error status", "status", resp.StatusCode, "body", string(body))
		return "", fmt.Errorf("OpenAI API error (status %d): %s", resp.StatusCode, string(body))
	}

	var openAIResp openAIResponse
	if err := json.Unmarshal(body, &openAIResp); err != nil {
		log.Error("Failed to unmarshal OpenAI response JSON", "error", err)
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if openAIResp.Error != nil {
		log.Error("OpenAI API returned error", "error", openAIResp.Error.Message)
		return "", fmt.Errorf("OpenAI API error: %s", openAIResp.Error.Message)
	}

	if len(openAIResp.Choices) == 0 {
		log.Error("No choices in OpenAI response")
		return "", fmt.Errorf("no choices in OpenAI response")
	}

	log.Info("OpenAI API call successful, returning content")
	return openAIResp.Choices[0].Message.Content, nil
}

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

	if err := json.Unmarshal([]byte(response), &aiResponse); err != nil {
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
