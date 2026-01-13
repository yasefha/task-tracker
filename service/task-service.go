package service

import (
	"strings"
	"task-tracker/domain"
	"task-tracker/repository/file_repo"
	"time"
)

func AddTask(repo file_repo.FileRepo, description string) (task domain.Task, err error) {
	if strings.TrimSpace(description) == "" {
		return task, domain.EmptyDescriptionError{}
	}

	task = domain.Task{
		ID:          0, //nanti id yg bener di Save()
		Description: description,
		Status:      domain.StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	task, err = repo.SaveTask(task)

	if err != nil {
		return domain.Task{}, domain.UnableToSaveError{Cause: err}
	} else {
		return task, nil
	}
}
