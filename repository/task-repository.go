package repository

import "task-tracker/domain"

type TaskRepository interface {
	Save(task domain.Task) error
}
