package todo

import (
	"time"
)

type Task struct {
	uniqueId    string
	title       string
	description string
	createdOn   time.Time
	updatedOn   time.Time
}

func NewTask(title string, description string) Task {
	return Task{
		"",
		title,
		description,
		time.Now(),
		time.Now(),
	}
}

func (t *Task) SetUniqueID(generator UUIDGenerator) {
	t.uniqueId = generator.Generate()
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}

func (t *Task) SetTitle(title string) {
	t.title = title
}

func (t *Task) SetDescription(description string) {
	t.description = description
}

func (t Task) UniqueId() string {
	return t.uniqueId
}

func (t Task) UpdatedOn() time.Time {
	return t.updatedOn
}

func (t Task) CreatedOn() time.Time {
	return t.createdOn
}

type TaskRepository interface {
	FindTasks() ([]Task, error)
	GetTaskByUniqueId(uniqueId string) (Task, error)
	CreateTask(task Task) (Task, error)
	UpdateTask(task Task) (Task, error)
	RemoveTask(uniqueId string) error
}
