// Package openai provides OpenAI-powered decision engine for task selection
package openai

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/BelKirill/vikunja-mcp/models"
	"github.com/charmbracelet/log"
)

// SuggestFilter uses OpenAI to intelligently build a filter expression
func (e *OpenAIDecisionEngine) SuggestFilter(ctx context.Context, request *string) (*models.FilterSuggestionResponse, error) {
	log.Info("OpenAI: Suggesting a filter expression", "request", request)

	prompt := e.buildFilterSuggestionPrompt(*request)
	log.Debug("Generated filter suggestion prompt", "prompt", prompt)

	response, err := e.callOpenAI(ctx, prompt)
	if err != nil {
		log.Error("OpenAI API call failed", "error", err)
		return nil, fmt.Errorf("OpenAI API call failed: %w", err)
	}

	log.Debug("Raw OpenAI response", "response", response)

	// Parse the AI response
	suggestion, err := e.parseFilterSuggestionResponse(response)
	if err != nil {
		log.Error("Failed to parse AI response", "error", err)
		return nil, fmt.Errorf("failed to parse AI response: %w", err)
	}

	log.Info("OpenAI: Successfully generated filter suggestion",
		"filter", suggestion.Filter,
		"confidence", suggestion.Confidence)
	return suggestion, nil
}

// parseFilterSuggestionResponse parses OpenAI's filter suggestion response
func (e *OpenAIDecisionEngine) parseFilterSuggestionResponse(response string) (*models.FilterSuggestionResponse, error) {
	// Parse JSON response matching the expected format from the prompt
	var aiResponse struct {
		Filter     string  `json:"filter"`
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
		log.Error("Failed to unmarshal JSON response", "error", err, "response", response)

		// Return fallback response if parsing fails
		log.Warn("Falling back to basic filter due to parse error")
		return &models.FilterSuggestionResponse{
			Filter:     "done = false", // Basic fallback filter
			Reasoning:  "Failed to parse AI response, using basic incomplete tasks filter",
			Confidence: 0.3,
			Strategy:   "fallback",
			Fallback:   true,
		}, nil
	}

	// Validate required fields
	if aiResponse.Filter == "" {
		log.Warn("AI response missing filter field, using fallback")
		return &models.FilterSuggestionResponse{
			Filter:     "done = false",
			Reasoning:  "AI response missing filter expression, using basic incomplete tasks filter",
			Confidence: 0.3,
			Strategy:   "fallback",
			Fallback:   true,
		}, nil
	}

	// Set default confidence if not provided
	if aiResponse.Confidence == 0 {
		aiResponse.Confidence = 0.5
	}

	// Set default strategy if not provided
	if aiResponse.Strategy == "" {
		aiResponse.Strategy = "general"
	}

	return &models.FilterSuggestionResponse{
		Filter:     aiResponse.Filter,
		Reasoning:  aiResponse.Reasoning,
		Confidence: aiResponse.Confidence,
		Strategy:   aiResponse.Strategy,
		Fallback:   false, // This is a successful AI response
	}, nil
}

