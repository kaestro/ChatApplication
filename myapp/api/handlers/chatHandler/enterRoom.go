package chatHandler

import (
	"myapp/api/models"
	"myapp/api/service/generalService"
	"myapp/api/service/userService"
	"myapp/internal/chat"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Request Type: POST
// Headers: Session-Key
// Body: RoomRequest { roomName, emailAddress, password }
func EnterRoom(c *gin.Context) {
	roomRequest, err := generalService.ParseRoomRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userSessionKey, err := handleUserSession(c, &roomRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = handleRoomEntry(userSessionKey, roomRequest.RoomName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Entered room successfully"})
}

func handleUserSession(c *gin.Context, roomRequest *models.RoomRequest) (string, error) {
	loginInfo := roomRequest.GetLoginInfo()
	userSessionKey := generalService.GetSessionKeyFromHeader(c)
	userServiceUtil := userService.NewUserServiceUtil()

	userSessionKey, isLoggedIn := userServiceUtil.CheckUserLoggedIn(userSessionKey, loginInfo)
	if !isLoggedIn {
		user, err := userServiceUtil.AuthenticateUserByLoginInfo(loginInfo, userSessionKey)
		if err != nil {
			return "", err
		}

		userSessionKey, err = userServiceUtil.GenerateSessionKey(user)
		if err != nil {
			return "", err
		}
	}

	return userSessionKey, nil
}

func handleRoomEntry(userSessionKey string, roomName string) error {
	cm := chat.GetChatManager()
	err := cm.ClientEnterRoom(roomName, userSessionKey)
	if err != nil {
		return err
	}

	return nil
}
