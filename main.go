package main 

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

type todo struct {
	ID			string `json:"id"`
	Name		string `json:"title"`
	Completed	bool   `json:"completed"`
} 

var todos = []todo {
	{ID:"1",Name:"Clean Room",Completed:false},	
	{ID:"2",Name:"Read Book",Completed:false},	
	{ID:"3",Name:"Record Video",Completed:false},	
}

//var welcomeMessage = "Welcome"
//IndentedJSON will accept the code , and a string , num , slive , struct , or map
func getTodos(context *gin.Context) {
	log.Println("\033[33m getTodos function called\033[33m")
	context.IndentedJSON(http.StatusOK, todos)
}
func getWelcome(context *gin.Context) {
	log.Println("\033[33m getWelcome function called\033[33m")
	context.IndentedJSON(http.StatusOK,"HELLO TEST") //the second message is what is going to be displayed to the screen
}

func main() {
	log.Println("Starting Program")
	router := gin.Default() 
	router.GET("/",getWelcome)
	router.GET("/todos",getTodos) //get request to /todos will run getTodos
	router.Run("localhost:9090")	//router will be listening on 127.0.0.1 on port 9090 for http requests
}