package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/schweppesale/todo/cmd/server/api"
	"log"
	"net/http"
	"path"
)

type Server struct {
	taskService api.TaskService
}

func NewServer(taskService api.TaskService) Server {
	return Server{
		taskService,
	}
}

func (server Server) Run() {

	router := mux.NewRouter()
	router.HandleFunc("/api/todo/tasks", func(response http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			tasks, err := server.taskService.FindTasks()
			FormatResponse(response, tasks, err)
			break
		case "POST":
			request.ParseForm()
			title := request.Form.Get("title")
			description := request.Form.Get("description")
			task, err := server.taskService.CreateTask(title, description)
			FormatResponse(response, task, err)
			break
		}
	})

	router.HandleFunc("/api/todo/tasks/{taskId}", func(response http.ResponseWriter, request *http.Request) {
		uniqueID := path.Base(request.RequestURI)
		switch request.Method {
		case "GET":
			task, err := server.taskService.GetTaskByUniqueID(uniqueID)
			FormatResponse(response, task, err)
			break
		case "PATCH":
			request.ParseForm()
			title := request.Form.Get("title")
			description := request.Form.Get("description")
			task, err := server.taskService.UpdateTask(uniqueID, title, description)
			FormatResponse(response, task, err)
			break
		case "DELETE":
			err := server.taskService.RemoveTask(uniqueID)
			FormatResponse(response, "success true", err)
			break
		}
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}

func FormatResponse(response http.ResponseWriter, payload interface{}, err error) {
	response.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(response, err.Error(), http.StatusNotFound)
		return
	}

	serialized, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("FATAL ERROR: ", err)
		return
	}

	response.Write(serialized)
}
