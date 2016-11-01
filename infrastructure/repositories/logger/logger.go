package logger

import (
	"github.com/schweppesale/todo/domain/entities"
	"github.com/schweppesale/todo/domain/repositories"
	"log"
)

type TaskRepository struct {
	tasks repositories.TaskRepository
}

func (r TaskRepository) FindAll() (map[string]entities.Task, error) {
	result, err := r.tasks.FindAll()
	log.Print("TaskRepository.FindAll:", result, err)
	return result, err
}

func (r TaskRepository) GetTaskByUniqueId(uniqueId string) (entities.Task, error) {
	result, err := r.tasks.GetTaskByUniqueId(uniqueId)
	log.Print("TaskRepository.GetTaskByUniqueId:", uniqueId, result, err)
	return result, err
}

func (r TaskRepository) UpdateTask(task entities.Task) (entities.Task, error) {
	result, err := r.tasks.UpdateTask(task)
	log.Print("TaskRepository.UpdateTask:", result, err)
	return result, err
}

func (r TaskRepository) CreateTask(task entities.Task) (entities.Task, error) {
	result, err := r.tasks.CreateTask(task)
	log.Print("TaskRepository.CreateTask:", result, err)
	return result, err
}

func (r TaskRepository) RemoveTask(uniqueId string) {
	r.RemoveTask(uniqueId)
	log.Print("TaskRepository.RemoveTask:", uniqueId)
}

func NewTaskRepository(tasks repositories.TaskRepository) repositories.TaskRepository {
	return TaskRepository{
		tasks,
	}
}
