package review

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReviewService_Review_Empty(t *testing.T) {
	// Example: test Review with empty input (mock or minimal implementation)
	var s Service // replace with actual service if needed
	resp, err := s.Review(nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
}
