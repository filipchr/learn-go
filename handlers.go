package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Everything is GO")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Index: 1, Title: "learn Go!", Completed: false},
		Todo{Index: 2, Title: "Drink coffe", Completed: false},
	}

	err := json.NewEncoder(w).Encode(todos)
	if err != nil {
		panic(err)
	}

	//fmt.Fprintln(w, "Show all todos", jsonString)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Show", todoId)
}
