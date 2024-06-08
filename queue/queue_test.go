package queue

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestAddTask(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	ctx := context.Background()
	queue := NewQueue(ctx, logger)

	exec := func() (interface{}, error) {
		return "result", nil
	}
	result := queue.AddTask(exec)

	assert.NotNil(t, result, "expected result to be non-nil")
	assert.Equal(t, Pending, result.Status, "expected status to be pending")
}

func TestGetResult_Success(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	ctx := context.Background()
	queue := NewQueue(ctx, logger)

	exec := func() (interface{}, error) {
		time.Sleep(100 * time.Millisecond)
		return "result", nil
	}
	result := queue.AddTask(exec)

	time.Sleep(200 * time.Millisecond) // Wait for the task to complete

	taskResult, err := queue.GetResult(result.ID)
	assert.NoError(t, err, "unexpected error")
	assert.Equal(t, Success, taskResult.Status, "expected status to be success")
	assert.Equal(t, "result", taskResult.Result, "expected result to be 'result'")
}

func TestGetResult_Error(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	ctx := context.Background()
	queue := NewQueue(ctx, logger)

	exec := func() (interface{}, error) {
		time.Sleep(100 * time.Millisecond)
		return nil, errors.New("task error")
	}
	result := queue.AddTask(exec)

	time.Sleep(200 * time.Millisecond) // Wait for the task to complete

	taskResult, err := queue.GetResult(result.ID)
	assert.NoError(t, err, "unexpected error")
	assert.Equal(t, Error, taskResult.Status, "expected status to be error")
	assert.EqualError(t, taskResult.Error, "task error", "expected error to be 'task error'")
}

func TestCleanupOldResults(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	ctx := context.Background()
	queue := NewQueue(ctx, logger)

	exec := func() (interface{}, error) {
		return "result", nil
	}
	result := queue.AddTask(exec)

	time.Sleep(200 * time.Millisecond) // Wait for the task to complete

	// Manually set the completion time to more than 5 minutes ago
	queue.mu.Lock()
	for i, task := range queue.tasks {
		if task.ID == result.ID {
			queue.tasks[i].CompletionTime = time.Now().Add(-10 * time.Minute)
		}
	}
	queue.mu.Unlock()

	queue.cleanupOldResults()

	taskResult, err := queue.GetResult(result.ID)
	assert.NoError(t, err, "unexpected error")
	assert.Equal(t, Archived, taskResult.Status, "expected status to be archived")
	assert.Nil(t, taskResult.Result, "expected result to be nil")
}
