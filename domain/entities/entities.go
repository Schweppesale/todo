package entities

import (
	"time"
)

type Task struct {
	unique_id   string
	title       string
	description string
	created_on  time.Time
	updated_on  time.Time
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

func (t *Task) SetUniqueID(uniqueId string) {
	t.unique_id = uniqueId
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
	return t.unique_id
}

func (t Task) UpdatedOn() time.Time {
	return t.updated_on
}

func (t Task) CreatedOn() time.Time {
	return t.created_on
}
