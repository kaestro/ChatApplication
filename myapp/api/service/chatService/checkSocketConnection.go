// myapp/api/service/chatService/checkSocketConnection.go
package chatService

import (
	"myapp/internal/chat"
)

func CheckSocketConnection(sessionID string) error {
	cm := chat.GetChatManager()
	_, err := cm.GetClient(sessionID)

	return err
}
