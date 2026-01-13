package cli

import (
	"errors"
	"fmt"
	"strings"
	"task-tracker/domain"
	"task-tracker/repository"
	"task-tracker/service"
)

type App struct {
	repo repository.TaskRepository
}

func NewApp(repo repository.TaskRepository) *App {
	return &App{repo: repo}
}

func (a *App) Run(args []string) error {
	if len(args) < 2 {
		return errors.New("error")
	}

	switch args[1] {
	case "add":
		return handleAdd(args[2:], a.repo)
	case "list":
		return hadleList(args[2:], a.repo)
	default:
		return domain.InvalidCommandError{}
	}
}

func handleAdd(args []string, repo repository.TaskRepository) error {
	if len(args) == 0 {
		fmt.Println("[FAIL] Task not saved.")
		return errors.New("Description is required.")
	}

	description := strings.Join(args, " ")

	task, err := service.AddTask(repo, description)
	if err != nil {
		fmt.Println("[FAIL] Task not saved.")
		return err
	}

	PrintAddSuccess(task)
	return nil
}

func parseStatus(arg string) (domain.TaskStatus, error) {
	switch arg {
	case "todo":
		return domain.StatusTodo, nil
	case "in-progress":
		return domain.StatusInProgress, nil
	case "done":
		return domain.StatusDone, nil
	default:
		return "", errors.New("invalid status")
	}
}

func hadleList(args []string, repo repository.TaskRepository) error {
	var status *domain.TaskStatus

	if len(args) > 1 {
		return domain.InvalidCommandError{}
	}

	if len(args) == 1 {
		s, err := parseStatus(args[0])
		if err != nil {
			return err
		}
		status = &s
	}

	var tasks []domain.Task
	tasks, err := service.ListTask(repo, status)
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		return fmt.Errorf("No tasks found with status \"%s\".", *status)
	}

	PrintTasksList(tasks)

	return nil
}
