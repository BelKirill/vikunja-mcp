package peek

// TaskPeekRequest represents the payload for a task peek request.
// swagger:model TaskPeekRequest
type TaskPeekRequest struct {
	// ID of the task.
	// example: "task-123"
	ID string `json:"id"`
	// Fields to include in the response.
	// example: ["name", "status"]
	Fields []string `json:"fields,omitempty"`
}

// TaskPeekResponse represents the response for a task peek request.
// swagger:model TaskPeekResponse
type TaskPeekResponse struct {
	// T is the task identifier.
	// example: "task-123"
	T string `json:"t"`
	// F is a map containing selected fields and values.
	F map[string]interface{} `json:"f"`
}

// ProjectPeekRequest represents the payload for a project peek request.
// swagger:model ProjectPeekRequest
type ProjectPeekRequest struct {
	// ID of the project.
	// example: "project-456"
	ID string `json:"id"`
	// Fields to include in the response.
	// example: ["name", "description"]
	Fields []string `json:"fields,omitempty"`
}

// ProjectPeekResponse represents the response for a project peek request.
// swagger:model ProjectPeekResponse
type ProjectPeekResponse struct {
	// P is the project identifier.
	// example: "project-456"
	P string `json:"p"`
	// F is a map containing selected fields and values.
	F map[string]interface{} `json:"f"`
}

// APIError represents a standard API error response.
// swagger:model APIError
type APIError struct {
	// Code is the HTTP status code.
	// example: 400
	Code int `json:"code"`
	// Message is a human-readable error message.
	// example: "invalid JSON body"
	Message string `json:"message"`
}
