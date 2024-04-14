package chatHandler

import (
	"myapp/api/models"
	"myapp/api/service/chatService"
	"myapp/api/service/generalService"
	"myapp/internal/chat"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Request Type: POST
// Headers: Session-Key
// Body: RoomRequest { roomName, emailAddress, password }, ChatMessage { roomName, userName, content }
func SendMessage(c *gin.Context) {
	roomRequest, chatMessage, isProperRequest := parseRequestData(c)
	if !isProperRequest {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userSessionID, err := handleUserSession(c, &roomRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !chatService.IsUserInsideRoom(roomRequest, userSessionID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User is not inside the room"})
		return
	}

	err = RequestSendMessageToClientManager(userSessionID, chatMessage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Entered room successfully"})
}

func RequestSendMessageToClientManager(userSessionID string, chatMessage chat.ChatMessage) error {
	cm := chat.GetChatManager()
	err := cm.SendMessageToRoom(userSessionID, chatMessage)

	return err
}

func parseRequestData(c *gin.Context) (models.RoomRequest, chat.ChatMessage, bool) {
	roomRequest, err := generalService.ParseRoomRequest(c)
	if err != nil {
		return models.RoomRequest{}, chat.ChatMessage{}, true
	}

	chatMessage, err := generalService.ParseChatMessage(c)
	if err != nil {
		return models.RoomRequest{}, chat.ChatMessage{}, false
	}
	return roomRequest, chatMessage, true
}
