package engine

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/BelKirill/vikunja-mcp/models"
	"github.com/charmbracelet/log"
)

// extractJSON safely extracts the first valid JSON object from a string
// Returns the JSON string and any error encountered
func (e *FocusEngine) extractJSON(input string) (string, error) {
	log.Debug("extractJSON called", "input_length", len(input))
	if input == "" {
		log.Debug("extractJSON: input is empty")
		return "", nil
	}

	input = strings.TrimSpace(input)

	// Find the start of JSON object or array
	startObj := strings.Index(input, "{")
	startArr := strings.Index(input, "[")

	var start int
	var expectedEnd rune

	if startObj == -1 && startArr == -1 {
		return "", nil // No JSON found
	} else if startObj == -1 {
		start = startArr
		expectedEnd = ']'
	} else if startArr == -1 {
		start = startObj
		expectedEnd = '}'
	} else if startObj < startArr {
		start = startObj
		expectedEnd = '}'
	} else {
		start = startArr
		expectedEnd = ']'
	}

	// Safely track brace/bracket depth to find matching closing character
	depth := 0
	inString := false
	escaped := false

	for i := start; i < len(input); i++ {
		char := rune(input[i])

		// Handle escape sequences in strings
		if escaped {
			escaped = false
			continue
		}

		if char == '\\' {
			escaped = true
			continue
		}

		// Handle string boundaries
		if char == '"' {
			inString = !inString
			continue
		}

		// Only count braces/brackets outside of strings
		if !inString {
			if char == '{' || char == '[' {
				depth++
			} else if char == '}' || char == ']' {
				depth--
				if depth == 0 && char == expectedEnd {
					// Found complete JSON - validate it
					candidate := input[start : i+1]
					var js interface{}
					if json.Unmarshal([]byte(candidate), &js) == nil {
						return candidate, nil
					}
					// If JSON is invalid, continue searching
					break
				}
			}
		}
	}

	log.Debug("extractJSON: no valid JSON found")
	return "", nil // No valid JSON found
}

// embedMetadataInDescription safely embeds hyperfocus metadata into a task description
func (e *FocusEngine) embedMetadataInDescription(description string, metadata *models.HyperFocusMetadata) string {
	log.Debug("embedMetadataInDescription called", "desc_length", len(description), "metadata", metadata)
	if metadata == nil {
		log.Debug("No metadata provided, returning original description")
		return description
	}

	// Remove any existing JSON metadata first
	existingJSON, err := e.extractJSON(description)
	cleanDesc := description
	if err == nil && existingJSON != "" {
		log.Debug("Existing JSON found in description, cleaning", "existingJSON", existingJSON)
		cleanDesc = strings.Replace(description, existingJSON, "", 1)
		cleanDesc = strings.TrimSpace(cleanDesc)
	}

	// Create new metadata JSON
	metadataMap := map[string]interface{}{
		"energy":           metadata.Energy,
		"mode":             metadata.Mode,
		"extend":           metadata.Extend,
		"minutes":          metadata.Minutes,
		"estimate":         metadata.Estimate,
		"hyper_focus_comp": metadata.HyperFocusCompatability,
		"instructions":     metadata.Instructions,
		"dependencies":     metadata.Dependencies,
		"context_hints":    metadata.ContextualHints,
	}

	metadataJSON, err := json.Marshal(metadataMap)
	if err != nil {
		log.Error("Failed to marshal metadata", "error", err)
		return description // Return original on error
	}
	log.Debug("Metadata marshaled to JSON", "json", string(metadataJSON))

	// Combine clean description with new metadata
	if cleanDesc == "" {
		log.Debug("Clean description is empty, returning only metadata JSON")
		return string(metadataJSON)
	}
	log.Debug("Returning combined description and metadata JSON")
	return cleanDesc + " " + string(metadataJSON)
}

