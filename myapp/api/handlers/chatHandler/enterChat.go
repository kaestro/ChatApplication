// myapp/api/handlers/chatHandler/enterRoom.go
package chatHandler

import (
	"errors"
	"myapp/api/service/chatService"
	"myapp/api/service/userService"
	"net/http"

	"github.com/gin-gonic/gin"
)

// request type: GET
// Header: Upgrade, Connection, Sec-WebSocket-Version, Sec-WebSocket-Key, Session-Key
// Body: LoginInfo { EmailAddress, Password, LoginSessionID }
func EnterChat(c *gin.Context) {
	loginSessionInfo, err := chatService.ParseEnterLoginSessionInfo(c)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userServiceUtil := userService.NewUserServiceUtil()
	if err = userServiceUtil.AuthenticateUserByLoginSessionInfo(loginSessionInfo); err != nil {
		c.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"authentication error": err.Error()})
		return
	}

	hijacker, ok := c.Writer.(http.Hijacker)
	if !ok {
		c.Error(errors.New("the response writer does not support the Hijacker interface"))
		return
	}

	_, _, err = hijacker.Hijack()
	if err != nil {
		c.Error(err)
		return
	}

	// Handle the WebSocket connection
	if err := chatService.EnterChat(c, loginSessionInfo); err != nil {
		c.Error(err)
		return
	}
}
