package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/schweppesale/todo/domain/entities"
	"github.com/schweppesale/todo/domain/repositories"
	"github.com/schweppesale/todo/domain/services"
	"log"
	"net/http"
	"path"
)

func Run(Tasks repositories.TaskRepository, UUIDGen services.UniqueIdGenerator) {

	r := mux.NewRouter()
	r.HandleFunc("/api/todo/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == "GET":
			w.Write(serialize(Tasks.FindAll()))
			break
		case r.Method == "POST":
			r.ParseForm()
			w.Write(serialize(Tasks.SaveTask(entities.NewTask(UUIDGen, r.Form.Get("title")))))
			break
		}
	})

	r.HandleFunc("/api/todo/tasks/{taskId}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		uniqueId := path.Base(r.RequestURI)
		switch {
		case r.Method == "GET":
			w.Write(serialize(Tasks.GetTaskByUniqueId(uniqueId)))
			break
		case r.Method == "DELETE":
			task, _ := Tasks.GetTaskByUniqueId(uniqueId)
			Tasks.RemoveTask(task.UniqueId())
		}
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}

func serialize(payload interface{}, err error) []byte {
	if err != nil {
		response, _ := json.Marshal(err)
		return response
	}
	response, _ := json.Marshal(payload)
	return response

}
