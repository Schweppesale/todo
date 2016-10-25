package server

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/schweppesale/todo/domain/repositories"
	"github.com/schweppesale/todo/domain/entities"
	"github.com/schweppesale/todo/domain/services"
	"path"
	"github.com/gorilla/mux"
)

func Run(Tasks repositories.TaskRepository, UUIDGen services.UniqueIdGenerator) {

	r := mux.NewRouter()
	r.HandleFunc("/api/todo/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == "GET":
			tasks, _ := Tasks.FindAll()
			response, err := json.Marshal(tasks)
			if (err != nil) {
				log.Print(err)
			}
			w.Write(response)
			break
		case r.Method == "POST":
			r.ParseForm()
			task, _ := Tasks.SaveTask(entities.NewTask(UUIDGen, r.Form.Get("title")))
			response, err := json.Marshal(task)
			if (err != nil) {
				log.Print(err)
			}
			w.Write(response)
			break
		}
	})

	r.HandleFunc("/api/todo/tasks/{taskId}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		uniqueId := path.Base(r.RequestURI)
		switch {
		case r.Method == "GET":
			task, err := Tasks.GetTaskByUniqueId(uniqueId)
			if(err != nil) {
				w.Write(err)
				return
			}
			response, err := json.Marshal(task)
			if (err != nil) {
				log.Print(err)
				return
			}
			w.Write(response)
			break
		case r.Method == "DELETE":
			task, _ := Tasks.GetTaskByUniqueId(uniqueId)
			Tasks.RemoveTask(task.UniqueId())
		}
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}