package file_repo

import (
	"encoding/json"
	"os"
	"task-tracker/domain"
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
