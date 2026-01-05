package main

type EmptyDescriptionError struct{}

func (e EmptyDescriptionError) Error() string {
	return "task description cannot be empty"
}

type UnableToSaveError struct{}

func (e UnableToSaveError) Error() string {
	return "unable to save task"
}
