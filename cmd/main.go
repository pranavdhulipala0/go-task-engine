package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/pranavdhulipala/go-task-engine/internal/engine"
)

func main() {
	fmt.Println("Task engine starting...")
	time.Sleep(2 * time.Second)
	startTime := time.Now()
	ctx, cancel := context.WithCancel(context.Background())
	responses := make([]string, 0)

	tm := engine.NewTaskManager(ctx, cancel, 5)

	tasks := []engine.Task{
		{
			ID:         "A",
			Duration:   5 * time.Second,
			Retries:    0,
			MaxRetries: 3,
			Execute: func(ctx context.Context) error {
				time.Sleep(1000 * time.Second)
				return nil
			},
		},
		{
			ID:         "B",
			Duration:   3 * time.Second,
			Retries:    0,
			MaxRetries: 3,
			Execute: func(ctx context.Context) error {
				req, err := http.NewRequestWithContext(
					ctx,
					"GET",
					"https://jsonplaceholder.typicode.com/todos/1",
					nil,
				)
				if err != nil {
					return err
				}

				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					return err
				}
				defer resp.Body.Close()

				body, _ := io.ReadAll(resp.Body)
				responses = append(responses, string(body))

				return nil
			},
		},
		{
			ID:         "C",
			Duration:   10 * time.Second,
			Retries:    0,
			MaxRetries: 3,
			Execute: func(ctx context.Context) error {
				req, err := http.NewRequestWithContext(
					ctx,
					"GET",
					"https://jsonplaceholder.typicode.com/todos/1",
					nil,
				)
				if err != nil {
					return err
				}

				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					return err
				}
				defer resp.Body.Close()

				body, _ := io.ReadAll(resp.Body)
				responses = append(responses, string(body))

				return nil
			},
		},
		{
			ID:         "D",
			Duration:   4 * time.Second,
			Retries:    0,
			MaxRetries: 3,
			Execute: func(ctx context.Context) error {
				req, err := http.NewRequestWithContext(
					ctx,
					"GET",
					"https://jsonplaceholder.typicode.com/todos/1",
					nil,
				)
				if err != nil {
					return err
				}

				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					return err
				}
				defer resp.Body.Close()

				body, _ := io.ReadAll(resp.Body)
				responses = append(responses, string(body))

				return nil
			},
		},
		{
			ID:         "E",
			Duration:   4 * time.Second,
			Retries:    0,
			MaxRetries: 3,
			Execute: func(ctx context.Context) error {
				req, err := http.NewRequestWithContext(
					ctx,
					"GET",
					"https://jsonplaceholder.typicode.com/todos/1",
					nil,
				)
				if err != nil {
					return err
				}

				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					return err
				}
				defer resp.Body.Close()

				body, _ := io.ReadAll(resp.Body)
				responses = append(responses, string(body))

				return nil
			},
		},
		{
			ID:         "F",
			Duration:   4 * time.Second,
			Retries:    0,
			MaxRetries: 3,
			Execute: func(ctx context.Context) error {
				req, err := http.NewRequestWithContext(
					ctx,
					"GET",
					"https://jsonplaceholder.typicode.com/todos/1",
					nil,
				)
				if err != nil {
					return err
				}

				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					return err
				}
				defer resp.Body.Close()

				body, _ := io.ReadAll(resp.Body)
				responses = append(responses, string(body))

				return nil
			},
		},
		{
			ID:         "G",
			Duration:   4 * time.Second,
			Retries:    0,
			MaxRetries: 3,
			Execute: func(ctx context.Context) error {
				req, err := http.NewRequestWithContext(
					ctx,
					"GET",
					"https://jsonplaceholder.typicode.com/todos/1",
					nil,
				)
				if err != nil {
					return err
				}

				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					return err
				}
				defer resp.Body.Close()

				body, _ := io.ReadAll(resp.Body)
				responses = append(responses, string(body))

				return nil
			},
		},
		{
			ID:         "H",
			Duration:   4 * time.Second,
			Retries:    0,
			MaxRetries: 3,
			Execute: func(ctx context.Context) error {
				req, err := http.NewRequestWithContext(
					ctx,
					"GET",
					"https://jsonplaceholder.typicode.com/todos/1",
					nil,
				)
				if err != nil {
					return err
				}

				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					return err
				}
				defer resp.Body.Close()

				body, _ := io.ReadAll(resp.Body)
				responses = append(responses, string(body))

				return nil
			},
		},
		{
			ID:         "I",
			Duration:   4 * time.Second,
			Retries:    0,
			MaxRetries: 3,
			Execute: func(ctx context.Context) error {
				req, err := http.NewRequestWithContext(
					ctx,
					"GET",
					"https://jsonplaceholder.typicode.com/todos/1",
					nil,
				)
				if err != nil {
					return err
				}

				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					return err
				}
				defer resp.Body.Close()

				body, _ := io.ReadAll(resp.Body)
				responses = append(responses, string(body))

				return nil
			},
		},
		{
			ID:         "L",
			Duration:   4 * time.Second,
			Retries:    0,
			MaxRetries: 3,
			Execute: func(ctx context.Context) error {
				req, err := http.NewRequestWithContext(
					ctx,
					"GET",
					"https://jsonplaceholder.typicode.com/todos/1",
					nil,
				)
				if err != nil {
					return err
				}

				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					return err
				}
				defer resp.Body.Close()

				body, _ := io.ReadAll(resp.Body)
				responses = append(responses, string(body))

				return nil
			},
		},
		{
			ID:         "J",
			Duration:   4 * time.Second,
			Retries:    0,
			MaxRetries: 3,
			Execute: func(ctx context.Context) error {
				req, err := http.NewRequestWithContext(
					ctx,
					"GET",
					"https://jsonplaceholder.typicode.com/todos/1",
					nil,
				)
				if err != nil {
					return err
				}

				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					return err
				}
				defer resp.Body.Close()

				body, _ := io.ReadAll(resp.Body)
				responses = append(responses, string(body))

				return nil
			},
		},
		{
			ID:         "K",
			Duration:   4 * time.Second,
			Retries:    0,
			MaxRetries: 3,
			Execute: func(ctx context.Context) error {
				req, err := http.NewRequestWithContext(
					ctx,
					"GET",
					"https://jsonplaceholder.typicode.com/todos/1",
					nil,
				)
				if err != nil {
					return err
				}

				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					return err
				}
				defer resp.Body.Close()

				body, _ := io.ReadAll(resp.Body)
				responses = append(responses, string(body))

				return nil
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
	tm.Mu.RLock()
	fmt.Println("Failed tasks:", len(tm.FailedTasks))
	tm.Mu.RUnlock()

	fmt.Println("Responses:", responses)
}
