package todo

import (
	"time"
)

// Task is a domain entity
type Task struct {
	uniqueID    string
	title       string
	description string
	createdOn   time.Time
	updatedOn   time.Time
}

// NewTask serves as a Task constructor
func NewTask(title string, description string) Task {
	return Task{
		"",
		title,
		description,
		time.Now(),
		time.Now(),
	}
}

// SetUniqueID assigns this task a unique ID
func (t *Task) SetUniqueID(uniqueID string) {
	t.uniqueID = uniqueID
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
	t.updatedOn = time.Now()
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
	FindAll() ([]Task, error)
	GetByUniqueID(uniqueID string) (Task, error)
	Create(task Task) (Task, error)
	Update(task Task) (Task, error)
	Remove(uniqueID string) error
}
