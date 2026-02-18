package engine

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
}
