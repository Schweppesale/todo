package services

import (
	"github.com/schweppesale/todo/application/mappers"
	"github.com/schweppesale/todo/application/response"
	"github.com/schweppesale/todo/domain/entities"
	"github.com/schweppesale/todo/domain/repositories"
)

type TaskService struct {
	tasks  repositories.TaskRepository
	mapper mappers.TaskMapper
}

func (ts TaskService) FindAll() (map[string]response.TaskResponse, error) {
	tasks, err := ts.tasks.FindAll()
	result := make(map[string]response.TaskResponse, len(tasks))
	for k, v := range tasks {
		result[k] = ts.mapper.MapTaskResponse(v)
	}
	return result, err
}

func (ts TaskService) GetTaskByUniqueId(uniqueId string) (response.TaskResponse, error) {
	task, err := ts.tasks.GetTaskByUniqueId(uniqueId)
	return ts.mapper.MapTaskResponse(task), err
}

func (ts TaskService) CreateTask(title string, description string) (response.TaskResponse, error) {
	task, err := ts.tasks.CreateTask(entities.NewTask(title, description))
	return ts.mapper.MapTaskResponse(task), err
}

func (ts TaskService) UpdateTask(uniqueId string, title string) (response.TaskResponse, error) {
	task, err := ts.tasks.GetTaskByUniqueId(uniqueId)
	if err != nil {
		return response.TaskResponse{}, err
	}

	task.SetTitle(title)
	newTask, err := ts.tasks.UpdateTask(task)
	return ts.mapper.MapTaskResponse(newTask), err
}

func (ts TaskService) RemoveTask(uniqueId string) {
	ts.tasks.RemoveTask(uniqueId)
}

func NewTaskService(tasks repositories.TaskRepository, mapper mappers.TaskMapper) TaskService {
	return TaskService{
		tasks,
		mapper,
	}
}
