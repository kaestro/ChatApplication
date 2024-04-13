// myapp/api/handlers/chatHandler/enterRoom.go
package chatHandler

import (
	"myapp/api/service/chatService"
	"net/http"

	"github.com/gin-gonic/gin"
)

// myapp/api/handlers/chatHandler/enterRoom.go
func EnterRoom(c *gin.Context) {
	if err := chatService.ValidateUpgradeHeader(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req, err := chatService.ParseAndAuthenticateRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := chatService.EnterChatRoom(c, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Entered room successfully"})
}
