package repositories

import (
	"github.com/schweppesale/todo/domain/entities"
)

type TaskRepository interface {
	FindAll() (map[string]entities.Task, error)
	GetTaskByUniqueId(uniqueId string) (entities.Task, error)
	SaveTask(task entities.Task) (entities.Task, error)
	RemoveTask(uniqueId string)
}