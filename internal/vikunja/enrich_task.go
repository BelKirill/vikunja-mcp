package vikunja

import (
	"encoding/json"
	"strings"

	"github.com/BelKirill/vikunja-mcp/models"
	"github.com/charmbracelet/log"
)

// extractJSON safely extracts the first valid JSON object from a string
// Returns the JSON string and any error encountered
func extractJSON(input string) (string, error) {
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
func embedMetadataInDescription(description string, metadata *models.HyperFocusMetadata) string {
	log.Debug("embedMetadataInDescription called", "desc_length", len(description), "metadata", metadata)
	if metadata == nil {
		log.Debug("No metadata provided, returning original description")
		return description
	}

	// Remove any existing JSON metadata first
	existingJSON, err := extractJSON(description)
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

func enrichTask(task *models.RawTask) (*models.Task, error) {
	log.Info("enrichTask called", "task_id", task.ID, "title", task.Title)
	meta, err := extractJSON(task.Description)
	if err != nil {
		log.Error("Failed to extract JSON from task description", "error", err, "task_id", task.ID)
		meta = ""
	}

	enrichedTask := &models.Task{
		RawTask:          task,             // Embed the raw task
		CleanDescription: task.Description, // Will be cleaned below
	}

	if meta == "" {
		log.Debug("No JSON metadata found, using defaults", "task_id", task.ID)
		// No JSON metadata found - use defaults
		enrichedTask.Metadata = &models.HyperFocusMetadata{
			Energy:                  "medium", // Default energy level
			Mode:                    "quick",  // Default mode
			Extend:                  false,    // Default no extension
			Minutes:                 25,       // Default pomodoro
			Estimate:                25,       // Default estimate same as minutes
			HyperFocusCompatability: 3,        // Default middle compatibility
		}
		// Description stays as-is since no JSON to remove
	} else {
		log.Debug("JSON metadata found, attempting to unmarshal", "json", meta, "task_id", task.ID)
		// Parse the JSON metadata
		var hyperfocusData models.HyperFocusMetadata
		if err := json.Unmarshal([]byte(meta), &hyperfocusData); err != nil {
			log.Error("Failed to unmarshal hyperfocus metadata", "error", err, "json", meta, "task_id", task.ID)
			// JSON exists but invalid - use defaults
			enrichedTask.Metadata = &models.HyperFocusMetadata{
				Energy:                  "medium",
				Mode:                    "quick",
				Extend:                  false,
				Minutes:                 25,
				Estimate:                25,
				HyperFocusCompatability: 3,
			}
		} else {
			log.Debug("Valid JSON metadata found and unmarshaled", "task_id", task.ID, "metadata", hyperfocusData)
			// Valid JSON metadata found
			enrichedTask.Metadata = &hyperfocusData

			// Validate and set defaults for missing fields
			if hyperfocusData.Energy == "" {
				enrichedTask.Metadata.Energy = "medium"
			}
			if hyperfocusData.Mode == "" {
				enrichedTask.Metadata.Mode = "quick"
			}
			if hyperfocusData.Minutes == 0 {
				enrichedTask.Metadata.Minutes = 25
			}
			if hyperfocusData.Estimate == 0 {
				// If no estimate provided, use minutes as estimate
				enrichedTask.Metadata.Estimate = enrichedTask.Metadata.Minutes
			}
			if hyperfocusData.HyperFocusCompatability == 0 {
				enrichedTask.Metadata.HyperFocusCompatability = 3 // Default middle
			}
		}

		// Clean the description by removing the JSON metadata
		enrichedTask.CleanDescription = strings.Replace(task.Description, meta, "", 1)
		enrichedTask.CleanDescription = strings.TrimSpace(enrichedTask.CleanDescription)
		log.Debug("Cleaned description after removing JSON metadata", "clean_description", enrichedTask.CleanDescription)
	}
	log.Info("enrichTask returning", "task_id", task.ID, "has_metadata", enrichedTask.Metadata != nil)
	return enrichedTask, nil
}

func enrichTasks(tasks []models.RawTask) ([]models.Task, error) {
	log.Info("enrichTasks called", "task_count", len(tasks))
	enrichedTasks := make([]models.Task, 0, len(tasks))

	for _, task := range tasks {
		log.Debug("Enriching task", "task_id", task.ID)
		enriched, err := enrichTask(&task)
		if err != nil {
			log.Error("Failed to enrich task", "error", err, "task_id", task.ID)
			continue
		}
		enrichedTasks = append(enrichedTasks, *enriched)
	}
	log.Info("enrichTasks returning", "enriched_count", len(enrichedTasks))
	return enrichedTasks, nil
}

// // enrichMinimalTask enriches a MinimalTask with additional metadata or computed fields.
// func enrichMinimalTask(task *models.MinimalTask) *models.MinimalTask {
// 	log.Debug("enrichMinimalTask called", "task", task)
// 	if task == nil {
// 		log.Debug("enrichMinimalTask: task is nil")
// 		return nil
// 	}
// 	if task.Priority == 0 {
// 		log.Debug("Setting default priority", "task_id", task.TaskID)
// 		task.Priority = 3 // default priority
// 	}
// 	if task.Metadata == nil {
// 		log.Debug("Setting default metadata", "task_id", task.TaskID)
// 		task.Metadata = &models.HyperFocusMetadata{
// 			Energy:                  "medium",
// 			Mode:                    "quick",
// 			Extend:                  false,
// 			Minutes:                 25,
// 			Estimate:                25,
// 			HyperFocusCompatability: 3,
// 		}
// 	}
// 	if task.Title != "" && (containsIgnoreCase(task.Title, "urgent") || containsIgnoreCase(task.Description, "urgent")) {
// 		log.Info("Task marked as urgent", "task_id", task.TaskID)
// 		task.Priority = 5
// 	}
// 	log.Debug("enrichMinimalTask returning", "task", task)
// 	return task
// }

// // containsIgnoreCase checks if substr is in s, case-insensitive
// func containsIgnoreCase(s, substr string) bool {
// 	log.Debug("containsIgnoreCase called", "s", s, "substr", substr)
// 	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
// }
