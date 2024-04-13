// myapp/api/service/chatService/parseEnterRoomRequest.go
package chatService

import (
	"encoding/base64"
	"myapp/api/models"

	"github.com/gin-gonic/gin"
)

func IsUpgradeHeaderValid(c *gin.Context) bool {
	upgradeHeader := c.GetHeader("Upgrade")
	connectionHeader := c.GetHeader("Connection")
	secWebSocketVersionHeader := c.GetHeader("Sec-WebSocket-Version")
	secWebSocketKeyHeader := c.GetHeader("Sec-WebSocket-Key")

	isSecWebSocketKeyValid := false
	decodedKey, err := base64.StdEncoding.DecodeString(secWebSocketKeyHeader)
	if err == nil && len(decodedKey) == 16 {
		isSecWebSocketKeyValid = true
	}

	return upgradeHeader == "websocket" &&
		connectionHeader == "upgrade" &&
		secWebSocketVersionHeader == "13" &&
		isSecWebSocketKeyValid
}

func ParseEnterRoomRequest(c *gin.Context) (models.RoomRequest, error) {
	var req models.RoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return models.RoomRequest{}, err
	}
	return req, nil
}

func ParseEnterChatRequest(c *gin.Context) (models.LoginInfo, error) {
	var loginInfo models.LoginInfo
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		return models.LoginInfo{}, err
	}
	return loginInfo, nil
}
