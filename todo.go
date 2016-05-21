package main

import "time"

/*
Todo model
*/
type Todo struct {
	Index     int       `json:"index"`
	Title     string    `json:"title"`
	Due       time.Time `json:"due"`
	Completed bool      `json:"completed"`
}

/*
Todos array of Todo models
*/
type Todos []Todo
