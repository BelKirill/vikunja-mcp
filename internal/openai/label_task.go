// Package openai provides OpenAI-powered decision engine for task labeling
package openai

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/BelKirill/vikunja-mcp/models"
	"github.com/charmbracelet/log"
)

// LabelTask uses OpenAI to suggest appropriate labels for a task
func (e *OpenAIDecisionEngine) LabelTask(ctx context.Context, task *models.RawTask, availableLabels []models.PartialLabel) ([]models.PartialLabel, error) {
	log.Info("OpenAI: Labeling task", "task_id", task.ID, "title", task.Title)

	systemPrompt := e.buildLabelTaskSystemPrompt(availableLabels)
	userPrompt := e.buildLabelTaskUserPrompt(task)

	log.Debug("Generated label task prompts", "system", systemPrompt, "user", userPrompt)

	// Call OpenAI with both system and user prompts
	response, err := e.callOpenAIWithSystem(ctx, systemPrompt, userPrompt)
	if err != nil {
		log.Error("OpenAI API call failed", "error", err)
		return nil, fmt.Errorf("OpenAI API call failed: %w", err)
	}

	log.Debug("Raw OpenAI response", "response", response)

	suggestedLabels, err := e.parseLabelTaskResponse(response, availableLabels)
	if err != nil {
		log.Error("Failed to parse label response", "error", err)
		return nil, fmt.Errorf("failed to parse label response: %w", err)
	}

	log.Info("OpenAI: Successfully suggested labels", "count", len(suggestedLabels))
	return suggestedLabels, nil
}

// buildLabelTaskSystemPrompt creates the system prompt with available labels
func (e *OpenAIDecisionEngine) buildLabelTaskSystemPrompt(availableLabels []models.PartialLabel) string {
	// Convert available labels to JSON for the prompt
	labelsJSON := "["
	for i, label := range availableLabels {
		labelData := map[string]interface{}{
			"id":          label.ID,
			"title":       label.Title,
			"description": label.Description,
		}
		labelBytes, _ := json.Marshal(labelData)
		labelsJSON += string(labelBytes)
		if i < len(availableLabels)-1 {
			labelsJSON += ","
		}
	}
	labelsJSON += "]"

	return fmt.Sprintf(`You are a productivity AI that helps categorize tasks with appropriate labels for an ADHD developer.

Your job: Analyze the given task and suggest which of the available labels should be applied.

AVAILABLE LABELS:
%s

Guidelines:
- Only suggest labels that genuinely apply to the task
- Consider task complexity, type of work, priority indicators, and context
- For ADHD productivity: prioritize labels that help with energy management, context switching, and focus
- Return 0-5 labels maximum (don't over-label)
- Focus on actionable categorization

Response format: Return a JSON array of label IDs that should be applied.
Example: [1, 3, 7] or [] for no labels

Return only the JSON array, no other text.`, labelsJSON)
}

// buildLabelTaskUserPrompt creates the user prompt with the task details
func (e *OpenAIDecisionEngine) buildLabelTaskUserPrompt(task *models.RawTask) string {
	taskJSON, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		log.Error("Failed to marshal task to JSON", "error", err)
		// Fallback to basic string representation
		return fmt.Sprintf("Task ID: %d\nTitle: %s\nDescription: %s\nPriority: %d",
			task.ID, task.Title, task.Description, task.Priority)
	}

	return fmt.Sprintf("Analyze this task and suggest appropriate labels:\n\n%s", string(taskJSON))
}

// parseLabelTaskResponse parses OpenAI's label suggestion response
func (e *OpenAIDecisionEngine) parseLabelTaskResponse(response string, availableLabels []models.PartialLabel) ([]models.PartialLabel, error) {
	cleanJSON, err := e.sanitizeResponse(response)
	if err != nil {
		log.Error("Error sanitizing label response", "response", response)
		return nil, fmt.Errorf("failed to sanitize label response: %w", err)
	}

	var labelIDs []int
	if err := json.Unmarshal([]byte(cleanJSON), &labelIDs); err != nil {
		log.Error("Failed to unmarshal label response", "error", err, "response", cleanJSON)
		return nil, fmt.Errorf("failed to parse label response: %w", err)
	}

	// Convert label IDs back to PartialLabel objects
	var suggestedLabels []models.PartialLabel
	labelMap := make(map[int]models.PartialLabel)

	// Create lookup map for available labels
	for _, label := range availableLabels {
		labelMap[label.ID] = label
	}

	// Find matching labels
	for _, id := range labelIDs {
		if label, exists := labelMap[id]; exists {
			suggestedLabels = append(suggestedLabels, label)
		} else {
			log.Warn("AI suggested non-existent label", "label_id", id)
		}
	}

	return suggestedLabels, nil
}
