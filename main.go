package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: 1, Title: "Learn Go", Completed: false},
	{ID: 2, Title: "Build a REST API", Completed: false},
	{ID: 3, Title: "Deploy to production", Completed: false},
}

func main() {
	r := gin.Default()

	r.GET("/todos", getTodos)
	r.POST("/todos", createTodo)
	r.PUT("/todos/:id", updateTodo)
	r.DELETE("/todos/:id", deleteTodo)

	r.Run("localhost:8989") // Run on port 8080
}