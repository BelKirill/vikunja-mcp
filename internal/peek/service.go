package peek

import (
	"time"
)

// --- Service Interfaces ------------------------------------------------------

// TaskService defines how to fetch details about one task.
type TaskService interface {
	Peek(taskID string, fields []string) (map[string]interface{}, error)
}

// ProjectService defines how to fetch details about one project.
type ProjectService interface {
	Peek(projID string, fields []string) (map[string]interface{}, error)
}

// --- Default stubs (override via setter) ------------------------------------

var (
	taskService    TaskService    = &stubTaskService{}
	projectService ProjectService = &stubProjectService{}
)

func SetTaskService(svc TaskService) {
	taskService = svc
}

func SetProjectService(svc ProjectService) {
	projectService = svc
}

// --- Stub implementations -----------------------------------------------------

// Users should override these via SetTaskService / SetProjectService

type stubTaskService struct{}

func (s *stubTaskService) Peek(id string, fields []string) (map[string]interface{}, error) {
	now := time.Now().UTC().Format(time.RFC3339)
	result := make(map[string]interface{})
	for _, f := range fields {
		switch f {
		case "title":
			result[f] = "<stub title>"
		case "due":
			result[f] = now[:10]
		default:
			result[f] = nil
		}
	}
	return result, nil
}

type stubProjectService struct{}

func (s *stubProjectService) Peek(id string, fields []string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	for _, f := range fields {
		switch f {
		case "name":
			result[f] = "<stub project name>"
		default:
			result[f] = nil
		}
	}
	return result, nil
}
