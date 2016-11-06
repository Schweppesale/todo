package todo

import (
	"github.com/schweppesale/todo/tasks/memory/satori"
	"testing"
)

func NewTestTask() Task {
	return NewTask("title", "description")
}

func TestTask_SetTitle(t *testing.T) {
	task := NewTestTask()
	title := "new title"
	task.SetTitle(title)
	if task.Title() != title {
		t.Error("Title has not been changed!")
	}
}

func TestTask_SetDescription(t *testing.T) {
	task := NewTestTask()
	description := "description2"
	task.SetDescription(description)
	if task.Description() != description {
		t.Error("Description has not been changed!")
	}
}

func TestTask_SetUniqueID(t *testing.T) {
	task := NewTestTask()
	id := task.UniqueID()
	generator := satori.NewUUIDGenerator()
	task.SetUniqueID(generator.Generate())

	if id == task.UniqueID() {
		t.Error("Unique ID has not been updated!")
	}
}
