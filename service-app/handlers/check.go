package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Check(c *gin.Context) {
	c.JSON(http.StatusOK, "Check handler called!")
}
