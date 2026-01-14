package repository

import "task-tracker/domain"

type TaskRepository interface {
	SaveTask(task domain.Task) (domain.Task, error)
	ListTask(status domain.TaskStatus) ([]domain.Task, error)
	UpdateTaskStatus(status domain.TaskStatus, ID int) (domain.Task, error)
	UpdateTaskDescription(description string, ID int) (domain.Task, error)
	DeleteTask(ID int) (domain.Task, error)
	DeleteAllTask() error
}
