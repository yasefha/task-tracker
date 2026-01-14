package file_repo

import (
	"encoding/json"
	"os"
	"task-tracker/domain"
	"time"
)

type FileRepo struct {
	Dir string
}

type taskFileState struct {
	LastID int
	Tasks  []domain.Task
}

func loadState(repo *FileRepo) (taskFileState, error) {
	var state taskFileState

	data, err := os.ReadFile(repo.Dir)
	if err != nil {
		return state, err
	}

	err = json.Unmarshal(data, &state)
	if err != nil {
		return state, err
	}

	return state, nil
}

func saveState(state taskFileState, repo *FileRepo) error {
	bytes, err := json.Marshal(state)
	if err != nil {
		return err
	}

	err = os.WriteFile(repo.Dir, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (repo *FileRepo) SaveTask(task domain.Task) (domain.Task, error) {
	state, err := loadState(repo)
	if err != nil {
		return domain.Task{}, err
	}

	state.LastID++
	task.ID = state.LastID
	state.Tasks = append(state.Tasks, task)

	err = saveState(state, repo)
	if err != nil {
		return domain.Task{}, err
	}

	return task, nil
}

func (repo *FileRepo) ListTask(status domain.TaskStatus) (tasks []domain.Task, err error) {
	state, err := loadState(repo)
	if err != nil {
		return nil, err
	}

	if len(state.Tasks) == 0 {
		return nil, domain.EmptyTaskError{}
	}

	if status == "" {
		return state.Tasks, nil
	}

	var filtered []domain.Task
	for _, task := range state.Tasks {
		if task.Status == status {
			filtered = append(filtered, task)
		}
	}

	return filtered, nil
}

func (repo *FileRepo) UpdateTaskStatus(status domain.TaskStatus, ID int) (domain.Task, error) {
	state, err := loadState(repo)
	if err != nil {
		return domain.Task{}, err
	}

	for i := range state.Tasks {
		if state.Tasks[i].ID == ID {
			state.Tasks[i].Status = status
			state.Tasks[i].UpdatedAt = time.Now()

			if err := saveState(state, repo); err != nil {
				return domain.Task{}, err
			}

			return state.Tasks[i], nil
		}
	}

	return domain.Task{}, domain.TaskNotFoundError{}
}

func (repo *FileRepo) UpdateTaskDescription(description string, ID int) (domain.Task, error) {
	state, err := loadState(repo)
	if err != nil {
		return domain.Task{}, err
	}

	for i := range state.Tasks {
		if state.Tasks[i].ID == ID {
			state.Tasks[i].Description = description
			state.Tasks[i].UpdatedAt = time.Now()

			if err := saveState(state, repo); err != nil {
				return domain.Task{}, err
			}

			return state.Tasks[i], nil
		}
	}

	return domain.Task{}, domain.TaskNotFoundError{}
}

func (repo *FileRepo) DeleteTask(ID int) (domain.Task, error) {
	state, err := loadState(repo)
	if err != nil {
		return domain.Task{}, err
	}

	for i := range state.Tasks {
		if state.Tasks[i].ID == ID {
			deletedTask := state.Tasks[i]

			state.Tasks = append(state.Tasks[:i], state.Tasks[i+1:]...)
			if err := saveState(state, repo); err != nil {
				return domain.Task{}, err
			}

			return deletedTask, nil
		}
	}
	return domain.Task{}, domain.TaskNotFoundError{}
}

func (repo *FileRepo) DeleteAllTask() error {
	state := taskFileState{
		LastID: 0,
		Tasks:  nil,
	}

	return saveState(state, repo)
}
