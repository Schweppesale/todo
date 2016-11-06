package todo

import (
	"testing"
)

func NewTestTask() Task {
	return NewTask("title", "description")
}

func TestNewTask(t *testing.T) {

	task := NewTestTask()
	title := "new title"
	task.SetTitle(title)
	if task.Title() != title {
		t.Error("Title has not been changed!")
	}

	description := "description2"
	task.SetDescription(description)
	if task.Description() != description {
		t.Error("Description has not been changed!")
	}

	id := task.UniqueID()
	task.SetUniqueID("uniqueid")

	if id == task.UniqueID() {
		t.Error("Unique ID has not been updated!")
	}
}
