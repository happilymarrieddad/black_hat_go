package main

import "fmt"

func worker(workerID int, tasks chan Task) {
	fmt.Printf("Starting worker %d\n", workerID)
	for task := range tasks {
		fmt.Printf("Worker %d Starting Type: %s ID: %s \n", workerID, task.Type(), task.ID())
		err := task.Complete()
		if err != nil {
			fmt.Printf("Worker %d was unable to complete task %s because of error '%s'\n", workerID, task.ID(), err.Error())
		}
	}
}
