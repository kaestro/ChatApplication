// myapp/api/handlers/chatHandler/enterRoom.go
package chatHandler

import (
	"myapp/api/service/chatService"
	"net/http"

	"github.com/gin-gonic/gin"
)

// request type: GET
// Header: Upgrade, Connection, Sec-WebSocket-Version, Sec-WebSocket-Key, Session-Key
// Body: LoginInfo { EmailAddress, Password, LoginSessionID }
func EnterChat(c *gin.Context) {
	if err := chatService.ValidateUpgradeHeader(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginInfo, err := chatService.ParseEnterChatAndAuthenticateUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := chatService.EnterChat(c, loginInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Entered room successfully"})
}
