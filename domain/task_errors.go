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
