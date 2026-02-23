package main

import (
	"context"
	"fmt"
	"time"

	"github.com/pranavdhulipala/go-task-engine/internal/engine"
	"github.com/pranavdhulipala/go-task-engine/internal/utils"
)

func main() {
	fmt.Println("Task engine starting...")
	time.Sleep(1 * time.Second)
	startTime := time.Now()
	ctx, cancel := context.WithCancel(context.Background())
	responses := make([]string, 0)

	tm := engine.NewTaskManager(ctx, cancel, 5)
	tasks := utils.GenerateTasks(&responses) // Pass pointer to slice

	time.Sleep(2 * time.Second)

	for _, t := range tasks {
		tm.Submit(t)
	}

	// tm.Cancel("C")
	// tm.Cancel("X")

	tm.Shutdown()
	endTime := time.Now()
	fmt.Println("⏰ Total time taken:", endTime.Sub(startTime))
	tm.Mu.RLock()
	fmt.Println("Failed tasks:", len(tm.FailedTasks))
	for id, task := range tm.FailedTasks {
		fmt.Printf("  - %s (retries: %d)\n", id, task.Retries)
	}
	tm.Mu.RUnlock()

	fmt.Println("Responses:", len(responses))
}
