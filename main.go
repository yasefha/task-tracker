package main

import (
	"fmt"
	"os"
	"task-tracker/cli"
	"task-tracker/repository/file_repo"
)

func main() {
	repo := &file_repo.FileRepo{}
	repo.Dir = "data/tasks.json"

	app := cli.NewApp(repo)

	args := os.Args
	err := app.Run(args)
	if err != nil {
		fmt.Println(err.Error())
		// os.Exit(1)
	}
}
