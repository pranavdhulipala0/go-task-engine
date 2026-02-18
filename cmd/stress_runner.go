package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"sync/atomic"
	"time"

	"github.com/pranavdhulipala/go-task-engine/internal/engine"
	"github.com/pranavdhulipala/go-task-engine/internal/models"
)

const (
	NUM_TASKS   = 100000
	NUM_WORKERS = 500
)

type StressTestMetrics struct {
	tasksSubmitted int64
	tasksCompleted int64
	tasksFailed    int64
	tasksCancelled int64
	tasksTimedOut  int64

	// Memory and goroutine tracking
	initialGoroutines int
	finalGoroutines   int

	// For tracking unique executions
	executionCount int64
}

func (m *StressTestMetrics) IncrementSubmitted() {
	atomic.AddInt64(&m.tasksSubmitted, 1)
}

func (m *StressTestMetrics) IncrementExecuted() {
	atomic.AddInt64(&m.executionCount, 1)
}

// Simple task execution function
func createSimpleTask(taskId string, metrics *StressTestMetrics) models.Task {
	// Random duration between 1ms and 500ms
	duration := time.Duration(rand.Intn(500)+1) * time.Millisecond

	// Random failure rate (15% chance of failure)
	shouldFail := rand.Float32() < 0.15

	// Random timeout (10% chance of timeout by taking longer than duration)
	shouldTimeout := rand.Float32() < 0.10

	executeFunc := func(ctx context.Context) error {
		metrics.IncrementExecuted()

		// Simulate work time
		workDuration := duration
		if shouldTimeout {
			// Make it take 2x longer than allowed to force timeout
			workDuration = duration * 2
		}

		// Simulate work with context checking
		select {
		case <-time.After(workDuration):
			if shouldFail {
				return fmt.Errorf("simulated task failure for task %s", taskId)
			}
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return models.Task{
		ID:         taskId,
		Duration:   duration,
		Execute:    executeFunc,
		Retries:    0,
		MaxRetries: rand.Intn(3), // Random max retries 0-2
	}
}

func StressMain() {
	fmt.Printf("🚀 Starting stress test with %d tasks and %d workers\n", NUM_TASKS, NUM_WORKERS)

	metrics := &StressTestMetrics{
		initialGoroutines: runtime.NumGoroutine(),
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create task manager
	taskManager := engine.NewTaskManager(ctx, cancel, NUM_WORKERS)

	// Track submission metrics
	startTime := time.Now()

	// Submit all tasks
	fmt.Printf("📤 Submitting %d tasks...\n", NUM_TASKS)
	for i := 0; i < NUM_TASKS; i++ {
		taskId := fmt.Sprintf("task-%d", i)
		task := createSimpleTask(taskId, metrics)

		executionId := taskManager.Submit(task)
		if executionId != "Task manager is shutting down" {
			metrics.IncrementSubmitted()
		}

		// Brief pause every 5000 tasks to prevent overwhelming
		if i%5000 == 0 && i > 0 {
			fmt.Printf("  📊 Submitted %d tasks so far...\n", i)
			time.Sleep(10 * time.Millisecond)
		}
	}

	submissionTime := time.Since(startTime)
	fmt.Printf("✅ All tasks submitted in %v\n", submissionTime)

	// Wait for completion
	fmt.Println("⏳ Waiting for all tasks to complete...")
	shutdownStart := time.Now()
	taskManager.Shutdown()
	shutdownTime := time.Since(shutdownStart)

	// Get final metrics
	metrics.finalGoroutines = runtime.NumGoroutine()

	// Calculate results
	totalTime := time.Since(startTime)

	// Force garbage collection to get accurate final goroutine count
	runtime.GC()
	runtime.GC()
	time.Sleep(100 * time.Millisecond)
	finalGoroutinesAfterGC := runtime.NumGoroutine()

	// Analyze engine state
	tm := taskManager
	pendingCount := len(tm.PendingTasks)
	failedCount := len(tm.FailedTasks)

	// Print comprehensive results
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("📊 STRESS TEST RESULTS")
	fmt.Println(strings.Repeat("=", 80))

	fmt.Printf("⏱️  TIMING:\n")
	fmt.Printf("   Total execution time: %v\n", totalTime)
	fmt.Printf("   Task submission time: %v\n", submissionTime)
	fmt.Printf("   Shutdown time: %v\n", shutdownTime)
	fmt.Printf("   Tasks per second (submission): %.2f\n", float64(NUM_TASKS)/submissionTime.Seconds())
	fmt.Printf("   Tasks per second (total): %.2f\n", float64(NUM_TASKS)/totalTime.Seconds())

	fmt.Printf("\n📈 TASK METRICS:\n")
	submitted := atomic.LoadInt64(&metrics.tasksSubmitted)
	executed := atomic.LoadInt64(&metrics.executionCount)

	fmt.Printf("   Tasks submitted: %d\n", submitted)
	fmt.Printf("   Tasks executed: %d\n", executed)
	fmt.Printf("   Tasks pending: %d\n", pendingCount)
	fmt.Printf("   Tasks failed: %d\n", failedCount)

	fmt.Printf("\n🔍 ANALYSIS:\n")
	finished := executed
	leaked := submitted - finished
	stuck := int64(pendingCount)
	duplicated := executed - submitted // If positive, means duplicate executions

	fmt.Printf("   ✅ Finished (executed): %d\n", finished)
	fmt.Printf("   🚰 Leaked (submitted - finished): %d\n", leaked)
	fmt.Printf("   ⚠️  Stuck in pending: %d\n", stuck)
	fmt.Printf("   🔄 Duplicate executions: %d\n", max(0, duplicated))

	fmt.Printf("\n🧠 MEMORY & GOROUTINES:\n")
	fmt.Printf("   Initial goroutines: %d\n", metrics.initialGoroutines)
	fmt.Printf("   Final goroutines: %d\n", metrics.finalGoroutines)
	fmt.Printf("   Final goroutines (after GC): %d\n", finalGoroutinesAfterGC)
	fmt.Printf("   Goroutine leak: %d\n", finalGoroutinesAfterGC-metrics.initialGoroutines)

	// Memory stats
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("   Memory allocated: %.2f MB\n", float64(memStats.Alloc)/1024/1024)
	fmt.Printf("   Total allocations: %.2f MB\n", float64(memStats.TotalAlloc)/1024/1024)
	fmt.Printf("   GC cycles: %d\n", memStats.NumGC)

	fmt.Printf("\n📊 SUCCESS METRICS:\n")
	if submitted > 0 {
		successRate := float64(executed) * 100 / float64(submitted)
		fmt.Printf("   Execution rate: %.2f%%\n", successRate)

		if failedCount > 0 {
			fmt.Printf("   Final failure rate: %.2f%%\n", float64(failedCount)*100/float64(executed))
		}
	}

	// Final verdict
	fmt.Printf("\n🏆 ENGINE HEALTH CHECK:\n")
	goroutineLeak := finalGoroutinesAfterGC - metrics.initialGoroutines

	if leaked <= 0 {
		fmt.Printf("   ✅ No leaked tasks detected\n")
	} else {
		fmt.Printf("   ⚠️  %d tasks may be leaked\n", leaked)
	}

	if stuck == 0 {
		fmt.Printf("   ✅ No stuck tasks detected\n")
	} else {
		fmt.Printf("   ⚠️  %d tasks stuck in pending\n", stuck)
	}

	if duplicated <= 0 {
		fmt.Printf("   ✅ No duplicate executions detected\n")
	} else {
		fmt.Printf("   ⚠️  %d duplicate executions detected\n", duplicated)
	}

	if goroutineLeak <= NUM_WORKERS {
		fmt.Printf("   ✅ Goroutine usage looks healthy (leak: %d)\n", goroutineLeak)
	} else {
		fmt.Printf("   ⚠️  Potential goroutine leak: %d extra goroutines\n", goroutineLeak)
	}

	if leaked <= 0 && stuck == 0 && duplicated <= 0 && goroutineLeak <= NUM_WORKERS {
		fmt.Printf("\n🎉 CONGRATULATIONS! Your engine survived the stress test!\n")
		fmt.Printf("🏗️  You've built a production-grade task system!\n")
	} else {
		fmt.Printf("\n🔧 Your engine needs some tuning, but it's getting there!\n")

		// Provide specific recommendations
		if leaked > 0 || stuck > 0 {
			fmt.Printf("💡 Recommendation: Check task completion and WaitGroup synchronization\n")
		}
		if duplicated > 0 {
			fmt.Printf("💡 Recommendation: Review task deduplication logic\n")
		}
		if goroutineLeak > NUM_WORKERS {
			fmt.Printf("💡 Recommendation: Check for goroutine leaks in worker management\n")
		}
	}

	fmt.Println(strings.Repeat("=", 80))
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
