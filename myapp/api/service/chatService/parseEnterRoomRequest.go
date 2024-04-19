// myapp/api/service/chatService/parseEnterRoomRequest.go
package chatService

import (
	"encoding/base64"
	"errors"
	"myapp/api/models"
	"myapp/jsonProperties"
	"myapp/types"
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

func ParseEnterChatRequest(c *gin.Context) (models.LoginSessionInfo, error) {
	sessionID := c.GetHeader(jsonProperties.SessionKey)
	emailAddress := c.GetHeader(jsonProperties.EmailAddress)

	if sessionID == "" || emailAddress == "" {
		return models.LoginSessionInfo{}, errors.New("invalid request")
	}

	loginSessionInfo := models.NewLoginSessionInfo(emailAddress, types.LoginSessionID(sessionID))

	return loginSessionInfo, nil
}
