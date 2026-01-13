package cli

import (
	"errors"
	"strings"
	"task-tracker/repository/file_repo"
	"task-tracker/service"
)

func Run(args []string, repo file_repo.FileRepo) error {
	if len(args) < 2 {
		return errors.New("error")
	}

	switch args[1] {
	case "add":
		return handledAdd(args[2:], repo)
	default:
		return errors.New("Command is invalid. See \"task help\" for more information.")
	}
}

func handledAdd(args []string, repo file_repo.FileRepo) error {
	if len(args) == 0 {
		return errors.New("description is required")
	}

	description := strings.Join(args, " ")

	task, err := service.AddTask(repo, description)
	if err != nil {
		return err
	}

	PrintAddSuccess(task)
	return nil
}
