// myapp/api/service/chatService/publishWebsocket.go
package chatService

import (
	"myapp/internal/chat"
	"net/http"
)

func PublishWebSocket(w http.ResponseWriter, r *http.Request, sessionKey string) error {
	chatManager := chat.GetChatManager()

	err := chatManager.ProvideClientToUser(w, r, sessionKey)
	return err
}
