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
	clientManager *clientManager
	roomManager   *roomManager
}

func NewChatManager() *ChatManager {
	cmInstance = getClientManager()
	rmInstance = getRoomManager()

	chatManagerOnce.Do(func() {
		cmInstance = getClientManager()
		rmInstance = getRoomManager()

		chatManager = &ChatManager{
			upgrader: websocket.Upgrader{
				ReadBufferSize:  readBufferSize,
				WriteBufferSize: writeBufferSize,
			},
			clientManager: cmInstance,
			roomManager:   rmInstance,
		}
	})

	return chatManager
}

func (cm *ChatManager) ProvideClientToUser(w http.ResponseWriter, r *http.Request, loginSessionID string) error {
	conn, err := cm.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	cmInstance = getClientManager()
	client, err := cmInstance.registerNewClient(loginSessionID, conn)
	if err != nil {
		return err
	}

	if client == nil {
		return errors.New("failed to register new client")
	}

	return nil
}

func (cm *ChatManager) RemoveClientFromUser(loginSessionID string) {
	cmInstance = getClientManager()
	cmInstance.unRegisterClient(loginSessionID)
}
