package cli

import (
	"fmt"
	"task-tracker/domain"
)

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
		"Description: %s\nStatus: %s\nCreated at: %s\nUpdated at: %s",
		t.Description,
		t.Status,
		t.CreatedAt,
		t.UpdatedAt,
	)
}

func PrintAddSuccess(task domain.Task) {
	view := TaskView(task)

	fmt.Println("[SUCC] Task added (ID:", view.ID, ")")
	fmt.Println(view.String())
}

func PrintTasksList(tasks []domain.Task) {
	fmt.Println("TASKS LIST")

	for _, task := range tasks {
		view := TaskView(task)
		fmt.Println("ID:", view.ID)
		fmt.Println(view.String())
	}
}

func PrintUpdateSuccess(task domain.Task) {
	view := TaskView(task)

	fmt.Println("[SUCC] Task updated (ID:", view.ID, ")")
	fmt.Println(view.String())
}

func PrintDeleteSuccess(task domain.Task) {
	view := TaskView(task)

	fmt.Println("[SUCC] Task deleted (ID: ", view.ID, ")")
	fmt.Println(view.String())
}
