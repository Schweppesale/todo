package logger

import (
	"github.com/schweppesale/todo"
	"log"
)

type TaskRepository struct {
	tasks todo.TaskRepository
}

func NewTaskRepository(tasks todo.TaskRepository) todo.TaskRepository {
	return TaskRepository{
		tasks,
	}
}

func (r TaskRepository) FindAll() ([]todo.Task, error) {
	log.Print("TaskRepository.FindAll:")
	result, err := r.tasks.FindAll()
	log.Print("TaskRepository.FindAll:", result)
	return result, err
}

func (r TaskRepository) GetByUniqueID(uniqueID string) (todo.Task, error) {
	log.Print("TaskRepository.GetByUniqueID:", uniqueID)
	result, err := r.tasks.GetByUniqueID(uniqueID)
	log.Print("TaskRepository.GetByUniqueID:", result, err)
	return result, err
}

func (r TaskRepository) Update(task todo.Task) (todo.Task, error) {
	log.Print("TaskRepository.Update:", task)
	result, err := r.tasks.Update(task)
	log.Print("TaskRepository.Update:", result, err)
	return result, err
}

func (r TaskRepository) Create(task todo.Task) (todo.Task, error) {
	log.Print("TaskRepository.Create:", task)
	result, err := r.tasks.Create(task)
	log.Print("TaskRepository.Create:", result, err)
	return result, err
}

func (r TaskRepository) Remove(uniqueID string) error {
	log.Print("TaskRepository.Remove:", uniqueID)
	return r.tasks.Remove(uniqueID)
}
