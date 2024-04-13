// myapp/api/service/chatService/publishAndCheckConnection.go
package chatService

import (
	"myapp/internal/chat"
	"net/http"
)

func CheckSocketConnection(sessionID string) error {
	cm := chat.GetChatManager()
	_, err := cm.GetClient(sessionID)

	return err
}

func PublishWebSocket(w http.ResponseWriter, r *http.Request, sessionKey string) error {
	chatManager := chat.GetChatManager()

	err := chatManager.ProvideClientToUser(w, r, sessionKey)
	return err
}
