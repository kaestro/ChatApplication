// myapp/api/service/chatService/chatServiceUtil.go
package chatService

import (
	"myapp/api/models"
	"myapp/internal/chat"
	"myapp/types"
)

func IsUserInsideRoom(roomRequest models.RoomRequest, userSessionKey string) bool {
	cm := chat.GetChatManager()
	return cm.IsClientInsideRoom(roomRequest.RoomName, types.LoginSessionID(userSessionKey))
}
