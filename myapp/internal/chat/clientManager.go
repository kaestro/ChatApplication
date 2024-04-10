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

// Client에 대한 CRUD를 담당하는 ClientManager
// Singleton 객체로 구현되어 있다.
type ClientManager struct {
	clients map[string]*Client // key: loginSessionID, value: Client object
}

func GetClientManager() *ClientManager {
	clientOnce.Do(func() {
		clientManager = &ClientManager{
			clients: make(map[string]*Client),
		}
	})

	return clientManager
}

func (cm *ClientManager) CheckClient(loginSessionID string) bool {
	_, ok := cm.clients[loginSessionID]
	return ok
}

// TODO: fmt 대신 별개의 로거를 사용하도록 변경
func (cm *ClientManager) GetClient(loginSessionID string) *Client {
	if !cm.CheckClient(loginSessionID) {
		fmt.Println("Client with sessionID", loginSessionID, "does not exist")
		return nil
	}

	return cm.clients[loginSessionID]
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
