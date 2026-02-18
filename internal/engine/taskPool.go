package engine

import (
	"context"
	"io"
	"net/http"
	"sync"
	"time"
)

func GenerateTasks(responses *[]string) []Task {
	var mu sync.Mutex // Mutex to protect slice access

	return []Task{
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
