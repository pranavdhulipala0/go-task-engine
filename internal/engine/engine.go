package engine

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type TaskManager struct {
	TaskQueue      chan Task
	Wg             sync.WaitGroup
	TaskWg         sync.WaitGroup
	ctx            context.Context
	cancel         context.CancelFunc
	PendingTasks   map[string]Task
	FailedTasks    map[string]Task
	Mu             sync.RWMutex
	executingTasks sync.Map // map[string]*sync.Once
}

func NewTaskManager(ctx context.Context, cancel context.CancelFunc, workers int) *TaskManager {
	tm := &TaskManager{
		ctx:          ctx,
		cancel:       cancel,
		TaskQueue:    make(chan Task),
		PendingTasks: make(map[string]Task),
		FailedTasks:  make(map[string]Task),
		Mu:           sync.RWMutex{},
	}

	for i := 0; i < workers; i++ {
		tm.Wg.Add(1)
		go tm.Worker(ctx, i, tm.TaskQueue, &tm.Wg)
	}

	return tm
}

func (tm *TaskManager) Submit(task Task) string {
	// Generate ExecutionId before sending to channel
	task.ExecutionId = uuid.New().String()

	tm.TaskWg.Add(1) // increment BEFORE sending
	select {
	case tm.TaskQueue <- task:
		tm.AddTaskToQueue(task, tm.PendingTasks)
		fmt.Println("✅ Submitted Task:", task.ID, "\nExecutionId:", task.ExecutionId)
		return task.ExecutionId
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
			tm.ExecuteTask(ctx, task, workerId)

		case <-tm.ctx.Done():
			fmt.Println("💤 Worker", workerId, "shutting down")
			return
		}
	}
}

func (tm *TaskManager) ExecuteTask(ctx context.Context, task Task, workerId int) {
	//Mark this task as done at the end of this execution
	defer tm.TaskWg.Done()

	// Use sync.Once to ensure each task ExecutionId executes only once
	once, _ := tm.executingTasks.LoadOrStore(task.ExecutionId, &sync.Once{})
	taskOnce := once.(*sync.Once) //This is a type assertion from any to sync.Once

	taskOnce.Do(func() {
		// Remove task from pending tasks
		tm.Mu.Lock()
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
				//Retry the Task if it has not reached the max retries
				if task.Retries < task.MaxRetries {
					task.Retries++
					tm.Submit(task)
					return
				}
				tm.AddTaskToQueue(task, tm.FailedTasks)
				fmt.Println("🔥 WORKER:", workerId, " Task failed:", task.ID, "\nExecutionId:", task.ExecutionId, err)
			} else {
				fmt.Println("✅ WORKER:", workerId, " Task completed:", task.ID, "\nExecutionId:", task.ExecutionId)
			}

		case <-taskCtx.Done():
			//Retry the Task if it has not reached the max retries
			if task.Retries < task.MaxRetries {
				tm.RetryTask(task)
				return
			}
			tm.AddTaskToQueue(task, tm.FailedTasks)
			fmt.Println("⏰ WORKER:", workerId, " Task timed out:", task.ID, "ExecutionId:", task.ExecutionId)
		}
	})
}

func (tm *TaskManager) RetryTask(task Task) {
	task.Retries++
	fmt.Println("😩 Retrying task:", task.ID, "for the", task.Retries, "time")
	tm.Submit(task) // This will get a NEW ExecutionId
}

func (tm *TaskManager) AddTaskToQueue(task Task, queue map[string]Task) {
	tm.Mu.Lock()
	queue[task.ID] = task
	tm.Mu.Unlock()
}
