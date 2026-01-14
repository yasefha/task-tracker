package main

import (
	"fmt"
	"os"

	"github.com/yasefha/task-tracker/cli"
	"github.com/yasefha/task-tracker/repository/file_repo"
)

func main() {
	repo := &file_repo.FileRepo{}
	repo.Dir = "data/tasks.json"

	app := cli.NewApp(repo)

	args := os.Args
	err := app.Run(args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
