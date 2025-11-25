package main

import (
	"testing"
	"time"
)

func TestWorkerWithCustomParams(t *testing.T) {
	// You can change these values to test different scenarios
	maxTasks := 5
	numTasks := 8
	maxDuration := 100 * time.Millisecond

	tasks := make(chan Task, numTasks)
	results := make(chan Result, numTasks)

	// Start worker (same function as main.go)
	go worker(tasks, results, maxTasks, maxDuration)

	// Send tasks
	for i := 1; i <= numTasks; i++ {
		tasks <- Task{ID: i}
	}
	close(tasks)

	// Collect results and report errors
	hadErrors := false
	for r := range results {
		if r.Err != nil {
			t.Logf("[Error] Task %d: %v", r.TaskID, r.Err)
			hadErrors = true
		} else {
			t.Logf("%s", r.Msg)
		}
	}

	if hadErrors {
		t.Errorf("Worker produced errors. See logs above.")
	}
}
