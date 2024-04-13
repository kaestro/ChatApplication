// myapp/api/handlers/chatHandler/createRoom.go
package chatHandler

import (
	"myapp/api/models"
	"myapp/api/service"
	"myapp/internal/chat"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Request Type: POST
// Headers: Session-Key
// Body: RoomRequest { roomName, emailAddress, password }
func CreateRoom(ginContext *gin.Context) {
	roomRequest, err := service.ParseRoomRequest(ginContext)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = handleUserSession(ginContext, &roomRequest)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !handleRoomCreation(roomRequest.RoomName) {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create room"})
		return
	}

	if !handleUserRoomEntry(&roomRequest) {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Creator failed to enter room"})
		return
	}

	ginContext.JSON(http.StatusOK, gin.H{"message": "Room created successfully"})
}

func handleRoomCreation(roomName string) bool {
	cm := chat.GetChatManager()
	err := cm.CreateRoom(roomName)

	return err == nil
}

func handleUserRoomEntry(roomRequest *models.RoomRequest) bool {
	cm := chat.GetChatManager()
	err := cm.ClientEnterRoom(roomRequest.RoomName, roomRequest.LoginSessionID)

	return err == nil
}
