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

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func main() {
	r := gin.Default()

	r.GET("/todos", getTodos)

	r.Run("localhost:8989") // Run on port 8989
}