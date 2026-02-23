package engine

import (
	"fmt"
	"time"

	"github.com/pranavdhulipala/go-task-engine/internal/models"
)

type PriorityQueue []models.Task

func (pq PriorityQueue) Len() int { return len(pq) }

// Aging-based comparison
func (pq PriorityQueue) Less(i, j int) bool {
	now := time.Now()

	waitedI := int(now.Sub(pq[i].CreatedAt).Milliseconds())
	waitedJ := int(now.Sub(pq[j].CreatedAt).Milliseconds())

	scoreI := pq[i].Priority*1000 + waitedI
	scoreJ := pq[j].Priority*1000 + waitedJ

	return scoreI > scoreJ // max-heap behavior
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	task := x.(models.Task)
	*pq = append(*pq, task)
	// fmt.Printf("PQ: Added task %s (priority %d) to queue. Queue now has %d tasks\n", task.ID, task.Priority, len(*pq))
	// pq.printQueue("After Push")
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	task := old[n-1]
	*pq = old[:n-1]
	// fmt.Printf("PQ: Popped task %s (priority %d) from queue. Queue now has %d tasks\n", task.ID, task.Priority, len(*pq))
	// pq.printQueue("After Pop")
	return task
}

func (pq *PriorityQueue) PrintQueue(operation string) {
	fmt.Printf("PQ: %s - Current queue state: [", operation)
	for i, task := range *pq {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%s(P%d)", task.ID, task.Priority)
	}
	fmt.Println("]")
}
