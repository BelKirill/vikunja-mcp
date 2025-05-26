package peek

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskService_Peek_Empty(t *testing.T) {
	// Example: test Peek with empty input (mock or minimal implementation)
	var s stubTaskService // replace with actual service if needed
	resp, err := s.Peek("", nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
}
