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

	return fmt.Sprintf(`You are a productivity AI that categorizes tasks for an ADHD developer. Be selective and specific with labels.

AVAILABLE LABELS:
%s

LABEL CATEGORIES & SPECIFIC USAGE:
• TECHNICAL: coding, development, API work, debugging, architecture
• CREATIVE: design, writing, brainstorming, planning, ideation  
• ADMINISTRATIVE: paperwork, forms, compliance, bureaucracy (rare)
• COMMUNICATION: meetings, emails, calls, collaboration
• LEARNING: study, research, tutorials, certification prep
• CONTEXT: online, quick, urgent (when specifically relevant)
• FOCUS: ONLY for tasks requiring 30+ minutes of sustained concentration

STRICT RULES:
1. **"focus" is NOT automatic** - Only for genuine deep work (30+ min concentration)
2. **Be minimalist** - 1-2 labels usually sufficient, 3 maximum
3. **Match the primary work type** - What is this task actually doing?
4. **Quick tasks (< 15 min) rarely need "focus"**
5. **Routine tasks rarely need special labels**

GOOD EXAMPLES:
• "Fix authentication bug" → [technical] 
• "Design API architecture" → [technical, design]
• "5-minute standup meeting" → [communication]
• "Fill out expense report" → [admin]
• "Study 2-hour course module" → [learning, focus]
• "Quick email check" → [communication]
• "Write comprehensive documentation" → [technical, writing]

BAD EXAMPLES (avoid these):
• "Fix small bug" → [technical, focus] ❌ (no focus needed for quick fix)
• "Simple task" → [focus] ❌ (focus is overused)
• "Any development work" → [admin] ❌ (admin is wrong category)

Response format: Return ONLY a JSON array of label IDs.
Example: [14] or [3, 14] or [] for no labels

Be selective - not every task needs multiple labels or "focus"!`, labelsJSON)
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

	return fmt.Sprintf(`Categorize this task with minimal, specific labels.

Ask yourself:
1. What type of work is this primarily? (technical/creative/communication/learning/admin)
2. Does this genuinely need sustained concentration (30+ min)? If yes, add "focus"
3. Are there other specific context labels that help? (online/quick/urgent)

Be selective - most tasks need only 1-2 labels maximum.

TASK TO ANALYZE:
%s

Return only the JSON array of the most essential label IDs for this task.`, string(taskJSON))
}
