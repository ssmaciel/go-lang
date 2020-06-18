package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type Task struct {
	Name string
	Done bool
}

func CreateTasks() []Task {
	tasks := []Task{
		{"Tarefa 1", true},
		{"Tarefa 2", true},
		{"Tarefa 3", false},
		{"Tarefa 4", false},
		{"Tarefa 5", true},
		{"Tarefa 6", false},
	}
	return tasks
}

func taskApiHhandler(w http.ResponseWriter, r *http.Request) {
	tasks := CreateTasks()
	 js, err := json.Marshal(tasks)
	 if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	  }
	 w.Header().Set("Content-Type", "application/json")
	 w.Write(js)
}

func taskViewHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("tasks.html"))
	
	tasks := []Task{
		{Name:"Tarefa 1", Done: true},
		{"Tarefa 2", true},
		{"Tarefa 3", false},
		{"Tarefa 4", false},
		{"Tarefa 5", true},
		{"Tarefa 6", false},
	}


	tmpl.Execute(w, tasks)
}

func main() {
	http.HandleFunc("/tasks", taskViewHandler)
	http.HandleFunc("/api/tasks", taskApiHhandler)
	fmt.Println("Servidor iniciado...")
	http.ListenAndServe(":5000", nil)
}
