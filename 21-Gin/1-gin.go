package main

import (
	"net/http"

	"github.com/gin-gonic/gin" //gin framework installation:- go get -u github.com/gin-gonic/gin
)

type User struct {
	Name string `json:name`
}

func main() {
	router := gin.Default()

	router.GET("/home", Home)
	router.Run(":8080")
}

// func(*Context)
func Home(c *gin.Context) {
	u := User{
		Name: "John",
	}
	//c.String(http.StatusOK, "this is my home page")
	//using the map to send the json response
	// c.JSON(http.StatusOK, gin.H{"msg": "this is my home page"})
	c.JSON(http.StatusOK, gin.H{"user": u})
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{}) //to send error and serve the request

}
