package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID   int
	Data []byte
}

type TaskChannel chan Task

func worker(taskChannel TaskChannel, wg *sync.WaitGroup) {
	fmt.Println("Worker started")
	for task := range taskChannel {
		fmt.Printf("Starting task %d\n", task.ID)

		time.Sleep(time.Second * 2)
		wg.Done()
	}
}

func main() {
	wg := &sync.WaitGroup{}
	numWorkers := 5
	taskChannel := make(TaskChannel)

	data := []Task{}

	for i := 0; i < 20; i++ {
		data = append(data, Task{ID: i + 1})
	}

	wg.Add(len(data))

	for i := 0; i < numWorkers; i++ {
		go worker(taskChannel, wg)
	}

	// This is just to delay so the workers are started and ready
	time.Sleep(time.Second)

	for _, d := range data {
		taskChannel <- d
	}

	wg.Wait()
	close(taskChannel)

	fmt.Println("Finished!!")
}
