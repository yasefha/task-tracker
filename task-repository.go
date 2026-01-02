package main

type TaskRepository interface {
	Save(task Task) error
}
