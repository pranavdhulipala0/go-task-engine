package engine

import (
	"context"
	"fmt"
	"sync"
)

type TaskManager struct {
	TaskQueue    chan Task
	Wg           sync.WaitGroup
	TaskWg       sync.WaitGroup
	ctx          context.Context
	cancel       context.CancelFunc
	PendingTasks map[string]Task
	Mu           sync.RWMutex
}

func NewTaskManager(ctx context.Context, cancel context.CancelFunc, workers int) *TaskManager {
	tm := &TaskManager{
		ctx:          ctx,
		cancel:       cancel,
		TaskQueue:    make(chan Task),
		PendingTasks: make(map[string]Task),
		Mu:           sync.RWMutex{},
	}

	for i := 0; i < workers; i++ {
		tm.Wg.Add(1)
		go tm.Worker(ctx, i, tm.TaskQueue, &tm.Wg)
	}

	return tm
}

func (tm *TaskManager) Submit(task Task) string {
	tm.TaskWg.Add(1) // increment BEFORE sending
	select {
	case tm.TaskQueue <- task:
		tm.Mu.Lock()
		tm.PendingTasks[task.ID] = task
		tm.Mu.Unlock()
		fmt.Println("✅ Submitted Task:", task.ID)
		return task.ID
	case <-tm.ctx.Done():
		tm.TaskWg.Done()
		return "Task manager is shutting down"
	}
}

func (tm *TaskManager) Cancel(ID string) string {
	tm.Mu.Lock()
	defer tm.Mu.Unlock()

	if _, exists := tm.PendingTasks[ID]; exists {
		delete(tm.PendingTasks, ID) // remove from pending
		fmt.Println("❌ Cancelled Task:", ID)
		tm.TaskWg.Done() // mark as done to prevent deadlock
		return "Cancelled Task"
	}
	fmt.Println("❌ Task not found or already executed:", ID)
	return "Task not found or already executed"
}

func (tm *TaskManager) Shutdown() {
	tm.TaskWg.Wait()    // ⬅ wait for all tasks to be executed
	close(tm.TaskQueue) // ⬅ then close channel to signal workers
	tm.cancel()         // ⬅ cancel context (optional here)
	tm.Wg.Wait()        // ⬅ wait for workers to exit
	fmt.Println("🛑 Shutting down Task Manager...")
}

func (tm *TaskManager) Worker(ctx context.Context, workerId int, queue chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("🚀 Worker:", workerId, "started")

	for {
		select {
		case task, ok := <-tm.TaskQueue:
			if !ok {
				fmt.Println("🔌 Channel closed, worker", workerId, "shutting down")
				return
			}
			tm.ExecuteTask(ctx, task)

		case <-tm.ctx.Done():
			fmt.Println("💤 Worker", workerId, "shutting down")
			return
		}
	}
}

func (tm *TaskManager) ExecuteTask(ctx context.Context, task Task) {

	//Lock and check if the task exists.
	tm.Mu.Lock()
	_, exists := tm.PendingTasks[task.ID]
	if !exists { // task already cancelled or executed
		fmt.Println("❌ Task already executed or cancelled:", task.ID)
		tm.Mu.Unlock()
		return
	}
	delete(tm.PendingTasks, task.ID)
	tm.Mu.Unlock()

	//Create a Timeout Context for the Task.
	taskCtx, cancel := context.WithTimeout(ctx, task.Duration)
	defer cancel()

	//Create a buffered channel to store status of the task execution
	done := make(chan error, 1)

	//Run the function in a Go Routine
	go func() {
		done <- task.Execute(taskCtx)
	}()

	//Listen to the Status Channel for each Task -> If the task fails, print the error, if it times out, print timeout error
	select {
	case err := <-done:
		if err != nil {
			fmt.Println("🔥 Task failed:", task.ID, err)
		} else {
			fmt.Println("✅ Task completed:", task.ID)
		}

	case <-taskCtx.Done():
		fmt.Println("⏰ Task timed out:", task.ID)
	}

	//Mark this task as done
	tm.TaskWg.Done()
}
