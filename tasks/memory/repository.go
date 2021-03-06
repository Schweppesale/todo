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
	uuidService UUIDGenerator
}

func NewTaskRepository(UuidService UUIDGenerator) todo.TaskRepository {
	return TaskRepository{UuidService}
}

func (r TaskRepository) FindAll() ([]todo.Task, error) {
	container.Lock()
	result := make([]todo.Task, 0, len(container.keys))
	for _, value := range container.keys {
		result = append(result, container.tasks[value])
	}
	container.Unlock()
	return result, nil
}

func (r TaskRepository) GetByUniqueID(uniqueID string) (todo.Task, error) {
	if val, ok := container.tasks[uniqueID]; ok {
		return val, nil
	} else {
		return todo.Task{}, errors.New("Task does not exist!")
	}
}

func (r TaskRepository) Create(task todo.Task) (todo.Task, error) {
	task.SetUniqueID(r.uuidService.Generate())
	container.Lock()
	container.tasks[task.UniqueID()] = task
	container.keys = append(container.keys, task.UniqueID())
	container.Unlock()
	return task, nil
}

func (r TaskRepository) Update(task todo.Task) (todo.Task, error) {
	container.Lock()
	container.tasks[task.UniqueID()] = task
	container.Unlock()
	return task, nil
}

func (r TaskRepository) Remove(uniqueID string) error {
	container.Lock()
	delete(container.tasks, uniqueID)
	for key, value := range container.keys { //@todo optimize
		if value == uniqueID {
			container.keys = append(container.keys[:key], container.keys[key+1:]...)
		}
	}
	container.Unlock()
	return nil
}

type UUIDGenerator interface {
	Generate() string
}
