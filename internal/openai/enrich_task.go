// Package openai provides OpenAI-powered decision engine for task enrichment
package openai

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/BelKirill/vikunja-mcp/models"
	"github.com/charmbracelet/log"
)

// EnrichTask uses OpenAI to enrich a task with additional information
func (e *OpenAIDecisionEngine) EnrichTask(ctx context.Context, task *models.RawTask) (*models.HyperFocusMetadata, error) {
	log.Info("OpenAI: Enriching task", "task", task)

	prompt := e.buildEnrichTaskPrompt(task)
	log.Debug("Generated enrich task prompt", "prompt", prompt)

	response, err := e.callOpenAI(ctx, prompt)
	if err != nil {
		log.Error("OpenAI API call failed", "error", err)
		return nil, fmt.Errorf("OpenAI API call failed: %w", err)
	}

	log.Debug("Raw OpenAI response", "response", response)

	enriched, err := e.parseEnrichTaskResponse(response)
	if err != nil {
		log.Error("Failed to parse AI response", "error", err)
		return nil, fmt.Errorf("failed to parse AI response: %w", err)
	}

	log.Info("OpenAI: Successfully enriched task", "enriched", enriched)
	return enriched, nil
}

// parseEnrichTaskResponse parses OpenAI's enrichment response
func (e *OpenAIDecisionEngine) parseEnrichTaskResponse(response string) (*models.HyperFocusMetadata, error) {
	var metadata models.HyperFocusMetadata
	if err := json.Unmarshal([]byte(response), &metadata); err != nil {
		log.Error("Failed to unmarshal enrichment response", "error", err, "response", response)
		return nil, fmt.Errorf("failed to parse enrichment response: %w", err)
	}
	return &metadata, nil
}

// buildEnrichTaskPrompt creates a prompt for task enrichment
func (e *OpenAIDecisionEngine) buildEnrichTaskPrompt(task *models.RawTask) string {
	taskJSON, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		log.Error("Failed to marshal task to JSON", "error", err)
		taskJSON = []byte(fmt.Sprintf("%+v", task))
	}

	return fmt.Sprintf(`# Vikunja AI task analyzer
		You are a productivity assistant helping an ADHD developer estimate task metadata. Analyze this task and provide JSON metadata.

		TASK: %s

		Provide JSON with these fields:
		- energy: "low" | "medium" | "high" | "social" (cognitive load required)
		- mode: "deep" | "quick" | "admin" (work style needed)
		- minutes: number (realistic time estimate 15-240)
		- hyper_focus_comp: 1-5 (ADHD compatibility: 1=avoid, 5=perfect for hyperfocus)
		- extend: boolean (can this task naturally extend beyond estimate?)
		- instructions: string (free text guidance for approaching this specific task)

		Consider:
		- Integration tasks typically need sustained focus but have clear endpoints
		- Immediate feedback (seeing AI replace hardcoded values) provides dopamine
		- Medium complexity - not trivial but not research-heavy
		- Well-defined completion criteria reduce decision fatigue

		Return only the JSON object.
	`, taskJSON)
}
