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

func ParseAndAuthenticateRequest(c *gin.Context) (models.RoomRequest, error) {
	req, err := ParseEnterRoomRequest(c)
	if err != nil {
		return models.RoomRequest{}, err
	}

	loginInfo := models.NewLoginInfo(req.EmailAddress, req.Password)
	_, err = userService.NewUserServiceUtil().AuthenticateUser(loginInfo, req.LoginSessionID)
	if err != nil {
		return models.RoomRequest{}, err
	}

	return req, nil
}

func EnterChatRoom(c *gin.Context, req models.RoomRequest) error {
	cm := chat.GetChatManager()
	cm.ProvideClientToUser(c.Writer, c.Request, req.LoginSessionID)
	cm.ClientEnterRoom(req.RoomName, req.LoginSessionID)

	return nil
}
