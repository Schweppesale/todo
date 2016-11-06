package memory

import (
	"github.com/schweppesale/todo"
	"github.com/schweppesale/todo/infrastructure/satori"
	"testing"
)

func NewTestTaskRepository() todo.TaskRepository {
	return NewTaskRepository(satori.NewUUIDGenerator())
}

func TestTaskRepository_CreateTask(t *testing.T) {
	tasks := NewTestTaskRepository()
	newTask := todo.NewTask("title", "description")
	storedTask, err := tasks.CreateTask(newTask)
	if err != nil {
		t.Error(err)
	}

	fetchedTask, err := tasks.GetTaskByUniqueID(storedTask.UniqueID())
	if err != nil {
		t.Error(err)
	}

	if fetchedTask.UniqueID() != storedTask.UniqueID() {
		t.Error("Failed to fetch task")
	}
}
