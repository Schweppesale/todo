package memory

import (
	"errors"
	"github.com/schweppesale/todo"
	"sync"
)

var container = struct {
	sync.RWMutex
	tasks map[string]todo.Task
	keys  []string
}{tasks: make(map[string]todo.Task)}

type TaskRepository struct {
	uuidService todo.UUIDGenerator
}

func NewTaskRepository(UuidService todo.UUIDGenerator) todo.TaskRepository {
	return TaskRepository{UuidService}
}

func (r TaskRepository) FindTasks() ([]todo.Task, error) {
	result := make([]todo.Task, 0, len(container.keys))
	for _, value := range container.keys {
		result = append(result, container.tasks[value])
	}
	return result, nil
}

func (r TaskRepository) GetTaskByUniqueId(uniqueId string) (todo.Task, error) {
	if val, ok := container.tasks[uniqueId]; ok {
		return val, nil
	} else {
		return todo.Task{}, errors.New("Task does not exist!")
	}
}

func (r TaskRepository) CreateTask(task todo.Task) (todo.Task, error) {
	task.SetUniqueID(r.uuidService)
	container.Lock()
	container.tasks[task.UniqueId()] = task
	container.keys = append(container.keys, task.UniqueId())
	container.Unlock()
	return task, nil
}

func (r TaskRepository) UpdateTask(task todo.Task) (todo.Task, error) {
	container.Lock()
	container.tasks[task.UniqueId()] = task
	container.Unlock()
	return task, nil
}

func (r TaskRepository) RemoveTask(uniqueId string) error {
	container.Lock()
	delete(container.tasks, uniqueId)
	for key, value := range container.keys { //@todo optimize
		if value == uniqueId {
			container.keys = append(container.keys[:key], container.keys[key+1:]...)
		}
	}
	container.Unlock()
	return nil
}
