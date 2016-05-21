package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

/*
Route model
*/
type Route struct {
	Method      string
	Pattern     string
	Name        string
	HandlerFunc http.HandlerFunc
}

/*
Routes is an array of Route Models
*/
type Routes []Route

/*
NewRouter is the main router function
*/
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

// Construct the routes
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
