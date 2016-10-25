package main

import (
	"github.com/schweppesale/todo/application/server"
	"github.com/schweppesale/todo/infrastructure/repositories/memory"
	"github.com/schweppesale/todo/infrastructure/services/unique_id_generator/satori"
)

func main() {
	//@todo CLI arguments
	server.Run(new(memory.TaskRepository), new(satori.UniqueIdGenerator))
}
