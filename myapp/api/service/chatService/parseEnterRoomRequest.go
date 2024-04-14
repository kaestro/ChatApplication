// myapp/api/service/chatService/parseEnterRoomRequest.go
package chatService

import (
	"encoding/base64"
	"myapp/api/models"
	"myapp/api/service"
	"strings"

	"github.com/gin-gonic/gin"
)

func IsHandshakeAndKeyHeadersValid(c *gin.Context) bool {
	return IsHandshakeHeadersValid(c) && IsSecWebsocketKeyValid(c)
}

func IsSecWebsocketKeyValid(c *gin.Context) bool {
	secWebSocketKeyHeader := c.GetHeader("Sec-WebSocket-Key")
	decodedKey, err := base64.StdEncoding.DecodeString(secWebSocketKeyHeader)

	result := (err == nil && len(decodedKey) == 16)
	return result
}

func IsHandshakeHeadersValid(c *gin.Context) bool {
	headers := map[string]string{
		"Upgrade":               "websocket",
		"Connection":            "upgrade",
		"Sec-WebSocket-Version": "13",
	}

	for header, expectedValue := range headers {
		if !strings.EqualFold(c.GetHeader(header), expectedValue) {
			return false
		}
	}
	return true
}

func ParseEnterChatRequest(c *gin.Context) (models.LoginInfo, error) {
	loginInfo, err := service.ParseLoginInfo(c)
	if err != nil {
		return models.LoginInfo{}, err
	}
	return loginInfo, nil
}
