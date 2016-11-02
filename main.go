package main

import (
	"github.com/schweppesale/todo/application/mappers"
	"github.com/schweppesale/todo/application/services"
	"github.com/schweppesale/todo/infrastructure/repositories/logger"
	"github.com/schweppesale/todo/infrastructure/repositories/memory"
	"github.com/schweppesale/todo/infrastructure/services/uuid/satori"
)

func main() {
	mapper := mappers.NewTaskMapper()
	uuidService := satori.NewUuidService()
	tasks := logger.NewTaskRepository(memory.NewTaskRepository(uuidService))
	taskService := services.NewTaskService(tasks, mapper)
	responseHandler := NewJsonResponseHandler()

	server := NewServer(taskService, responseHandler)
	server.Run()
}
