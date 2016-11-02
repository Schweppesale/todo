package memory

import (
	"errors"
	"github.com/schweppesale/todo/domain/entities"
	"github.com/schweppesale/todo/domain/repositories"
	"github.com/schweppesale/todo/domain/services"
	"sync"
)

var container = struct {
	sync.RWMutex
	tasks map[string]entities.Task
	keys  []string
}{tasks: make(map[string]entities.Task)}

type TaskRepository struct {
	uuidService services.UuidService
}

func NewTaskRepository(UuidService services.UuidService) repositories.TaskRepository {
	return TaskRepository{UuidService}
}

func (r TaskRepository) FindAll() ([]entities.Task, error) {
	result := make([]entities.Task, 0, len(container.keys))
	for _, value := range container.keys {
		result = append(result, container.tasks[value])
	}
	return result, nil
}

func (r TaskRepository) GetTaskByUniqueId(uniqueId string) (entities.Task, error) {
	if val, ok := container.tasks[uniqueId]; ok {
		return val, nil
	} else {
		return entities.Task{}, errors.New("Task does not exist!")
	}
}

func (r TaskRepository) CreateTask(task entities.Task) (entities.Task, error) {
	task.SetUniqueID(r.uuidService.Generate())
	container.Lock()
	container.tasks[task.UniqueId()] = task
	container.keys = append(container.keys, task.UniqueId())
	container.Unlock()
	return task, nil
}

func (r TaskRepository) UpdateTask(task entities.Task) (entities.Task, error) {
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
