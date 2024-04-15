// myapp/api/service/chatService/chatService.go
package chatService

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"myapp/api/models"
	"myapp/internal/chat"

	"github.com/gin-gonic/gin"
)

func ValidateUpgradeHeader(c *gin.Context) error {
	if !IsHandshakeAndKeyHeadersValid(c) {
		return errors.New("invalid Upgrade header")
	}
	return nil
}

func ParseEnterLoginSessionInfo(c *gin.Context) (models.LoginSessionInfo, error) {
	loginSessionInfo, err := ParseEnterChatRequest(c)
	if err != nil {
		return models.LoginSessionInfo{}, err
	}

	return loginSessionInfo, nil
}

func EnterChat(c *gin.Context, req models.LoginSessionInfo) error {
	cm := chat.GetChatManager()
	err := cm.ProvideClientToUser(c.Writer, c.Request, req.LoginSessionID)

	return err
}

func EnterChatRoom(c *gin.Context, req models.RoomRequest) error {
	cm := chat.GetChatManager()
	err := cm.ClientEnterRoom(req.RoomName, req.LoginSessionID)

	return err
}

func GenerateRandomSocketKey() (string, error) {
	key := make([]byte, 16)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	secWebSocketKey := base64.StdEncoding.EncodeToString(key)
	return secWebSocketKey, nil
}
