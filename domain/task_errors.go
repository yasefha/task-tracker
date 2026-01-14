package domain

type EmptyDescriptionError struct{}

func (e EmptyDescriptionError) Error() string {
	return "task description cannot be empty"
}

type UnableToSaveError struct {
	Cause error
}

func (e UnableToSaveError) Error() string {
	return "unable to save task"
}

type EmptyTaskError struct{}

func (e EmptyTaskError) Error() string {
	return "No tasks found. Add a task using \"task add \"<description>\"\"."
}

type InvalidCommandError struct{}

func (e InvalidCommandError) Error() string {
	return "Command is invalid. See \"task help\" for more information."
}

type TaskNotFoundError struct{}

func (e TaskNotFoundError) Error() string {
	return e.Error()
}
