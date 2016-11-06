package api

import (
	"github.com/schweppesale/todo/tasks/memory"
	"github.com/schweppesale/todo/tasks/memory/satori"
	"testing"
)

func NewTestTaskService() TaskService {
	uuidGenerator := satori.NewUUIDGenerator()
	tasks := memory.NewTaskRepository(uuidGenerator)
	return NewTaskService(tasks)
}

func TestNewTaskService(t *testing.T) {
	title := "title"
	description := "description"
	service := NewTestTaskService()
	task, err := service.Create(title, description)
	if err != nil {
		t.Error(err)
	}

	if task.Title != title || task.Description != description {
		t.Error("Task not set!")
	}

	uniqueID := task.UniqueID
	updatedTitle := "test2"
	updatedDescription := "description2"
	updatedTask, err := service.Update(uniqueID, updatedTitle, updatedDescription)
	if err != nil {
		t.Error(err)
	}

	if updatedTask.Title != updatedTitle || updatedTask.Description != updatedDescription {
		t.Error("Task not updated!")
	}

	fetchedTask, err := service.GetByUniqueID(task.UniqueID)
	if err != nil {
		t.Error(err)
	}

	if fetchedTask.Title != updatedTitle {
		t.Error("Invalid task!")
	}

	err = service.Remove(uniqueID)
	if err != nil {
		t.Error(err)
	}

	deletedTask, err := service.GetByUniqueID(uniqueID)
	if err == nil {
		t.Error("Task should have been removed: ", deletedTask)
	}
}
