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

func (ts TaskService) FindTasks() ([]Task, error) {
	tasks, err := ts.tasks.FindTasks()
	result := make([]Task, len(tasks))
	for k, v := range tasks {
		result[k] = NewTask(v)
	}
	return result, err
}

func (ts TaskService) GetTaskByUniqueID(uniqueID string) (Task, error) {
	task, err := ts.tasks.GetTaskByUniqueID(uniqueID)
	return NewTask(task), err
}

func (ts TaskService) CreateTask(title string, description string) (Task, error) {
	task, err := ts.tasks.CreateTask(todo.NewTask(title, description))
	return NewTask(task), err
}

func (ts TaskService) UpdateTask(uniqueID string, title string, description string) (Task, error) {
	task, err := ts.tasks.GetTaskByUniqueID(uniqueID)
	if err != nil {
		return Task{}, err
	}
	task.SetTitle(title)
	task.SetDescription(description)
	newTask, err := ts.tasks.UpdateTask(task)
	return NewTask(newTask), err
}

func (ts TaskService) RemoveTask(uniqueID string) error {
	return ts.tasks.RemoveTask(uniqueID)
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
