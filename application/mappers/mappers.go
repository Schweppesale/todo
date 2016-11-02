package mappers

import (
	"github.com/schweppesale/todo/application/response"
	"github.com/schweppesale/todo/domain/entities"
	"time"
)

type TaskMapper struct {
}

func NewTaskMapper() TaskMapper {
	return TaskMapper{}
}

func (ts TaskMapper) MapTaskResponse(task entities.Task) response.TaskResponse {
	return response.TaskResponse{
		task.UniqueId(),
		task.Title(),
		task.Description(),
		task.UpdatedOn().Format(time.UnixDate),
		task.CreatedOn().Format(time.UnixDate),
	}
}
