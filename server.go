package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/schweppesale/todo/application/services"
	"log"
	"net/http"
	"path"
)

type Server struct {
	taskService     services.TaskService
	responseHandler HttpResponseHandler
}

func (server Server) Run() {

	router := mux.NewRouter()
	router.HandleFunc("/api/todo/tasks", func(response http.ResponseWriter, request *http.Request) {
		switch {
		case request.Method == "GET":
			tasks, err := server.taskService.FindAll()
			server.responseHandler.Format(response, tasks, err)
			break
		case request.Method == "POST":
			request.ParseForm()
			title := request.Form.Get("title")
			description := request.Form.Get("description")
			task, err := server.taskService.CreateTask(title, description)
			server.responseHandler.Format(response, task, err)
			break
		}
	})

	router.HandleFunc("/api/todo/tasks/{taskId}", func(response http.ResponseWriter, request *http.Request) {
		uniqueId := path.Base(request.RequestURI)
		switch {
		case request.Method == "GET":
			task, err := server.taskService.GetTaskByUniqueId(uniqueId)
			server.responseHandler.Format(response, task, err)
			break
		case request.Method == "DELETE":
			server.taskService.RemoveTask(uniqueId)
			server.responseHandler.Format(response, "success true", nil)
			break
		}
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}

func NewServer(tasks services.TaskService, responseHandler HttpResponseHandler) Server {
	return Server{
		tasks,
		responseHandler,
	}
}

type HttpResponseHandler interface {
	Format(response http.ResponseWriter, payload interface{}, err error)
}

type JsonResponseHandler struct{}

func (handler JsonResponseHandler) Format(response http.ResponseWriter, payload interface{}, err error) {
	response.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(response, err.Error(), http.StatusNotFound)
	} else {
		serialized, _ := json.Marshal(payload)
		response.Write(serialized)
	}
}

func NewJsonResponseHandler() JsonResponseHandler {
	return JsonResponseHandler{}
}
