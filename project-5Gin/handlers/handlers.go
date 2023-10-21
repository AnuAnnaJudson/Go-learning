package handlers

import (
	"log"
	"net/http"
	"project-5/models"
	"strconv"

	"github.com/gin-gonic/gin" //gin framework installation:- go get -u github.com/gin-gonic/gin
)

func GetUser(c *gin.Context) {

	//set content type

	userIdString := c.Param("userId") // this will fetch the

	//convert this to uint (strconv.ParseUint)
	userId, err := strconv.ParseUint(userIdString, 10, 64)

	// if error happens report to the client in json // write status code
	if err != nil {

		log.Println("Error: ", err)

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "Invalid string"})

		return

	}

	//appErr := map[string]string{"Message": http.StatusText(http.StatusBadRequest)}

	//call the FetchUser function from the models package

	uData, err := models.FetchUser(userId)

	if err != nil {

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})

		return

	}

	// write to the client with user data or error

	c.JSON(http.StatusOK, uData)

}
