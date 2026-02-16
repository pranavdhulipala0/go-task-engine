package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/pranavdhulipala/go-task-engine/internal/engine"
)

func main() {
	fmt.Println("Task engine starting...")
	time.Sleep(2 * time.Second)
	ctx, cancel := context.WithCancel(context.Background())

	tm := engine.NewTaskManager(ctx, cancel, 5)

	tasks := []engine.Task{
		{
			ID: "A",
			Execute: func() error {
				time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
				fmt.Println("Executing task A")
				return nil
			},
		},
		{
			ID: "B",
			Execute: func() error {
				time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
				fmt.Println("Executing task B")
				return nil
			},
		},
		{
			ID: "C",
			Execute: func() error {
				time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
				fmt.Println("Executing task C")
				return nil
			},
		},
		{
			ID: "D",
			Execute: func() error {
				time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
				fmt.Println("Executing task D")
				return nil
			},
		},
	}

	time.Sleep(2 * time.Second)

	for _, t := range tasks {
		tm.Submit(t)
	}

	tm.Cancel("C")
	tm.Cancel("X")

	tm.Shutdown()
}
