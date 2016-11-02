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
	service := services.NewTaskService(tasks, mapper)

	server := NewServer(service)
	server.Run()
}
