// /myapp/api/handlers/handlePing.go
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlePing(c *gin.Context) {
	response := gin.H{
		"message": "pong",
	}

	if response["message"] != "pong" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ping request failed"})
		return
	}

	c.JSON(http.StatusOK, response)
}
