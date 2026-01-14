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

	fmt.Printf("[SUCCESS] Task added (ID: %d) \n", view.ID)
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
		fmt.Println("ID:", view.ID)
		printSeparator()
		fmt.Println(view.String())
		printSeparator()
	}
}

func PrintUpdateSuccess(task domain.Task) {
	view := TaskView(task)

	fmt.Printf("[SUCCESS] Task updated (ID: %d) \n", view.ID)
	printSeparator()
	fmt.Println(view.String())
	printSeparator()
}

func PrintDeleteSuccess(task domain.Task) {
	view := TaskView(task)

	fmt.Printf("[SUCCESS] Task deleted (ID: %d) \n", view.ID)
	printSeparator()
	fmt.Println(view.String())
	printSeparator()
}

func PrintHelp() {
	fmt.Println(`════════════════════════════════════════
            TASK TRACKER CLI
════════════════════════════════════════

USAGE
────────────────────────────────────────
  task <command> [arguments]


COMMANDS
────────────────────────────────────────
  add <description>
      Add a new task

  list [status]
      List tasks.
      If no status is provided, all tasks will be shown.
      Status options: todo | in-progress | done

  update-status <id> <status>
      Update task status

  update-desc <id> <description>
      Update task description

  delete <id>
      Delete a task by ID

  reset [--confirm]
      Delete all tasks permanently


EXAMPLES
────────────────────────────────────────
  task add "Buy milk"
  task list
  task list todo
  task update-status 1 done
  task update-desc 2 "Fix bug"
  task delete 3
  task reset --confirm

════════════════════════════════════════
`)
}

func printWelcome() {
	fmt.Println(`════════════════════════════════════════
            TASK TRACKER
             by Yasef Hatam
════════════════════════════════════════

A simple command-line tool to help you
manage your daily tasks efficiently.

You can add, list, update, and delete tasks
directly from your terminal.

────────────────────────────────────────
Type "task help" to see available commands
════════════════════════════════════════
`)
}
