package cli

import (
	"fmt"
	"task-tracker/domain"
)

func printSeparator() {
	fmt.Println("────────────────────────────────────────")
}

type TaskOutput struct {
	ID          int
	Description string
	Status      string
	CreatedAt   string
	UpdatedAt   string
}

func TaskView(task domain.Task) TaskOutput {
	return TaskOutput{
		ID:          task.ID,
		Description: task.Description,
		Status:      string(task.Status),
		CreatedAt:   task.CreatedAt.Format("2006-01-02 15:04"),
		UpdatedAt:   task.UpdatedAt.Format("2006-01-02 15:04"),
	}
}

func (t TaskOutput) String() string {
	return fmt.Sprintf(
		"  Description : %s\n  Status      : %s\n  Created At  : %s\n  Updated At  : %s",
		t.Description,
		t.Status,
		t.CreatedAt,
		t.UpdatedAt,
	)
}

func PrintAddSuccess(task domain.Task) {
	view := TaskView(task)

	fmt.Println(green("[SUCCESS]"), "Task added , (", cyan("ID"), ":", fmt.Sprint(view.ID))
	printSeparator()
	fmt.Println(view.String())
	printSeparator()
}

func PrintTasksList(tasks []domain.Task) {
	fmt.Println("════════════════════════════════════════")
	fmt.Println("				 TASKS LIST 				")
	fmt.Println("════════════════════════════════════════")

	for _, task := range tasks {
		view := TaskView(task)
		fmt.Println(cyan("ID:"), fmt.Sprint(view.ID))
		printSeparator()
		fmt.Println(view.String())
		printSeparator()
	}
}

func PrintUpdateSuccess(task domain.Task) {
	view := TaskView(task)

	fmt.Println(green("[SUCCESS]"), "Task updated , (", cyan("ID"), ":", fmt.Sprint(view.ID))
	printSeparator()
	fmt.Println(view.String())
	printSeparator()
}

func PrintDeleteSuccess(task domain.Task) {
	view := TaskView(task)

	fmt.Println(green("[SUCCESS]"), "Task deleted , (", cyan("ID"), ":", fmt.Sprint(view.ID))
	printSeparator()
	fmt.Println(view.String())
	printSeparator()
}

func PrintHelp() {
	fmt.Println("════════════════════════════════════════")
	fmt.Println(bold(green("	   TASKS TRACKER CLI    ")))
	fmt.Println("════════════════════════════════════════")
	fmt.Println()

	fmt.Println(bold("USAGE"))
	printSeparator()
	fmt.Println(`task <command> [arguments]`)
	fmt.Println()

	fmt.Println()
	fmt.Println(bold("COMMANDS"))
	printSeparator()
	fmt.Println("Add a new task:", bold(cyan("add")), magenta("<description>\n"))
	fmt.Println("List tasks: ", bold(cyan("list")), magenta("[status]"))
	fmt.Println("If no status is provided, all tasks will be shown.")
	fmt.Println("Status options: todo | in-progress | done\n")
	fmt.Println("Update task status: ", bold(cyan("update-status")), magenta("[ID]"), magenta("[status]\n"))
	fmt.Println("Update task description: ", bold(cyan("update-description")), magenta("[ID]"), magenta("<description>\n"))
	fmt.Println("Delete a task by ID:", bold(cyan("delete")), magenta("[ID]\n"))
	fmt.Println("Delete all tasks permanently:", bold(cyan("reset")), gray("[--confirm]"))
	fmt.Println()

	fmt.Println()
	fmt.Println(bold("EXAMPLES"))
	printSeparator()
	fmt.Println(bold(green("task")), cyan("add"), magenta("\"Buy milk\""))
	fmt.Println(bold(green("task")), cyan("list"), magenta("todo"))
	fmt.Println(bold(green("task")), cyan("updaye-desc"), magenta("2"), magenta("\"Fix bug\""))
	fmt.Println(bold(green("task")), cyan("delete"), magenta("3"))
	fmt.Println(bold(green("task")), cyan("reset"), gray("--confirm"))
	fmt.Println()
	fmt.Println("════════════════════════════════════════")
}

func printWelcome() {
	fmt.Println("════════════════════════════════════════")
	fmt.Println(bold(blue("             TASK TRACKER             ")))
	fmt.Println("            by Yasef Hatam            ")
	fmt.Println("════════════════════════════════════════")
	fmt.Println()

	fmt.Println(`A simple command-line tool to help you
manage your daily tasks efficiently.

You can add, list, update, and delete tasks
directly from your terminal.`)
	printSeparator()

	fmt.Println()
	fmt.Println("Type", cyan("\"task help\""), "to see available commands.")
	fmt.Println("════════════════════════════════════════")
}
