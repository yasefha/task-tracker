package repository

import "task-tracker/domain"

type TaskRepository interface {
	SaveTask(task domain.Task) (domain.Task, error)
	ListTask(status *domain.TaskStatus) ([]domain.Task, error)
}
