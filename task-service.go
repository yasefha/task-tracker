package main

import (
	"strings"
	"time"
)

func AddTask(repo TaskRepository, description string) (task Task, err error) {
	if strings.TrimSpace(description) == "" {
		return task, EmptyDescriptionError{}
	}

	task = Task{
		ID:          0, //nanti id yg bener di Save()
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = repo.Save(task)

	if err != nil {
		return Task{}, UnableToSaveError{}
	} else {
		return task, nil
	}
}
