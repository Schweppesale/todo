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

func (ts TaskService) GetTaskByUniqueId(uniqueId string) (Task, error) {
	task, err := ts.tasks.GetTaskByUniqueId(uniqueId)
	return NewTask(task), err
}

func (ts TaskService) CreateTask(title string, description string) (Task, error) {
	task, err := ts.tasks.CreateTask(todo.NewTask(title, description))
	return NewTask(task), err
}

func (ts TaskService) UpdateTask(uniqueId string, title string, description string) (Task, error) {
	task, err := ts.tasks.GetTaskByUniqueId(uniqueId)
	if err != nil {
		return Task{}, err
	}
	task.SetTitle(title)
	task.SetDescription(description)
	newTask, err := ts.tasks.UpdateTask(task)
	return NewTask(newTask), err
}

func (ts TaskService) RemoveTask(uniqueId string) error {
	return ts.tasks.RemoveTask(uniqueId)
}

type Task struct {
	UniqueId    string
	Title       string
	Description string
	CreatedOn   time.Time
	UpdatedOn   time.Time
}

func NewTask(task todo.Task) Task {
	return Task{
		task.UniqueId(),
		task.Title(),
		task.Description(),
		task.CreatedOn(),
		task.UpdatedOn(),
	}
}
