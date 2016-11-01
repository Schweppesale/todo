package entities

import (
	"time"
)

type Task struct {
	unique_id  string
	title      string
	created_on string
	updated_on string
}

func (t *Task) SetUniqueID(uniqueId string) {
	t.unique_id = uniqueId
}

func (t Task) Title() string {
	return t.title
}

func (t *Task) ChangeTitle(title string) {
	t.title = title
}

func (t Task) UniqueId() string {
	return t.unique_id
}

func (t Task) UpdatedOn() string {
	return t.updated_on
}

func NewTask(title string) Task {
	return Task{
		"",
		title,
		time.Now().Format(time.UnixDate),
		time.Now().Format(time.UnixDate),
	}
}
