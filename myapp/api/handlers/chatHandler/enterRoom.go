// myapp/api/handlers/chatHandler/enterRoom.go
package chatHandler

import (
	"myapp/api/models"
	"myapp/api/service/chatService"
	"net/http"

	"github.com/gin-gonic/gin"
)

// myapp/api/handlers/chatHandler/enterRoom.go
func EnterChat(c *gin.Context) {
	if err := chatService.ValidateUpgradeHeader(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req, err := chatService.ParseChatRequestAndAuthenticateUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Refactor this to use the new EnterChatRoom function
	loginInfo := models.NewLoginInfo(req.EmailAddress, req.Password)
	loginInfo.LoginSessionID = req.LoginSessionID

	if err := chatService.EnterChat(c, loginInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Entered room successfully"})
}
