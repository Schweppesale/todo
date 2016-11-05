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

func (r TaskRepository) FindTasks() ([]todo.Task, error) {
	log.Print("TaskRepository.FindAll:")
	result, err := r.tasks.FindTasks()
	log.Print("TaskRepository.FindAll:", result)
	return result, err
}

func (r TaskRepository) GetTaskByUniqueId(uniqueId string) (todo.Task, error) {
	log.Print("TaskRepository.GetTaskByUniqueId:", uniqueId)
	result, err := r.tasks.GetTaskByUniqueId(uniqueId)
	log.Print("TaskRepository.GetTaskByUniqueId:", result, err)
	return result, err
}

func (r TaskRepository) UpdateTask(task todo.Task) (todo.Task, error) {
	log.Print("TaskRepository.UpdateTask:", task)
	result, err := r.tasks.UpdateTask(task)
	log.Print("TaskRepository.UpdateTask:", result, err)
	return result, err
}

func (r TaskRepository) CreateTask(task todo.Task) (todo.Task, error) {
	log.Print("TaskRepository.CreateTask:", task)
	result, err := r.tasks.CreateTask(task)
	log.Print("TaskRepository.CreateTask:", result, err)
	return result, err
}

func (r TaskRepository) RemoveTask(uniqueId string) error {
	log.Print("TaskRepository.RemoveTask:", uniqueId)
	return r.tasks.RemoveTask(uniqueId)
}
