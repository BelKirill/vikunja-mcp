package focus

import (
	"github.com/BelKirill/vikunja-mcp/pkg/mcp"
	"github.com/charmbracelet/log"
)

// RegisterMCPTools registers focus-related MCP tools
func RegisterMCPTools(server *mcp.Server) error {
	log.Debug("Registering MCP tools for focus package")
	vikunjaSvc, err := NewService()
	if err != nil {
		log.Error("Failed to create focus service for MCP tools", "error", err)
		return err
	}

	// Register the daily-focus tool
	log.Debug("Registering daily-focus tool with MCP server")
	server.RegisterTool(mcp.Tool{
		Name:        "daily-focus",
		Description: "Get AI-recommended tasks for focus session based on energy/mode",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"energy": map[string]interface{}{
					"type":        "string",
					"enum":        []string{"low", "medium", "high", "social"},
					"description": "Current energy level",
				},
				"mode": map[string]interface{}{
					"type":        "string",
					"enum":        []string{"deep", "quick", "admin"},
					"description": "Work mode preference",
				},
				"hours": map[string]interface{}{
					"type":        "integer",
					"description": "Target work hours (default: 5)",
					"minimum":     1,
					"maximum":     12,
				},
				"max_items": map[string]interface{}{
					"type":        "integer",
					"description": "Maximum tasks to return (default: 10)",
					"minimum":     1,
					"maximum":     50,
				},
				"instructions": map[string]interface{}{
					"type":        "string",
					"description": "Additional details about the current request of the user",
				},
			},
		},
	}, func(args map[string]interface{}) (interface{}, error) {
		log.Debug("daily-focus tool invoked", "args", args)
		return handleDailyFocus(vikunjaSvc, args)
	})

	// Register task metadata reader
	log.Debug("Registering get-full-task tool with MCP server")
	server.RegisterTool(mcp.Tool{
		Name:        "get-full-task",
		Description: "Provide all the details possible for one task",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"task_id": map[string]interface{}{
					"type":        "integer",
					"description": "Vikunja task ID",
				},
			},
			"required": []string{"task_id"},
		},
	}, func(args map[string]interface{}) (interface{}, error) {
		log.Debug("get-full-task tool invoked", "args", args)
		return handleGetFullTask(vikunjaSvc, args)
	})

	// Register filtered task retrieval tool
	log.Debug("Registering get-filtered-tasks tool with MCP server")
	server.RegisterTool(mcp.Tool{
		Name:        "get-filtered-tasks",
		Description: "Retrieve tasks using AI-generated or manual filter expressions",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"filter": map[string]interface{}{
					"type":        "string",
					"description": "Vikunja filter expression (e.g., 'done = false && priority >= 3')",
				},
				"natural_request": map[string]interface{}{
					"type":        "string",
					"description": "Natural language request - will generate filter automatically if no filter provided",
				},
				"project_id": map[string]interface{}{
					"type":        "integer",
					"description": "Project ID to filter tasks within",
				},
				"limit": map[string]interface{}{
					"type":        "integer",
					"description": "Maximum number of tasks to return (default: 50)",
					"minimum":     1,
					"maximum":     200,
				},
			},
			"anyOf": []map[string]interface{}{
				{"required": []string{"filter"}},
				{"required": []string{"natural_request"}},
			},
		},
	}, func(args map[string]interface{}) (interface{}, error) {
		log.Debug("get-filtered-tasks tool invoked", "args", args)
		return handleGetFilteredTasks(vikunjaSvc, args)
	})

	// Register the upsert_task tool
	log.Debug("Registering upsert_task tool with MCP server")
	server.RegisterTool(mcp.Tool{
		Name:        "upsert_task",
		Description: "Create a new task or update an existing task (including marking complete)",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"task_id": map[string]interface{}{
					"type":        "integer",
					"description": "Task ID to update (omit to create new task)",
				},
				"title": map[string]interface{}{
					"type":        "string",
					"description": "Task title",
				},
				"done": map[string]interface{}{
					"type":        "boolean",
					"description": "Mark task as complete",
				},
				"priority": map[string]interface{}{
					"type":        "integer",
					"minimum":     1,
					"maximum":     5,
					"description": "Task priority (1-5)",
				},
				"description": map[string]interface{}{
					"type":        "string",
					"description": "Task description",
				},
				"hex_color": map[string]interface{}{
					"type":        "string",
					"description": "Colour of the task in Hexadecimal number, 6 characters",
					"pattern":     "^[0-9A-Fa-f]{6}$",
				},
				"project_id": map[string]interface{}{
					"type":        "integer",
					"description": "Project ID",
				},
			},
		},
	}, func(args map[string]interface{}) (interface{}, error) {
		log.Debug("upsert_task tool invoked", "args", args)
		return handleUpsertTask(vikunjaSvc, args)
	})

	log.Debug("All MCP tools registered for focus package")
	return nil
}
