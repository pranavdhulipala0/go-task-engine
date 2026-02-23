package models

import (
	"context"
	"time"
)

type Task struct {
	ID          string
	Duration    time.Duration
	Execute     func(context.Context) error
	Retries     int
	MaxRetries  int
	ExecutionId string
	State       TaskState
	Priority    int
	CreatedAt   time.Time
}

type TaskState string

const (
	StatePending    TaskState = "PENDING"
	StateInProgress TaskState = "IN_PROGRESS"
	StateCompleted  TaskState = "COMPLETED"
	StateFailed     TaskState = "FAILED"
	StateCancelled  TaskState = "CANCELLED"
)
