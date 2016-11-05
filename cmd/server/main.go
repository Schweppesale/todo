package main

import (
	"github.com/schweppesale/todo/cmd/server/api"
	"github.com/schweppesale/todo/cmd/server/http"
	"github.com/schweppesale/todo/infrastructure/logger"
	"github.com/schweppesale/todo/infrastructure/memory"
	"github.com/schweppesale/todo/infrastructure/satori"
)

func main() {
	uuidGenerator := satori.NewUUIDGenerator()
	tasks := logger.NewTaskRepository(memory.NewTaskRepository(uuidGenerator))
	taskService := api.NewTaskService(tasks)
	server := http.NewServer(taskService)
	server.Run()
}
