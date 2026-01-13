package service

import (
	"strings"
	"task-tracker/domain"
	"task-tracker/repository"
	"time"
)

func AddTask(repo repository.TaskRepository, description string) (task domain.Task, err error) {
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

func ListTask(repo repository.TaskRepository, status *domain.TaskStatus) ([]domain.Task, error) {
	return repo.ListTask(status)
}
