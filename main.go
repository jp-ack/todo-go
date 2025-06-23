package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Name      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Name: "Clean Room", Completed: false},
	{ID: "2", Name: "Read Book", Completed: false},
	{ID: "3", Name: "Record Video", Completed: false},
}

// IndentedJSON will accept the code , and a string , num , slice , struct , or map
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo
	// Try to bind the incoming JSON to newTodo
	if err := context.BindJSON(&newTodo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusAccepted, newTodo)
}

func delTodo(context *gin.Context) {

}

func main() {

	router := gin.Default()
	router.Run("localhost:9090") //router will be listening on 127.0.0.1 on port 9090 for http requests

	router.GET("/todos", getTodos) //get request to /todos will run getTodos, passing in the incomming http request as gin.Context
	router.POST("/todos", addTodo)
	router.DELETE("/todos", delTodo)

}
