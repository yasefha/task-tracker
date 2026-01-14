package cli

import (
	"errors"
	"fmt"
	"strconv"
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
	case "update-status":
		return handleUpdateStatus(args[2:], a.repo)
	case "update-desc":
		return handleUpdateDescription(args[2:], a.repo)
	case "delete":
		return handleDelete(args[2:], a.repo)
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
	var status domain.TaskStatus

	if len(args) > 1 {
		return domain.InvalidCommandError{}
	}

	if len(args) == 1 {
		s, err := parseStatus(args[0])
		if err != nil {
			return err
		}
		status = s
	}

	var tasks []domain.Task
	tasks, err := service.ListTask(repo, status)
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		return fmt.Errorf("No tasks found with status \"%s\".", status)
	}

	PrintTasksList(tasks)

	return nil
}

func handleUpdateStatus(args []string, repo repository.TaskRepository) error {
	if len(args) > 2 {
		return domain.InvalidCommandError{}
	}

	ID, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	s, err := parseStatus(args[1])
	if err != nil {
		return err
	}

	task, err := service.UpdateStatus(repo, ID, s)

	if errors.Is(err, domain.TaskNotFoundError{}) {
		fmt.Println("[FAIL] No task found with ID", ID)
	} else {
		PrintUpdateSuccess(task)
	}

	return nil
}

func handleUpdateDescription(args []string, repo repository.TaskRepository) error {
	if len(args) < 1 {
		return domain.InvalidCommandError{}
	}

	if len(args) < 2 {
		return domain.EmptyDescriptionError{}
	}

	ID, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	desc := strings.Join(args[1:], " ")

	task, err := service.UpdateDescription(repo, ID, desc)
	if errors.Is(err, domain.TaskNotFoundError{}) {
		fmt.Println("[FAIL] No task found with ID", ID)
	} else {
		PrintUpdateSuccess(task)
	}

	return nil
}

func handleDelete(args []string, repo repository.TaskRepository) error {
	if len(args) < 1 || len(args) > 1 {
		return domain.InvalidCommandError{}
	}

	ID, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	task, err := service.DeleteTaskByID(repo, ID)
	if errors.Is(err, domain.TaskNotFoundError{}) {
		fmt.Println("[FAIL] No task found with ID", ID)
	} else {
		PrintDeleteSuccess(task)
	}

	return nil
}
