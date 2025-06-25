package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"errors"

	"strconv"
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

func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	newTodo.ID = len(todos) + 1
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")

	todo, err := getTodoByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")

	todo, err := getTodoByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

func getTodoByID(id string) (*todo, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}
	for i, t := range todos {
		if t.ID == intID {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func main() {
	r := gin.Default()
	r.GET("/todos", getTodos)
	r.POST("/todos", addTodo)
	r.PATCH("/todos/:id", toggleTodoStatus)
	r.GET("/todos/:id", getTodo)
	r.Run("localhost:8989") // Run on port 8989
}