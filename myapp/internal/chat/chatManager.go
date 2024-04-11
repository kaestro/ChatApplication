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
	upgrader websocket.Upgrader
}

func NewChatManager() *ChatManager {
	chatManagerOnce.Do(func() {
		chatManager = &ChatManager{
			upgrader: websocket.Upgrader{
				ReadBufferSize:  readBufferSize,
				WriteBufferSize: writeBufferSize,
			},
		}
	})

	return chatManager
}

func (cm *ChatManager) ProvideClientToUser(w http.ResponseWriter, r *http.Request, loginSessionID string) error {
	conn, err := cm.upgradeToWebsocket(w, r)
	if err != nil {
		return err
	}

	err = cm.registerNewClient(loginSessionID, conn)
	return err
}

func (cm *ChatManager) upgradeToWebsocket(w http.ResponseWriter, r *http.Request) (Conn, error) {
	conn, err := cm.upgrader.Upgrade(w, r, nil)
	return conn, err
}

func (cm *ChatManager) registerNewClient(loginSessionID string, conn Conn) error {
	cmInstance = getClientManager()
	_, err := cmInstance.registerNewClient(loginSessionID, conn)

	return err
}

func (cm *ChatManager) RemoveClientFromUser(loginSessionID string) {
	cmInstance = getClientManager()
	cmInstance.unRegisterClient(loginSessionID)
}

func (cm *ChatManager) CreateRoom(roomName string) error {
	rmInstance = getRoomManager()
	room := rmInstance.createNewRoom(roomName)

	if room == nil {
		return errors.New("failed to create new room")
	}

	return nil
}

func (cm *ChatManager) RemoveRoom(roomName string) error {
	rmInstance = getRoomManager()
	return rmInstance.removeRoom(roomName)
}

func (cm *ChatManager) ClientEnterRoom(roomName, loginSessionID string) error {
	room, err := cm.getRoom(roomName)
	if err != nil {
		return err
	}

	client, err := cm.getClient(loginSessionID)
	if err != nil {
		return err
	}

	err = room.addClient(client)
	return err
}

func (cm *ChatManager) getClient(loginSessionID string) (*client, error) {
	cmInstance = getClientManager()
	client := cmInstance.getClientByLoginSessionID(loginSessionID)
	if client == nil {
		return nil, errors.New("user of loginSessionID " + loginSessionID + " does not exist")
	}
	return client, nil
}

func (cm *ChatManager) getRoom(roomName string) (*room, error) {
	rmInstance = getRoomManager()
	room := rmInstance.getRoom(roomName)
	if room == nil {
		return nil, errors.New("room of roomName " + roomName + " does not exist")
	}
	return room, nil
}
