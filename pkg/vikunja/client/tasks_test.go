package client

import (
	"context"
	"errors"
	"testing"

	"github.com/BelKirill/vikunja-mcp/models"
	"github.com/stretchr/testify/assert"
)

// mockClient is a minimal mock of Client for testing.
type mockClient struct {
	Client
	getFunc func(ctx context.Context, endpoint string, out interface{}) error
}

func (m *mockClient) Get(ctx context.Context, endpoint string, out interface{}) error {
	return m.getFunc(ctx, endpoint, out)
}

func TestGetAllTasks_Success(t *testing.T) {
	expected := []models.RawTask{
		{ID: 1, Title: "Task 1"},
		{ID: 2, Title: "Task 2"},
	}
	mc := &mockClient{
		getFunc: func(ctx context.Context, endpoint string, out interface{}) error {
			tasks := out.(*[]models.RawTask)
			*tasks = expected
			return nil
		},
	}

	ctx := context.Background()
	tasks, err := mc.GetAllTasks(ctx)
	assert.NoError(t, err)
	assert.Equal(t, expected, tasks)
}

func TestGetAllTasks_Error(t *testing.T) {
	mc := &mockClient{
		getFunc: func(ctx context.Context, endpoint string, out interface{}) error {
			return errors.New("api error")
		},
	}

	ctx := context.Background()
	tasks, err := mc.GetAllTasks(ctx)
	assert.Error(t, err)
	assert.Nil(t, tasks)
}

func TestGetTask_Success(t *testing.T) {
	expected := models.RawTask{ID: 42, Title: "Test Task"}
	mc := &mockClient{
		getFunc: func(ctx context.Context, endpoint string, out interface{}) error {
			task := out.(*models.RawTask)
			*task = expected
			return nil
		},
	}

	ctx := context.Background()
	task, err := mc.GetTask(ctx, 42)
	assert.NoError(t, err)
	assert.Equal(t, &expected, task)
}

func TestGetTask_Error(t *testing.T) {
	mc := &mockClient{
		getFunc: func(ctx context.Context, endpoint string, out interface{}) error {
			return errors.New("not found")
		},
	}
	ctx := context.Background()
	task, err := mc.GetTask(ctx, 99)
	assert.Error(t, err)
	assert.Nil(t, task)
}
