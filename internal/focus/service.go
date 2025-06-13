package focus

import (
	"context"

	vikunja "github.com/BelKirill/vikunja-mcp/internal/vikunja"
	"github.com/BelKirill/vikunja-mcp/models"
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
func (s *Service) GetFocusTasks(ctx context.Context, opts models.FocusOptions) ([]models.FocusResult, error) {
	tasks, err := s.Vikunja.GetIncompleteTasks(ctx)
	if err != nil {
		return nil, err
	}
	results := make([]models.FocusResult, 0, len(tasks))
	for _, t := range tasks {
		results = append(results, models.FocusResult{
			TaskID:      t.TaskID,
			Project:     t.Project,
			Metadata:    s.ParseHyperFocusMetadata(t.Description),
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

func (s *Service) ParseHyperFocusMetadata(desc string) *models.HyperfocusMetadata {
	// TODO: Parse metadata from description if needed
	return nil
}

func (s *Service) cleanDescription(desc *string) string {
	return *desc
}

// Add this method to the Service struct
func (s *Service) UpsertTask(ctx context.Context, task models.MinimalTask) (*models.MinimalTask, error) {
	return s.Vikunja.UpsertTask(ctx, task)
}

func (s *Service) CleanDescription(desc string) string {
	return s.cleanDescription(&desc)
}
