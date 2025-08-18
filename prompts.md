# Prompts Documentation

This document describes all the AI prompts used in the vikunja-mcp system for ADHD-optimized task management.

## Overview

The system uses a multi-tier LLM approach where each prompt serves a specific ADHD-optimized function in the task management workflow. All prompts follow structured JSON output patterns and include comprehensive ADHD-focused decision-making criteria.

## Prompt Locations and Functions

### 1. **Task Ranking Prompt** (`internal/openai/rank_tasks.go`)

**System Prompt** (`buildTaskRankingSystemPrompt`):
- **Purpose**: Instructs AI to rank tasks for ADHD-optimized focus sessions
- **Key Features**: 
  - Energy level compatibility (low/medium/high/social)
  - Mode alignment (deep/quick/admin) 
  - ADHD-specific factors (decision paralysis, hyperfocus, dopamine rewards)
  - Pomodoro integration with 25-minute base units
- **Output**: JSON with ranked tasks, scores, reasoning factors, and strategy

**User Prompt** (`buildTaskRankingUserPrompt`):
- **Purpose**: Provides current context and available tasks for ranking
- **Includes**: Energy level, focus mode, available time, task metadata

### 2. **Task Enrichment Prompt** (`internal/openai/enrich_task.go`)

**Single Prompt** (`buildEnrichTaskPrompt`):
- **Purpose**: Analyzes tasks to generate ADHD-specific metadata
- **Output**: JSON with energy requirements, mode, time estimates, hyperfocus compatibility (1-5), extend capability, and specific instructions
- **Focus**: Optimizes for ADHD workflows with clear completion criteria

### 3. **Task Labeling Prompts** (`internal/openai/label_task.go`)

**System Prompt** (`buildLabelTaskSystemPrompt`):
- **Purpose**: Intelligently categorizes tasks with appropriate labels
- **Categories**: Technical, Creative, Administrative, Communication, Learning, Context, Focus
- **Rules**: Minimalist approach (1-2 labels max), "focus" only for 30+ minute deep work
- **Output**: JSON array of label IDs

**User Prompt** (`buildLabelTaskUserPrompt`):
- **Purpose**: Provides task details for label analysis
- **Approach**: Asks specific categorization questions to ensure selectivity

### 4. **Filter Suggestion Prompt** (`internal/openai/suggest_filter.go`)

**Single Comprehensive Prompt** (`buildFilterSuggestionPrompt`):
- **Purpose**: Converts natural language requests into Vikunja filter expressions
- **Features**: 
  - Complete Vikunja syntax documentation
  - Date math operations (relative dates)
  - Complex filter examples
  - SQL-like query expressions
- **Output**: JSON with filter expression, reasoning, confidence, and strategy

## Prompt Engineering Patterns

### Common Architecture
1. **Structured Output**: All prompts request JSON responses for reliable parsing
2. **Context-Aware**: User's current energy, mode, and time constraints are central
3. **ADHD-Optimized**: Focus on decision paralysis reduction, hyperfocus compatibility, and dopamine rewards
4. **Fallback Handling**: Each prompt includes error recovery strategies

### Files to Modify for Prompt Improvements
- **Task Ranking**: `internal/openai/rank_tasks.go:75-165`
- **Task Enrichment**: `internal/openai/enrich_task.go:56-84` 
- **Task Labeling**: `internal/openai/label_task.go:76-158`
- **Filter Suggestion**: `internal/openai/suggest_filter.go:102-295`

## Detailed Prompt Analysis

### Task Ranking System Prompt
```go
// Location: internal/openai/rank_tasks.go:75-116
func (e *OpenAIDecisionEngine) buildTaskRankingSystemPrompt(request *models.DecisionRequest) string
```

This prompt establishes the AI as an ADHD-friendly task management specialist with a comprehensive analysis framework:

1. **Energy Compatibility Analysis**: Matches task cognitive load to user's current energy state
2. **Mode Alignment**: Ensures tasks fit the intended work style (deep focus vs quick tasks)
3. **ADHD-Specific Considerations**:
   - Context switching costs between projects
   - Interest/motivation potential for dopamine rewards
   - Clear vs ambiguous task definitions
   - Cognitive load requirements

### Task Enrichment Prompt
```go
// Location: internal/openai/enrich_task.go:56-84
func (e *OpenAIDecisionEngine) buildEnrichTaskPrompt(task *models.RawTask) string
```

Focused on generating hyperfocus-compatible metadata:
- **Energy Estimation**: Cognitive load analysis (low/medium/high/social)
- **Work Mode Classification**: Deep work vs quick tasks vs administrative
- **Time Estimation**: Realistic 15-240 minute ranges
- **Hyperfocus Compatibility**: 1-5 scale for ADHD suitability
- **Extension Capability**: Whether tasks can naturally expand beyond estimates

### Task Labeling System Prompt
```go
// Location: internal/openai/label_task.go:76-132
func (e *OpenAIDecisionEngine) buildLabelTaskSystemPrompt(availableLabels []models.PartialLabel) string
```

Implements strict labeling rules to prevent label pollution:
- **Minimalist Philosophy**: 1-2 labels maximum, 3 absolute max
- **"Focus" Label Rules**: Only for genuine deep work requiring 30+ minutes
- **Category Specificity**: Technical, Creative, Communication, Learning, Administrative
- **Anti-Patterns**: Explicit examples of what NOT to label

### Filter Suggestion Prompt
```go
// Location: internal/openai/suggest_filter.go:102-295
func (e *OpenAIDecisionEngine) buildFilterSuggestionPrompt(request string) string
```

Comprehensive Vikunja query language documentation:
- **Complete Syntax Reference**: All available fields and operators
- **Date Math Operations**: Relative date expressions with anchors and time units
- **Complex Examples**: Real-world filter patterns for common use cases
- **Validation Rules**: Proper parentheses, string quoting, and syntax requirements

## Usage in Practice

Each prompt is called through the OpenAI decision engine with proper error handling and response sanitization. The system supports both single-prompt calls (`callOpenAI`) and system/user prompt pairs (`callOpenAIWithSystem`) for maximum flexibility.

The prompts work together to create a comprehensive ADHD-optimized task management experience:
1. **Enrichment** adds metadata to raw tasks
2. **Labeling** categorizes tasks for filtering
3. **Ranking** prioritizes tasks for focus sessions
4. **Filtering** creates dynamic task queries

All prompts emphasize explainable AI with reasoning provided alongside recommendations.