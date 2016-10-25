package server

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
  "github.com/user/todo/infrastructure/repositories/memory"
)

func Run() {

  var Tasks = new(memory.TaskRepository)
  http.HandleFunc("/api/todo/tasks", func(w http.ResponseWriter, r *http.Request) {
    tasks, err := json.Marshal(Tasks.FindAll())
    if err != nil {
      fmt.Println(err)
      return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(tasks)
  })

  log.Fatal(http.ListenAndServe(":8080", nil))
}
