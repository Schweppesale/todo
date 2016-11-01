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
	tasks services.TaskService
}

func (s Server) Run() {

	r := mux.NewRouter()
	r.HandleFunc("/api/todo/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == "GET":
			w.Write(s.serialize(s.tasks.FindAll()))
			break
		case r.Method == "POST":
			r.ParseForm()
			w.Write(s.serialize(s.tasks.CreateTask(r.Form.Get("title"))))
			break
		}
	})

	r.HandleFunc("/api/todo/tasks/{taskId}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		uniqueId := path.Base(r.RequestURI)
		switch {
		case r.Method == "GET":
			w.Write(s.serialize(s.tasks.GetTaskByUniqueId(uniqueId)))
			break
		case r.Method == "DELETE":
			task, _ := s.tasks.GetTaskByUniqueId(uniqueId)
			s.tasks.RemoveTask(task.UniqueId)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}

func (s Server) serialize(payload interface{}, err error) []byte {
	if err != nil {
		response, _ := json.Marshal(err)
		return response
	}
	response, _ := json.Marshal(payload)
	return response

}

func NewServer(tasks services.TaskService) Server {
	return Server{
		tasks,
	}
}
