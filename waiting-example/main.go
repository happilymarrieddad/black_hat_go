package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type ErrHandler struct {
	Errs []error
}

type Task struct {
	ID   int
	Data []byte
}

type TaskChannel chan Task

func worker(taskChannel TaskChannel, wg *sync.WaitGroup, errHandler *ErrHandler) {
	fmt.Println("Worker started")
	for task := range taskChannel {
		fmt.Printf("Starting task %d\n", task.ID)

		time.Sleep(time.Second * 2)
		if task.ID == 2 {
			fmt.Println("Added err")
			errHandler.Errs = append(errHandler.Errs, errors.New("some error"))
		}
		if task.ID == 10 {
			fmt.Println("Added err")
			errHandler.Errs = append(errHandler.Errs, errors.New("some error 2"))
		}
		wg.Done()
	}
}

func main() {
	errHandler := new(ErrHandler)
	wg := &sync.WaitGroup{}
	numWorkers := 5
	taskChannel := make(TaskChannel)

	data := []Task{}

	for i := 0; i < 20; i++ {
		data = append(data, Task{ID: i + 1})
	}

	wg.Add(len(data))

	for i := 0; i < numWorkers; i++ {
		go worker(taskChannel, wg, errHandler)
	}

	// This is just to delay so the workers are started and ready
	time.Sleep(time.Second)

	for _, d := range data {
		taskChannel <- d
	}

	wg.Wait()
	close(taskChannel)

	var err error
	for _, e := range errHandler.Errs {
		fmt.Println("handling err")
		err = errors.Join(err, e)
	}

	fmt.Println("Finished!!")
	fmt.Println("With errs: ")
	fmt.Println(err.Error())
}
