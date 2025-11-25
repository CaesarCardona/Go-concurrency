package main

import (
	"errors"
	"fmt"
	"time"
)

type Task struct {
	ID int
}

type Result struct {
	TaskID int
	Msg    string
	Err    error
}

func worker(tasks <-chan Task, results chan<- Result, maxTasks int, maxDuration time.Duration) {
	processed := 0
	for task := range tasks {
		if processed >= maxTasks {
			results <- Result{
				TaskID: task.ID,
				Err:    errors.New("max_tasks limit exceeded"),
			}
			continue
		}

		// Simulate work duration based on task ID and maxDuration
		var workDuration time.Duration
		if task.ID%2 == 0 {
			// Force some tasks to exceed maxDuration
			workDuration = maxDuration + 50*time.Millisecond
		} else {
			workDuration = maxDuration - 50*time.Millisecond
		}

		time.Sleep(workDuration)

		var err error
		if workDuration > maxDuration {
			err = fmt.Errorf("task %d exceeded max duration (%v)", task.ID, maxDuration)
		}

		results <- Result{
			TaskID: task.ID,
			Msg:    fmt.Sprintf("Processed task %d in %v", task.ID, workDuration),
			Err:    err,
		}

		processed++
	}
	close(results)
}


func main() {
	maxTasks := 5
	numTasks := 8
	maxDuration := 250 * time.Millisecond

	fmt.Println("=== CONFIGURATION ===")
	fmt.Println("Max tasks:", maxTasks)
	fmt.Println("Number of tasks to send:", numTasks)
	fmt.Println("Max allowed duration per task:", maxDuration)
	fmt.Println("=====================")

	tasks := make(chan Task, numTasks)
	results := make(chan Result, numTasks)

	// Start single worker
	go worker(tasks, results, maxTasks, maxDuration)

	// Send tasks
	for i := 1; i <= numTasks; i++ {
		tasks <- Task{ID: i}
	}
	close(tasks)

	// Read results
	for r := range results {
		if r.Err != nil {
			fmt.Printf("[Error] Task %d: %v\n", r.TaskID, r.Err)
		} else {
			fmt.Println(r.Msg)
		}
	}

	fmt.Println("All work completed.")
}

