package todo

import (
	"time"
)

// Task struct serves as an domain entity
type Task struct {
	uniqueID    string
	title       string
	description string
	createdOn   time.Time
	updatedOn   time.Time
}

// NewTask serves as a constructor
func NewTask(title string, description string) Task {
	return Task{
		"",
		title,
		description,
		time.Now(),
		time.Now(),
	}
}

// SetUniqueID invokes an interface which generates a new UUID
func (t *Task) SetUniqueID(generator UUIDGenerator) {
	t.uniqueID = generator.Generate()
	t.updatedOn = time.Now()
}

// Title returns the Task title
func (t Task) Title() string {
	return t.title
}

// Description returns the Task description
func (t Task) Description() string {
	return t.description
}

// SetTitle changes the title
func (t *Task) SetTitle(title string) {
	t.title = title
	t.updatedOn = time.Now()
}

// SetDescription changes the description
func (t *Task) SetDescription(description string) {
	t.description = description
}

// UniqueID returns the Task ID
func (t Task) UniqueID() string {
	return t.uniqueID
}

// UpdatedOn returns the Task's last modification date
func (t Task) UpdatedOn() time.Time {
	return t.updatedOn
}

// CreatedOn returns the Task's creation date
func (t Task) CreatedOn() time.Time {
	return t.createdOn
}

// TaskRepository serves as a layer of abstraction around our persistence layer
type TaskRepository interface {
	FindTasks() ([]Task, error)
	GetTaskByUniqueID(uniqueID string) (Task, error)
	CreateTask(task Task) (Task, error)
	UpdateTask(task Task) (Task, error)
	RemoveTask(uniqueID string) error
}
