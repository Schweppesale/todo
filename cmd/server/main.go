package main

import (
	"github.com/schweppesale/todo/cmd/server/api"
	"github.com/schweppesale/todo/cmd/server/http"
	"github.com/schweppesale/todo/tasks/logger"
	"github.com/schweppesale/todo/tasks/memory"
	"github.com/schweppesale/todo/tasks/memory/satori"
)

func main() {
	uuidGenerator := satori.NewUUIDGenerator()
	tasks := logger.NewTaskRepository(memory.NewTaskRepository(uuidGenerator))
	taskService := api.NewTaskService(tasks)
	server := http.NewServer(taskService)
	server.Run()
}
