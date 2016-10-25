package repositories

import (
	"github.com/user/todo/domain/entities"
)

type TaskRepository interface{
	FindAll()
	GetTaskByUniqueId(uniqueId string) entities.Task
	SaveTask(task entities.Task) entities.Task
	RemoveTask(uniqueId string)
}