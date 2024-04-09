// myapp/internal/chat/clientManager.go
package chat

import (
	"fmt"
	"sync"
)

var (
	clientOnce    sync.Once
	clientManager *ClientManager
)

// 모든 클라이언트를 관리하는 singleton 객체
// TODO: 갯수 제한 및 지속 시간 제한을 둘 수 있도록 변경
type ClientManager struct {
	clients map[string]*Client // key: sessionID, value: Client object
}

func GetClientManager() *ClientManager {
	clientOnce.Do(func() {
		clientManager = &ClientManager{
			clients: make(map[string]*Client),
		}
	})

	return clientManager
}

func (cm *ClientManager) CheckClient(sessionID string) bool {
	_, ok := cm.clients[sessionID]
	return ok
}

// TODO: fmt 대신 별개의 로거를 사용하도록 변경
func (cm *ClientManager) GetClient(sessionID string) *Client {
	if !cm.CheckClient(sessionID) {
		fmt.Println("Client with sessionID", sessionID, "does not exist")
		return nil
	}

	return cm.clients[sessionID]
}

func (cm *ClientManager) AddClient(client *Client) {
	if cm.CheckClient(client.loginSessionID) {
		fmt.Println("Client with sessionID", client.loginSessionID, "already exists")
		return
	}
	cm.clients[client.loginSessionID] = client
}

func (cm *ClientManager) RemoveClient(sessionID string) {
	if !cm.CheckClient(sessionID) {
		fmt.Println("Client with sessionID", sessionID, "does not exist")
		return
	}
	delete(cm.clients, sessionID)
}

func (cm *ClientManager) UpdateClientID(client *Client, loginSessionID string) {
	for savedID, savedClient := range cm.clients {
		if savedClient == client {
			cm.clients[loginSessionID] = client
			delete(cm.clients, savedID)
			break
		}
	}
}
