package main

import (
	"testing"
	"time"
)

func TestSingleWorkerMaxTasksAndTiming(t *testing.T) {
	maxTasks := 5
	numTasks := 8
	maxDuration := 250 * time.Millisecond

	tasks := make(chan Task, numTasks)
	results := make(chan Result, numTasks)

	// FIXED: pass maxDuration to match new worker signature
	go worker(tasks, results, maxTasks, maxDuration)

	// Send tasks
	for i := 1; i <= numTasks; i++ {
		tasks <- Task{ID: i}
	}
	close(tasks)

	successes := 0
	maxTasksErrors := 0
	timingErrors := 0

	for r := range results {
		if r.Err != nil {
			if r.Err.Error() == "max_tasks limit exceeded" {
				maxTasksErrors++
			} else {
				timingErrors++
				successes++ // count timing-error tasks as processed
			}
		} else {
			successes++
		}

	}

	if successes != maxTasks {
		t.Errorf("expected %d successes, got %d", maxTasks, successes)
	}
	if maxTasksErrors != numTasks-maxTasks {
		t.Errorf("expected %d max_tasks errors, got %d", numTasks-maxTasks, maxTasksErrors)
	}
	if timingErrors == 0 {
		t.Errorf("expected at least one timing error but got none")
	}
}

