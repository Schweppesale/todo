package api

import (
	"github.com/schweppesale/todo"
	"time"
)

type TaskService struct {
	tasks todo.TaskRepository
}

func NewTaskService(tasks todo.TaskRepository) TaskService {
	return TaskService{
		tasks,
	}
}

func (ts TaskService) FindAll() ([]Task, error) {
	tasks, err := ts.tasks.FindAll()
	result := make([]Task, len(tasks))
	for k, v := range tasks {
		result[k] = NewTask(v)
	}
	return result, err
}

func (ts TaskService) GetByUniqueID(uniqueID string) (Task, error) {
	task, err := ts.tasks.GetByUniqueID(uniqueID)
	return NewTask(task), err
}

func (ts TaskService) Create(title string, description string) (Task, error) {
	task, err := ts.tasks.Create(todo.NewTask(title, description))
	return NewTask(task), err
}

func (ts TaskService) Update(uniqueID string, title string, description string) (Task, error) {
	task, err := ts.tasks.GetByUniqueID(uniqueID)
	if err != nil {
		return Task{}, err
	}
	task.SetTitle(title)
	task.SetDescription(description)
	newTask, err := ts.tasks.Update(task)
	return NewTask(newTask), err
}

func (ts TaskService) Remove(uniqueID string) error {
	return ts.tasks.Remove(uniqueID)
}

type Task struct {
	UniqueID    string
	Title       string
	Description string
	CreatedOn   time.Time
	UpdatedOn   time.Time
}

func NewTask(task todo.Task) Task {
	return Task{
		task.UniqueID(),
		task.Title(),
		task.Description(),
		task.CreatedOn(),
		task.UpdatedOn(),
	}
}
