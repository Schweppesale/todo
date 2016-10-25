package memory

import (
	"github.com/user/todo/domain/entities"
	"sync"
)

var container = struct{
	sync.RWMutex
	tasks map[string] entities.Task
}{tasks: make(map[string] entities.Task)}

type TaskRepository struct {}

func(r TaskRepository) FindAll() map[string] entities.Task  {
	return container.tasks
}

func(r TaskRepository) GetTaskByUniqueId(uniqueId string) entities.Task {
	return container.tasks[uniqueId]
}

func(r TaskRepository) SaveTask(task entities.Task) entities.Task {
	container.Lock()
	container.tasks[task.UniqueId()] = task
	container.Unlock()
	return task
}

func(r TaskRepository) RemoveTask(uniqueId string) {
	container.Lock()
	delete(container.tasks, uniqueId)
	container.Unlock()
}

