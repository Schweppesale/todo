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
}{tasks: make(map[string]entities.Task)}

type TaskRepository struct {
	uuidGenerator services.UniqueIdGenerator
}

func (r TaskRepository) FindAll() (map[string]entities.Task, error) {
	return container.tasks, nil
}

func (r TaskRepository) GetTaskByUniqueId(uniqueId string) (entities.Task, error) {
	if val, ok := container.tasks[uniqueId]; ok {
		return val, nil
	} else {
		return entities.Task{}, errors.New("Task does not exist!")
	}
}

func (r TaskRepository) CreateTask(task entities.Task) (entities.Task, error) {
	task.SetUniqueID(r.uuidGenerator.Generate())
	container.Lock()
	container.tasks[task.UniqueId()] = task
	container.Unlock()
	return task, nil
}

func (r TaskRepository) UpdateTask(task entities.Task) (entities.Task, error) {
	container.Lock()
	container.tasks[task.UniqueId()] = task
	container.Unlock()
	return task, nil
}

func (r TaskRepository) RemoveTask(uniqueId string) {
	container.Lock()
	delete(container.tasks, uniqueId)
	container.Unlock()
}

func NewTaskRepository(uuidGenerator services.UniqueIdGenerator) repositories.TaskRepository {
	return TaskRepository{uuidGenerator}
}
