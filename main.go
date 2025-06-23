package main

import (
	"log"
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

// var welcomeMessage = "Welcome"
// IndentedJSON will accept the code , and a string , num , slice , struct , or map
func getTodos(context *gin.Context) {
	log.Println("\033[33m getTodos function called\033[33m")
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	log.Println("Add Todo Clicked ")
	var newTodo todo

	if err := context.BindJSON(&newTodo); // bind the JSON , if there is an error , it wont be nil
	err != nil {                          // if it is not nil , we will return
		return //hence we return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusAccepted, newTodo)

}

func main() {
	log.Println("Starting Program")
	router := gin.Default()
	router.GET("/todos", getTodos) //get request to /todos will run getTodos
	router.POST("/todos", addTodo)
	router.Run("localhost:9090") //router will be listening on 127.0.0.1 on port 9090 for http requests
}
