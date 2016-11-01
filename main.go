package main

import (
	"github.com/schweppesale/todo/application/mappers"
	"github.com/schweppesale/todo/application/services"
	"github.com/schweppesale/todo/infrastructure/repositories/logger"
	"github.com/schweppesale/todo/infrastructure/repositories/memory"
	"github.com/schweppesale/todo/infrastructure/services/unique_id_generator/satori"
)

func main() {
	mapper := mappers.NewTaskMapper()
	uuidGen := new(satori.UniqueIdGenerator)
	tasks := logger.NewTaskRepository(memory.NewTaskRepository(uuidGen))
	service := services.NewTaskService(tasks, mapper)

	server := NewServer(service)
	server.Run()
}
