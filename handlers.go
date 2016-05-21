package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*
Index main router
*/
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Everything is GO")
}

/*
TodoIndex shows all todos
*/
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Index: 1, Title: "learn Go!", Completed: false},
		Todo{Index: 2, Title: "Drink coffe", Completed: false},
	}

	err := json.NewEncoder(w).Encode(todos)
	if err != nil {
		panic(err)
	}

}

/*
TodoShow serves a single todo
*/
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoId"]
	fmt.Fprintln(w, "Show", todoID)
}
