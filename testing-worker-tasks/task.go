package main

// Task - a task to complete
type Task interface {
	ID() string
	Type() string
	Complete() error
}
