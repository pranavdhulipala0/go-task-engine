package main

import (
	"context"
	"fmt"
	"time"

	"github.com/pranavdhulipala/go-task-engine/internal/engine"
)

func main() {
	fmt.Println("Task engine starting...")
	time.Sleep(2 * time.Second)
	startTime := time.Now()
	ctx, cancel := context.WithCancel(context.Background())

	tm := engine.NewTaskManager(ctx, cancel, 5)

	tasks := []engine.Task{
		{
			ID: "A",
			Execute: func(ctx context.Context) error {
				for {
					select {
					case <-ctx.Done():
						return ctx.Err()
					default:
						time.Sleep(500 * time.Millisecond)
						fmt.Println("Executing task A")
					}
				}
			},
		},
		{
			ID: "B",
			Execute: func(ctx context.Context) error {
				for {
					select {
					case <-ctx.Done():
						return ctx.Err()
					default:
						time.Sleep(500 * time.Millisecond)
						fmt.Println("Executing task B")
					}
				}
			},
		},
		{
			ID: "C",
			Execute: func(ctx context.Context) error {
				for {
					select {
					case <-ctx.Done():
						return ctx.Err()
					default:
						time.Sleep(100000 * time.Millisecond)
						fmt.Println("Executing task C")
					}
				}
			},
		},
		{
			ID: "D",
			Execute: func(ctx context.Context) error {
				for {
					select {
					case <-ctx.Done():
						return ctx.Err()
					default:
						time.Sleep(500 * time.Millisecond)
						fmt.Println("Executing task D")
					}
				}
			},
		},
	}

	time.Sleep(2 * time.Second)

	for _, t := range tasks {
		tm.Submit(t)
	}

	// tm.Cancel("C")
	// tm.Cancel("X")

	tm.Shutdown()
	endTime := time.Now()
	fmt.Println("⏰ Total time taken:", endTime.Sub(startTime))
}
