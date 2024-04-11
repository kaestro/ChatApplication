// myapp/internal/chat/chatManager.go
package chat

import (
	"errors"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

const (
	readBufferSize  = 1024
	writeBufferSize = 1024
)

var (
	chatManagerOnce sync.Once
	chatManager     *ChatManager
)

// 채팅 서버의 전반적인 관리를 담당하는 ChatManager
// 채팅과 관련한 모든 요청은 ChatManager를 통해 이루어진다.
// Singleton 객체로 구현되어 있다.
type ChatManager struct {
	upgrader      websocket.Upgrader
	clientManager *ClientManager
	roomManager   *RoomManager
}

func NewChatManager() *ChatManager {
	clientManager = getClientManager()
	roomManager = GetRoomManager()

	chatManagerOnce.Do(func() {
		clientManager = getClientManager()
		roomManager = GetRoomManager()

		chatManager = &ChatManager{
			upgrader: websocket.Upgrader{
				ReadBufferSize:  readBufferSize,
				WriteBufferSize: writeBufferSize,
			},
			clientManager: clientManager,
			roomManager:   roomManager,
		}
	})

	return chatManager
}

func (cm *ChatManager) ProvideClientToUser(w http.ResponseWriter, r *http.Request, loginSessionID string) error {
	conn, err := cm.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	clientManager = getClientManager()
	client, err := clientManager.registerNewClient(loginSessionID, conn)
	if err != nil {
		return err
	}

	if client == nil {
		return errors.New("failed to register new client")
	}

	return nil
}

func (cm *ChatManager) RemoveClientFromUser(loginSessionID string) {
	clientManager = getClientManager()
	clientManager.unRegisterClient(loginSessionID)
}
