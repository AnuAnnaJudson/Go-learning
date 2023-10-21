package main

//run main.go then in cmd do curl localhost:8080/getuser or open this url in browser
import (
	"project-5/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() //create router using gin

	router.GET("/user/:userId", handlers.GetUser) //in get arguments are endpoit and handler function
	router.Run(":8080")                           //starts server at the port passed to it
}
