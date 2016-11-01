package mappers

import (
	"github.com/schweppesale/todo/application/response"
	"github.com/schweppesale/todo/domain/entities"
)

type TaskMapper struct {
}

func (ts TaskMapper) MapTaskResponse(task entities.Task) response.TaskResponse {
	return response.TaskResponse{
		task.UniqueId(),
		task.Title(),
		task.UpdatedOn(),
	}
}

func NewTaskMapper() TaskMapper {
	return TaskMapper{}
}
