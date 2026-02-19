package engine

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/pranavdhulipala/go-task-engine/internal/models"
)

type TaskManager struct {
	TaskQueue      chan models.Task
	Wg             sync.WaitGroup
	TaskWg         sync.WaitGroup
	ctx            context.Context
	cancel         context.CancelFunc
	PendingTasks   map[string]models.Task
	FailedTasks    map[string]models.Task
	Mu             sync.RWMutex
	executingTasks sync.Map // map[string]*sync.Once
}

func NewTaskManager(ctx context.Context, cancel context.CancelFunc, workers int) *TaskManager {
	// Configure logrus
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	tm := &TaskManager{
		ctx:          ctx,
		cancel:       cancel,
		TaskQueue:    make(chan models.Task, 100),
		PendingTasks: make(map[string]models.Task),
		FailedTasks:  make(map[string]models.Task),
		Mu:           sync.RWMutex{},
	}

	for i := 0; i < workers; i++ {
		tm.Wg.Add(1)
		go tm.Worker(ctx, i, tm.TaskQueue, &tm.Wg)
	}

	return tm
}

func (tm *TaskManager) Submit(task models.Task) string {
	// Generate ExecutionId before sending to channel
	task.ExecutionId = uuid.New().String()
	task.State = models.StatePending

	tm.Mu.Lock()
	if _, exists := tm.PendingTasks[task.ID]; !exists {
		tm.TaskWg.Add(1) // only once per job
		tm.PendingTasks[task.ID] = task
	}
	tm.Mu.Unlock()

	tm.TaskQueue <- task
	return task.ExecutionId
}

func (tm *TaskManager) Cancel(ID string) string {
	tm.Mu.Lock()
	defer tm.Mu.Unlock()

	if _, exists := tm.PendingTasks[ID]; exists {
		delete(tm.PendingTasks, ID) // remove from pending
		logrus.WithField("taskId", ID).Info("❌ Task cancelled successfully")
		tm.TaskWg.Done() // mark as done to prevent deadlock
		return "Cancelled Task"
	}
	logrus.WithField("taskId", ID).Warn("⚠️ Task not found or already executed")
	return "Task not found or already executed"
}

func (tm *TaskManager) Shutdown() {
	tm.TaskWg.Wait()    // ⬅ wait for all tasks to be executed
	close(tm.TaskQueue) // ⬅ then close channel to signal workers
	tm.cancel()         // ⬅ cancel context (optional here)
	tm.Wg.Wait()        // ⬅ wait for workers to exit
	logrus.Info("🛑 Task Manager shutting down")
}

func (tm *TaskManager) Worker(ctx context.Context, workerId int, queue chan models.Task, wg *sync.WaitGroup) {
	defer wg.Done()
	logrus.WithField("workerId", workerId).Info("🚀 Worker started")

	for {
		select {
		case task, ok := <-tm.TaskQueue:
			if !ok {
				logrus.WithField("workerId", workerId).Info("🔌 Worker shutting down - channel closed")
				return
			}
			tm.ExecuteTask(ctx, task, workerId)

		case <-tm.ctx.Done():
			logrus.WithField("workerId", workerId).Info("💤 Worker shutting down - context cancelled")
			return
		}
	}
}

func (tm *TaskManager) ExecuteTask(ctx context.Context, task models.Task, workerId int) {
	// Use sync.Once to ensure each task ExecutionId executes only once
	once, _ := tm.executingTasks.LoadOrStore(task.ExecutionId, &sync.Once{})
	taskOnce := once.(*sync.Once) //This is a type assertion from any to sync.Once

	taskOnce.Do(func() {
		//Mark this task as done at the end of this execution
		defer tm.TaskWg.Done()
		defer tm.executingTasks.Delete(task.ExecutionId)
		// Remove task from pending tasks
		tm.Mu.Lock()
		task.State = models.StateInProgress
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
					tm.RetryTask(task)
					return
				}
				tm.AddTaskToQueue(task, tm.FailedTasks)
				tm.Mu.Lock()
				task.State = models.StateFailed
				tm.Mu.Unlock()
				logrus.WithFields(logrus.Fields{
					"workerId":    workerId,
					"taskId":      task.ID,
					"executionId": task.ExecutionId,
					"retries":     task.Retries,
					"error":       err,
				}).Error("🔥 Task failed")
			} else {
				tm.Mu.Lock()
				task.State = models.StateCompleted
				tm.Mu.Unlock()
				logrus.WithFields(logrus.Fields{
					"workerId":    workerId,
					"taskId":      task.ID,
					"executionId": task.ExecutionId,
				}).Info("✅ Task completed successfully")
			}

		case <-taskCtx.Done():
			//Retry the Task if it has not reached the max retries
			if task.Retries < task.MaxRetries {
				tm.RetryTask(task)
				return
			}
			tm.AddTaskToQueue(task, tm.FailedTasks)
			tm.Mu.Lock()
			task.State = models.StateFailed
			tm.Mu.Unlock()
			logrus.WithFields(logrus.Fields{
				"workerId":    workerId,
				"taskId":      task.ID,
				"executionId": task.ExecutionId,
				"retries":     task.Retries,
			}).Warn("⏰ Task timed out")
		}
	})
}

func (tm *TaskManager) RetryTask(task models.Task) {
	task.Retries++
	logrus.WithFields(logrus.Fields{
		"taskId":     task.ID,
		"retryCount": task.Retries,
		"maxRetries": task.MaxRetries,
	}).Info("😩 Retrying task")
	tm.Submit(task) // This will get a NEW ExecutionId
}

func (tm *TaskManager) AddTaskToQueue(task models.Task, queue map[string]models.Task) {
	tm.Mu.Lock()
	queue[task.ID] = task
	tm.Mu.Unlock()
}
