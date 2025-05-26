package review

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReviewService_Review_Empty(t *testing.T) {
	// Example: test Review with empty input (mock or minimal implementation)
	var s stubReviewService // replace with actual service if needed
	resp, _ := s.Review(nil)
	assert.IsType(t, []ReviewItem{}, resp)
}
