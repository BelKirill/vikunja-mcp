package focus

import (
	"context"
	"testing"
	"time"

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
	opts := FocusOptions{
		Energy:   "low",
		Mode:     "",
		Hours:    5,
		MaxItems: 5,
		Date:     time.Now().String(),
	}

	items, _ := s.GetFocusTasks(ctx, opts)
	assert.IsType(t, FocusResult{}, items)
}

func TestFocusService_GetFocusTasks_Empty(t *testing.T) {
	s, err := NewService()
	if err != nil {
		t.Fatalf("failed to create service: %v", err)
	}
	ctx := context.Background()
	opts := FocusOptions{
		Energy:   "medium",
		Mode:     "deep",
		Hours:    3,
		MaxItems: 10,
		Date:     time.Now().Format("2006-01-02"),
	}
	items, err := s.GetFocusTasks(ctx, opts)
	assert.NoError(t, err)
	assert.IsType(t, []FocusResult{}, items)
}

func TestFocusService_ParseHyperFocusMetadata(t *testing.T) {
	s, err := NewService()
	if err != nil {
		t.Fatalf("failed to create service: %v", err)
	}
	meta := s.parseHyperFocusMetadata("some description")
	assert.Nil(t, meta)
}

func TestFocusService_CleanDescription(t *testing.T) {
	s, err := NewService()
	if err != nil {
		t.Fatalf("failed to create service: %v", err)
	}
	input := "desc"
	out := s.cleanDescription(input)
	assert.Equal(t, input, out)
}
