package focus

import (
	"context"
	"testing"
	"time"

	"github.com/BelKirill/vikunja-mcp/models"
	"github.com/stretchr/testify/assert"
)

// MockVikunjaService allows us to inject errors for testing
type MockVikunjaService struct {
	ShouldError bool
}

func (m *MockVikunjaService) GetIncompleteTasks(ctx context.Context) ([]interface{}, error) {
	if m.ShouldError {
		return nil, assert.AnError
	}
	return []interface{}{}, nil
}

func TestFocusService_Focus_Empty(t *testing.T) {
	// Example: test Focus with empty input (mock or minimal implementation)
	s, err := NewService() // replace with actual service if needed
	if err != nil {
		t.Fatalf("failed to create service: %v", err)
	}
	ctx := context.Background()
	opts := models.FocusOptions{
		Energy:     "low",
		Mode:       "",
		MaxMinutes: 5,
		MaxTasks:   5,
		Date:       time.Now().String(),
	}

	items, _ := s.GetFocusTasks(ctx, opts)
	assert.IsType(t, models.FocusResult{}, items)
}

func TestFocusService_GetFocusTasks_Empty(t *testing.T) {
	s, err := NewService()
	if err != nil {
		t.Fatalf("failed to create service: %v", err)
	}
	ctx := context.Background()
	opts := models.FocusOptions{
		Energy:     "medium",
		Mode:       "deep",
		MaxMinutes: 3,
		MaxTasks:   10,
		Date:       time.Now().Format("2006-01-02"),
	}
	items, err := s.GetFocusTasks(ctx, opts)
	assert.NoError(t, err)
	assert.IsType(t, []models.FocusResult{}, items)
}
