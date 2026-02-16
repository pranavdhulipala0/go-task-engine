package engine

import "time"

type Task struct {
	ID       string
	Duration time.Duration
	Execute  func() error
}