// buildFilterSuggestionPrompt creates a structured prompt for task filtering
func (e *OpenAIDecisionEngine) buildFilterSuggestionPrompt(request string) string {
	return fmt.Sprintf(`# Vikunja Filter Expression Builder

You are an expert at creating Vikunja filter expressions. Your task is to convert natural language requests into valid Vikunja filter syntax using SQL-like query expressions.

## Required filter: %s

## Core Syntax Rules

### Available Fields (All in snake_case)
- **done** - Whether task is completed (true/false)
- **priority** - Priority level (1-5, where 5 is highest)
- **percent_done** - Completion percentage (0-100)
- **due_date** - Task due date
- **start_date** - Task start date  
- **end_date** - Task end date
- **done_at** - When task was completed
- **assignees** - Users assigned to task
- **labels** - Labels/tags on task
- **project_id** - Project id of the project the task belongs to (only for saved filters, not project-level filters)
- **reminders** - Task reminders
- **created** - When task was created
- **updated** - When task was last modified

### Operators
- **=** - Equal to
- **!=** - Not equal to  
- **>** - Greater than
- **>=** - Greater than or equal to
- **<** - Less than
- **<=** - Less than or equal to
- **like** - Pattern matching with '%%' wildcards
- **in** - Match any value in comma-separated list
- **not in** - Match values NOT in comma-separated list

### Logical Operators
- **&&** - AND (all conditions must be true)
- **||** - OR (any condition can be true)
- **( )** - Parentheses for grouping

### String Values
- Single words: 'priority = 4'
- Multi-word strings: 'labels = "urgent task"'
- Dates: 'due_date = "2024-03-11"'

### Date Math (Relative Dates)
All date fields support relative date expressions:

**Anchors:**
- 'now' - Current date/time
- '"2024-03-11||"' - Specific date as anchor

**Time Units:**
- 's' - Seconds
- 'm' - Minutes  
- 'h' - Hours
- 'd' - Days
- 'w' - Weeks
- 'M' - Months
- 'y' - Years

**Operations:**
- '+1d' - Add one day
- '-1d' - Subtract one day
- '/d' - Round down to start of day
- '/w' - Round down to start of week

**Examples:**
- 'now' - Right now
- 'now+24h' - In 24 hours
- 'now/d' - Today at 00:00
- 'now/w' - Start of this week
- 'now+30d' - In 30 days
- 'now-1w/d' - One week ago at start of day

## Common Filter Patterns

### By Project
'''
project_id = 8
project_id in 8, 10
project_id != 1
'''

### By Labels  
'''
labels = "urgent"
labels in "bug", "feature"
labels like "%%urgent%%"
'''

### By Completion Status
'''
done = false
done = true
percent_done >= 50
percent_done < 100
'''

### By Due Date
'''
due_date < now
due_date >= now/d && due_date <= now/d+1d
due_date > now+7d
'''

### By Priority
'''
priority >= 3
priority = 5
priority in 4, 5
'''

### By Assignees
'''
assignees = "john@example.com"
assignees in "user1", "user2"
'''

## Complex Filter Examples

### Overdue High Priority Tasks
'''
due_date < now && priority >= 4 && done = false
'''

### This Week's Tasks
'''
due_date >= now/w && due_date <= now/w+1w
'''

### Urgent Tasks Not Yet Started
'''
labels = "urgent" && percent_done = 0 && done = false
'''

### Tasks Due Today or Tomorrow
'''
(due_date >= now/d && due_date <= now/d+1d) || (due_date >= now/d+1d && due_date <= now/d+2d)
'''

### Work Tasks Assigned to Me with High Priority
'''
project_id = 10 && assignees = "me@company.com" && priority >= 4
'''

### Tasks Created This Month but Not Done
'''
created >= now/M && done = false
'''

### Excluding Completed or Low Priority Tasks
'''
done = false && priority not in 1, 2
'''

## Response Format

When building filters, always:

1. **Analyze the request** - Identify key criteria (project, labels, dates, status, etc.)
2. **Choose appropriate fields** - Map natural language to exact field names
3. **Apply correct operators** - Use proper comparison and logical operators
4. **Handle dates intelligently** - Use date math for relative expressions
5. **Validate syntax** - Ensure proper parentheses and string quoting
6. **Provide the filter** - Give the complete, ready-to-use expression

## Required Response Format (JSON):
'''json
{
  "filter": "done = false && priority >= 3",
  "reasoning": "Filtered for incomplete tasks with medium to high priority based on the request",
  "confidence": 0.85,
  "strategy": "priority_focused"
}
'''

## Example Interactions

**User**: "Show me urgent tasks that are overdue"
**Filter**: 'labels = "urgent" && due_date < now && done = false'

**User**: "Find tasks in the Marketing project due this week"  
**Filter**: 'project = "Marketing" && due_date >= now/w && due_date <= now/w+1w'

**User**: "Get all completed tasks from last month"
**Filter**: 'done = true && done_at >= now-1M/M && done_at <= now/M'

Now respond with the appropriate Vikunja filter expression for the user's request.

Respond with valid JSON only.`,
		request)
}
