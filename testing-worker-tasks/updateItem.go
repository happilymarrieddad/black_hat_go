package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type updateItemInDB struct {
	id string
}

// NewUpdateItemInDB Create a new item to db
func NewUpdateItemInDB() updateItemInDB {
	return updateItemInDB{id: uuid.New().String()}
}

func (a updateItemInDB) ID() string {
	return a.id
}

func (a updateItemInDB) Type() string {
	return "updateItemInDB"
}

func (a updateItemInDB) Complete() error {
	fmt.Printf("Starting task %s Complete\n", a.id)
	time.Sleep(time.Second * 2)
	fmt.Printf("Finished task %s Complete\n", a.id)

	return nil
}
