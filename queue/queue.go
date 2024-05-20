package queue

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

// TaskStatus represents the status of a task.
type TaskStatus string

const (
	Pending  TaskStatus = "pending"
	Success  TaskStatus = "success"
	Error    TaskStatus = "error"
	Archived TaskStatus = "archived"
)

// Task represents a single task with an ID, a function to execute, status, error, and completion time.
type Task struct {
	ID             string
	Exec           func() (interface{}, error)
	Status         TaskStatus
	Error          error
	CompletionTime time.Time
}

type TaskResult struct {
	ID             string
	Status         TaskStatus
	Error          error
	CompletionTime time.Time
	Result         interface{}
}

// Queue represents a task queue.
type Queue struct {
	tasks   []Task
	results map[string]interface{}
	mu      sync.Mutex
	wg      sync.WaitGroup
	ctx     context.Context
	log     *zap.Logger
}

// NewQueue creates a new Queue.
func NewQueue(ctx context.Context, log *zap.Logger) *Queue {
	queue := &Queue{
		tasks:   []Task{},
		results: make(map[string]interface{}),
		ctx:     ctx,
		log:     log,
	}
	go queue.cleanupWorker()
	return queue
}

// AddTask adds a task to the queue and runs it.
func (q *Queue) AddTask(exec func() (interface{}, error)) *TaskResult {
	q.mu.Lock()
	defer q.mu.Unlock()

	taskID := uuid.New().String()
	task := Task{
		ID:     taskID,
		Exec:   exec,
		Status: Pending,
	}
	q.tasks = append(q.tasks, task)

	// Run the task immediately in a separate goroutine
	go q.runTask(task)

	return &TaskResult{
		ID:     task.ID,
		Status: task.Status,
	}
}

// runTask runs a single task.
func (q *Queue) runTask(task Task) {
	q.wg.Add(1)
	defer q.wg.Done()

	result, err := task.Exec()
	q.mu.Lock()
	defer q.mu.Unlock()

	if err != nil {
		task.Status = Error
		task.Error = err
	} else {
		task.Status = Success
		task.CompletionTime = time.Now()
		q.results[task.ID] = result
	}

	// Update task status in the queue
	for i, existingTask := range q.tasks {
		if existingTask.ID == task.ID {
			q.tasks[i] = task
			break
		}
	}
}

// GetResult retrieves the result and status of a task by its ID.
func (q *Queue) GetResult(id string) (*TaskResult, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	for _, task := range q.tasks {
		if task.ID == id {
			return &TaskResult{
				ID:             task.ID,
				Status:         task.Status,
				Error:          task.Error,
				Result:         q.results[id],
				CompletionTime: task.CompletionTime,
			}, nil
		}
	}
	return nil, errors.New("task not found")
}

// cleanupWorker periodically cleans up old task results.
func (q *Queue) cleanupWorker() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			q.cleanupOldResults()
		case <-q.ctx.Done():
			return
		}
	}
}

// cleanupOldResults removes results older than 5 minutes and updates their status.
func (q *Queue) cleanupOldResults() {
	q.mu.Lock()
	defer q.mu.Unlock()

	now := time.Now()
	for _, task := range q.tasks {
		if task.Status == Success && now.Sub(task.CompletionTime) > 5*time.Minute {
			delete(q.results, task.ID)
			task.Status = Archived
			// Update task status in the queue
			for i, existingTask := range q.tasks {
				if existingTask.ID == task.ID {
					q.tasks[i].Status = Archived
					break
				}
			}
		}
	}
}
