package focus

import (
	"context"
	"fmt"

	vikunja "github.com/BelKirill/vikunja-mcp/internal/vikunja"
)

// --- Service -----------------------------------------------------------------

type Service struct {
	Vikunja *vikunja.Service
}

func NewService() (*Service, error) {
	vikunjaSvc, err := vikunja.NewService()
	if err != nil {
		return nil, err
	}
	return &Service{Vikunja: vikunjaSvc}, nil
}

// Example: Get all incomplete tasks and map to FocusResult
func (s *Service) GetFocusTasks(ctx context.Context, opts FocusOptions) ([]FocusResult, error) {
	tasks, err := s.Vikunja.GetIncompleteTasks(ctx)
	if err != nil {
		return nil, err
	}
	results := make([]FocusResult, 0, len(tasks))
	for _, t := range tasks {
		results = append(results, FocusResult{
			TaskID:      fmt.Sprintf("%v", t.ID),
			Project:     fmt.Sprintf("%v", t.ProjectID),
			Metadata:    s.parseHyperFocusMetadata(t.Description),
			Priority:    t.Priority,
			Title:       t.Title,
			Done:        t.Done,
			Description: t.Description,
			// TODO: Add Estimate if available in t
		})
	}
	// TODO: Add filtering/prioritization logic based on opts
	return results, nil
}

func (s *Service) parseHyperFocusMetadata(desc string) *HyperfocusMetadata {
	// TODO: Parse metadata from description if needed
	return nil
}

func (s *Service) cleanDescription(desc string) string {
	return desc
}
