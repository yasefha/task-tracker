package cli

import (
	"errors"
	"fmt"
	"strconv"
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
		printWelcome()
		return nil
	}

	switch args[1] {
	case "add":
		return handleAdd(args[2:], a.repo)
	case "list":
		return handleList(args[2:], a.repo)
	case "update-status":
		return handleUpdateStatus(args[2:], a.repo)
	case "update-desc":
		return handleUpdateDescription(args[2:], a.repo)
	case "delete":
		return handleDelete(args[2:], a.repo)
	case "reset":
		return handleReset(args[2:], a.repo)
	case "help":
		return handleHelp()
	default:
		return domain.InvalidCommandError{}
	}
}

func handleAdd(args []string, repo repository.TaskRepository) error {
	if len(args) < 1 {
		fmt.Println("[FAIL] Task not saved.")
		return domain.EmptyDescriptionError{}
	}

	if len(args) != 1 {
		fmt.Println("[FAIL] Task not saved.")
		return errors.New("Description must be provided as a single argument. \"<description>\".")
	}

	description := args[0]

	task, err := service.AddTask(repo, description)
	if err != nil {
		fmt.Println("[FAIL] Task not saved.")
		return domain.UnableToSaveError{Cause: err}
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
		return "", errors.New("Invalid status")
	}
}

func handleList(args []string, repo repository.TaskRepository) error {
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
		fmt.Println("[FAIL] Task not updated.")
		return domain.InvalidCommandError{}
	}

	ID, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	s, err := parseStatus(args[1])
	if err != nil {
		fmt.Println("[FAIL] Task not updated.")
		return err
	}

	task, err := service.UpdateStatus(repo, ID, s)

	if errors.Is(err, domain.TaskNotFoundError{}) {
		fmt.Println("[FAIL] No task found with ID", ID)
		return err
	} else {
		PrintUpdateSuccess(task)
	}

	return nil
}

func handleUpdateDescription(args []string, repo repository.TaskRepository) error {
	if len(args) < 1 {
		fmt.Println("[FAIL] Task not updated.")
		return domain.EmptyDescriptionError{}
	}

	if len(args) != 1 {
		fmt.Println("[FAIL] Task not updated.")
		return errors.New("Description must be provided as a single argument. \"<description>\".")
	}

	ID, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	desc := args[0]

	task, err := service.UpdateDescription(repo, ID, desc)
	if errors.Is(err, domain.TaskNotFoundError{}) {
		fmt.Println("[FAIL] No task found with ID", ID)
		return err
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
		return err
	} else {
		PrintDeleteSuccess(task)
	}

	return nil
}

func handleReset(args []string, repo repository.TaskRepository) error {
	if len(args) == 0 {
		fmt.Println("[WARNING] This will delete ALL tasks permanently.")
		fmt.Println("To confirm, run: task reset --confirm")
		return nil
	}

	if len(args) != 1 || args[0] != "--confirm" {
		return errors.New("Invalid option. use --confirm to proceed.")
	}

	err := service.DeleteAllTask(repo)
	if err != nil {
		return err
	}

	fmt.Println("[SUCCESS] All tasks deleted.")
	return nil
}

func handleHelp() error {
	PrintHelp()
	return nil
}