// Refactored EnrichTask to use the new helper functions and include labeling
func (e *FocusEngine) EnrichTask(ctx context.Context, task *models.RawTask, availableLabels []models.PartialLabel) (*models.Task, bool, error) {
	log.Info("enrichTask called", "task_id", task.ID, "title", task.Title)

	enrichedTask := &models.Task{
		RawTask:          task,             // Embed the raw task
		CleanDescription: task.Description, // Will be cleaned below
		Identifier:       task.Identifier,  // For human readable identifier like 'mcp-63'
	}

	// Step 1: Enrich metadata
	metadata, enriched, err := e.enrichMetadata(ctx, task)
	if err != nil {
		log.Warn("Couldn't get metadata enrichment", "err", err)
	} else if metadata != nil {
		enrichedTask.Metadata = metadata

		if enriched {
			// If we enriched with AI, embed the new metadata in description
			newDesc := e.embedMetadataInDescription(enrichedTask.RawTask.Description, metadata)
			enrichedTask.RawTask.Description = newDesc
		} else {
			// If metadata was already there, clean the description
			meta, _ := e.extractJSON(task.Description)
			if meta != "" {
				enrichedTask.CleanDescription = strings.Replace(task.Description, meta, "", 1)
				enrichedTask.CleanDescription = strings.TrimSpace(enrichedTask.CleanDescription)
				log.Debug("Cleaned description after removing JSON metadata", "clean_description", enrichedTask.CleanDescription)
			}
		}
	}

	// Step 2: Label task (get AI suggestions)
	if len(availableLabels) > 0 {
		suggestedLabels, err := e.LabelTask(ctx, task, availableLabels)
		if err != nil {
			log.Warn("Couldn't get label suggestions", "err", err)
		} else {
			enrichedTask.RawTask.Labels = suggestedLabels
			enriched = true
			log.Debug("Added suggested labels", "task_id", task.ID, "suggested_count", len(suggestedLabels))
		}
	} else {
		log.Debug("No available labels provided, skipping label suggestions")
	}

	log.Info("enrichTask returning", "task_id", task.ID, "has_metadata", enrichedTask.Metadata != nil, "suggested_labels_count", len(enrichedTask.RawTask.Labels))
	return enrichedTask, enriched, nil
}

// enrichMetadata handles just the metadata enrichment part
func (e *FocusEngine) enrichMetadata(ctx context.Context, task *models.RawTask) (*models.HyperFocusMetadata, bool, error) {
	log.Info("enrichMetadata called", "task_id", task.ID, "title", task.Title)

	meta, err := e.extractJSON(task.Description)
	if err != nil {
		log.Error("Failed to extract JSON from task description", "error", err, "task_id", task.ID)
		meta = ""
	}

	if meta == "" {
		log.Debug("No JSON metadata found, using AI enrichment", "task_id", task.ID)
		// No JSON metadata found - get AI enrichment
		newMeta, err := e.decisionEngine.EnrichTask(ctx, task)
		if err != nil {
			log.Warn("Couldn't get AI enrichment", "err", err)
			return nil, false, err
		}
		return newMeta, true, nil // true = was enriched by AI
	} else {
		log.Debug("JSON metadata found, attempting to unmarshal", "json", meta, "task_id", task.ID)
		// Parse existing JSON metadata
		var hyperfocusData models.HyperFocusMetadata
		if err := json.Unmarshal([]byte(meta), &hyperfocusData); err != nil {
			log.Error("Failed to unmarshal hyperfocus metadata", "error", err, "json", meta, "task_id", task.ID)
			return nil, false, err
		}
		log.Debug("Valid JSON metadata found and unmarshaled", "task_id", task.ID, "metadata", hyperfocusData)
		return &hyperfocusData, false, nil // false = already had metadata
	}
}

// LabelTask suggests appropriate labels for a task using AI
func (e *FocusEngine) LabelTask(ctx context.Context, task *models.RawTask, availableLabels []models.PartialLabel) ([]models.PartialLabel, error) {
	log.Info("LabelTask called", "task_id", task.ID, "title", task.Title, "available_labels_count", len(availableLabels))

	if len(availableLabels) == 0 {
		log.Debug("No available labels provided, returning empty slice")
		return []models.PartialLabel{}, nil
	}

	// Use the decision engine to get AI label suggestions
	suggestedLabels, err := e.decisionEngine.LabelTask(ctx, task, availableLabels)
	if err != nil {
		log.Error("Failed to get AI label suggestions", "error", err, "task_id", task.ID)
		return nil, err
	}

	log.Info("AI suggested labels", "task_id", task.ID, "suggested_count", len(suggestedLabels))
	return suggestedLabels, nil
}
