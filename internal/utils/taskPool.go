package utils

import (
	"context"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/pranavdhulipala/go-task-engine/internal/models"
)

func GenerateTasks(responses *[]string) []models.Task {
	var mu sync.Mutex // Mutex to protect slice access

	return []models.Task{
		{
			ID:         "A",
			Duration:   5 * time.Second,
			Retries:    0,
			MaxRetries: 3,
			Priority:   1,
			Execute: func(ctx context.Context) error {
				time.Sleep(500 * time.Millisecond)
				return nil
			},
		},
		{
			ID:         "B",
			Duration:   3 * time.Second,
			Retries:    0,
			MaxRetries: 3,
			Priority:   10,
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
				mu.Lock()
				*responses = append(*responses, string(body))
				mu.Unlock()

				return nil
			},
		},
		{
			ID:         "C",
			Duration:   10 * time.Second,
			Retries:    0,
			MaxRetries: 3,
			Priority:   5,
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
				mu.Lock()
				*responses = append(*responses, string(body))
				mu.Unlock()

				return nil
			},
		},
		{
			ID:         "D",
			Duration:   4 * time.Second,
			Retries:    0,
			MaxRetries: 3,
			Priority:   8,
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
				mu.Lock()
				*responses = append(*responses, string(body))
				mu.Unlock()

				return nil
			},
		},
		{
			ID:         "E",
			Duration:   4 * time.Second,
			Retries:    0,
			MaxRetries: 3,
			Priority:   2,
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
				mu.Lock()
				*responses = append(*responses, string(body))
				mu.Unlock()

				return nil
			},
		},
		{
			ID:         "F",
			Duration:   4 * time.Second,
			Retries:    0,
			MaxRetries: 3,
			Priority:   9,
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
				mu.Lock()
				*responses = append(*responses, string(body))
				mu.Unlock()

				return nil
			},
		},
		{
			ID:         "G",
			Duration:   4 * time.Second,
			Retries:    0,
			MaxRetries: 3,
			Priority:   3,
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
				mu.Lock()
				*responses = append(*responses, string(body))
				mu.Unlock()

				return nil
			},
		},
		{
			ID:         "H",
			Duration:   4 * time.Second,
			Retries:    0,
			MaxRetries: 3,
			Priority:   7,
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
				mu.Lock()
				*responses = append(*responses, string(body))
				mu.Unlock()

				return nil
			},
		},
		{
			ID:         "I",
			Duration:   4 * time.Second,
			Retries:    0,
			MaxRetries: 3,
			Priority:   4,
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
				mu.Lock()
				*responses = append(*responses, string(body))
				mu.Unlock()

				return nil
			},
		},
		{
			ID:         "L",
			Duration:   4 * time.Second,
			Retries:    0,
			MaxRetries: 3,
			Priority:   6,
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
				mu.Lock()
				*responses = append(*responses, string(body))
				mu.Unlock()

				return nil
			},
		},
		{
			ID:         "J",
			Duration:   4 * time.Second,
			Retries:    0,
			MaxRetries: 3,
			Priority:   1,
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
				mu.Lock()
				*responses = append(*responses, string(body))
				mu.Unlock()

				return nil
			},
		},
		{
			ID:         "K",
			Duration:   4 * time.Second,
			Retries:    0,
			MaxRetries: 3,
			Priority:   10,
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
				mu.Lock()
				*responses = append(*responses, string(body))
				mu.Unlock()

				return nil
			},
		},
	}
}
