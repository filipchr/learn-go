package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Method      string
	Pattern     string
	Name        string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {

		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{
		"Get",
		"/",
		"Index",
		Index,
	},
	Route{
		"Get",
		"/todos",
		"TodosIndex",
		TodoIndex,
	},
	Route{
		"Get",
		"/todos/{todoId}",
		"TodoShowIndex",
		TodoShow,
	},
}
