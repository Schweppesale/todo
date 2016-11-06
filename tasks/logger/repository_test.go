package logger

import (
	"github.com/schweppesale/todo"
	"github.com/schweppesale/todo/tasks/memory"
	"github.com/schweppesale/todo/tasks/memory/satori"
	"testing"
)

func NewTestTaskRepository() todo.TaskRepository {
	return NewTaskRepository(memory.NewTaskRepository(satori.NewUUIDGenerator()))
}

func TestNewTaskRepository(t *testing.T) {
	tasks := NewTestTaskRepository()
	newTask := todo.NewTask("title", "description")
	storedTask, err := tasks.Create(newTask)
	if err != nil {
		t.Error(err)
	}

	_, err = tasks.Update(storedTask)
	if err != nil {
		t.Error(err)
	}

	fetchedTask, err := tasks.GetByUniqueID(storedTask.UniqueID())
	if err != nil {
		t.Error(err)
	}

	err = tasks.Remove(fetchedTask.UniqueID())

}
