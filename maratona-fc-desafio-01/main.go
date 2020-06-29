package main

import (
	"fmt"
	"html/template"
	"net/http"
)


func taskViewHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("main.html"))
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", taskViewHandler)
	fmt.Println("Servidor iniciado...")
	http.ListenAndServe(":5000", nil)
}