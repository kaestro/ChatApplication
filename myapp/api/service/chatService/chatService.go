// myapp/api/service/chatService/chatService.go
package chatService

import (
	"errors"
	"myapp/api/models"
	"myapp/api/service/userService"
	"myapp/internal/chat"

	"github.com/gin-gonic/gin"
)

func ValidateUpgradeHeader(c *gin.Context) error {
	if !IsUpgradeHeaderValid(c) {
		return errors.New("invalid Upgrade header")
	}
	return nil
}

func ParseChatRequestAndAuthenticateUser(c *gin.Context) (models.LoginInfo, error) {
	loginInfo, err := ParseEnterChatRequest(c)
	if err != nil {
		return models.LoginInfo{}, err
	}

	_, err = userService.NewUserServiceUtil().AuthenticateUser(loginInfo, loginInfo.LoginSessionID)
	if err != nil {
		return models.LoginInfo{}, err
	}

	return loginInfo, nil
}

func EnterChat(c *gin.Context, req models.LoginInfo) error {
	cm := chat.GetChatManager()
	err := cm.ProvideClientToUser(c.Writer, c.Request, req.LoginSessionID)

	return err
}

func EnterChatRoom(c *gin.Context, req models.RoomRequest) error {
	cm := chat.GetChatManager()
	err := cm.ClientEnterRoom(req.RoomName, req.LoginSessionID)

	return err
}
