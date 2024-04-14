// myapp/api/handlers/chatHandler/getRoomList.go
package chatHandler

import (
	"myapp/api/service/generalService"
	"myapp/api/service/userService"
	"myapp/internal/chat"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Request Type: GET
// Headers: Session-Key
// Request Body: LoginInfo (emailaddress, password)
func GetRoomList(ginContext *gin.Context) {
	loginInfo, err := generalService.ParseLoginInfo(ginContext)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userSessionKey := generalService.GetSessionKeyFromHeader(ginContext)

	userServiceUtil := userService.NewUserServiceUtil()
	_, err = userServiceUtil.AuthenticateUser(loginInfo, userSessionKey)
	if err != nil {
		ginContext.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	cm := chat.GetChatManager()
	roomNames := cm.GetAllRoomNames()

	ginContext.JSON(http.StatusOK, gin.H{"roomNames": roomNames})
}
