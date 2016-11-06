package api

import (
	"github.com/schweppesale/todo/tasks/memory/satori"
	"github.com/schweppesale/todo/tasks/logger"
	"github.com/schweppesale/todo/tasks/memory"
	"testing"
)

func NewTestTaskService() TaskService {
	uuidGenerator := satori.NewUUIDGenerator()
	tasks := logger.NewTaskRepository(memory.NewTaskRepository(uuidGenerator))
	return NewTaskService(tasks)
}

func TestTaskService_UpdateTask(t *testing.T) {
	service := NewTestTaskService()
	task, err := service.CreateTask("test", "description")
	if(err != nil) {
		t.Error(err)
	}

	uniqueID := task.UniqueID
	task, err = service.UpdateTask(uniqueID, "test2", "description2")
	if(err != nil) {
		t.Error(err)
	}
}