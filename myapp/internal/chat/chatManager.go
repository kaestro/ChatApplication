// myapp/internal/chat/chatManager.go
package chat

import (
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	readBufferSize  = 1024
	writeBufferSize = 1024
)

// TODO: Design ChatManager
// 채팅 서버의 전반적인 관리를 담당하는 ChatManager
// 채팅과 관련한 모든 요청은 ChatManager를 통해 이루어진다.
type ChatManager struct {
}

func (cm *ChatManager) provideClientConnection(w http.ResponseWriter, r *http.Request, loginSessionID string) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  readBufferSize,
		WriteBufferSize: writeBufferSize,
	}

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		return
	}

	clientManager = getClientManager()
	clientManager.registerNewClient(loginSessionID, conn)
}

func (cm *ChatManager) provideChatRoomList(w http.ResponseWriter, r *http.Request) {
}

func (cm *ChatManager) provideClientList(w http.ResponseWriter, r *http.Request) {
}
