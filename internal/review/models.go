package review

// ReviewRequest mirrors the `360-review` contract.
// swagger:model ReviewRequest
// req: {"project_ids": ["proj-1", "proj-2"]}
type ReviewRequest struct {
	// ProjectIDs is the list of project identifiers.
	// example: ["proj-1", "proj-2"]
	ProjectIDs []string `json:"project_ids"`
}

// ReviewResponseItem is a slim response containing project, task, and due date.
// swagger:model ReviewResponseItem
type ReviewResponseItem struct {
	// P represents the project identifier.
	// example: "proj-1"
	P string `json:"p"`

	// T represents the task identifier.
	// example: "task-123"
	T string `json:"t"`

	// D is the due date in YYYY-MM-DD format.
	// example: "2025-05-26"
	D string `json:"d"`
}
