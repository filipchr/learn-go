package main

import "time"

type Todo struct {
	Index     int       `json:"index"`
	Title     string    `json:"title"`
	Due       time.Time `json:due`
	Completed bool      `json:"completed"`
}

type Todos []Todo
