package main

import (
	"fmt"
	"time"
)

func main() {
	tasks := make(chan Task)

	for i := 0; i < 10; i++ {
		go worker(i, tasks)
	}

	// Sending in Add Item Requests
	go func() {
		fmt.Println("Waited 10 seconds and now starting to send tasks")
		time.Sleep(time.Second * 10)

		for j := 0; j < 25; j++ {
			tasks <- NewAddItemToDB()
		}
	}()

	// Sending in Update Item Requests
	go func() {
		fmt.Println("Waited 10 seconds and now starting to send tasks")
		time.Sleep(time.Second * 10)

		for j := 0; j < 25; j++ {
			tasks <- NewUpdateItemInDB()
		}
	}()

	fmt.Println("Waiting 60 seconds to close this application")
	time.Sleep(time.Second * 60)
	close(tasks)

	return
}
