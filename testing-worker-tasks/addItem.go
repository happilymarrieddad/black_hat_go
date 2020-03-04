package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type addItemToDB struct {
	id string
}

// NewAddItemToDB Create a new item to db
func NewAddItemToDB() addItemToDB {
	return addItemToDB{id: uuid.New().String()}
}

func (a addItemToDB) ID() string {
	return a.id
}

func (a addItemToDB) Type() string {
	return "addItemToDB"
}

func (a addItemToDB) Complete() error {
	fmt.Printf("Starting task %s Complete\n", a.id)
	time.Sleep(time.Second * 2)
	fmt.Printf("Finished task %s Complete\n", a.id)

	return nil
}
