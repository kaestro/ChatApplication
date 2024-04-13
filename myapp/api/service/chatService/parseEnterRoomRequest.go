// myapp/api/service/chatService/parseEnterRoomRequest.go
package chatService

import (
	"encoding/base64"
	"myapp/api/models"
	"myapp/api/service"
	"strings"

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

	return strings.EqualFold(upgradeHeader, "websocket") &&
		strings.EqualFold(connectionHeader, "upgrade") &&
		strings.EqualFold(secWebSocketVersionHeader, "13") &&
		isSecWebSocketKeyValid
}

func ParseEnterChatRequest(c *gin.Context) (models.LoginInfo, error) {
	loginInfo, err := service.ParseLoginInfo(c)
	if err != nil {
		return models.LoginInfo{}, err
	}
	return loginInfo, nil
}
