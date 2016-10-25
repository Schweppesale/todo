package entities

import (
	"time"
	"github.com/schweppesale/todo/domain/services"
)

type Task struct {
	unique_id  string
	title      string
	created_on string
	updated_on string
}

func (t Task) UniqueId() string {
	return t.unique_id
}

func NewTask(UUIDGen services.UniqueIdGenerator, title string) Task {
	return Task{
		UUIDGen.Generate(),
		title,
		time.Now().Format(time.UnixDate),
		time.Now().Format(time.UnixDate),
	}
}