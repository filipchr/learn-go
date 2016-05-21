package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Setup the router
	r := mux.NewRouter().StrictSlash(true)

	// Attach the routes
	r.HandleFunc("/", Index)
	r.HandleFunc("/todos", Todos)
	r.HandleFunc("/todos/{todoId}", TodoShow)

	// Listen on port :8080
	log.Fatal(http.ListenAndServe(":8080", r))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Everything is GO")
}

func Todos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Show all todos")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Show", todoId)
}
