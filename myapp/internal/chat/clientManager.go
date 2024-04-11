// myapp/internal/chat/clientManager.go
package chat

import (
	"errors"
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

func getClientManager() *ClientManager {
	clientOnce.Do(func() {
		clientManager = &ClientManager{
			clients: make(map[string]*Client),
		}
	})

	return clientManager
}

func (cm *ClientManager) isClientRegistered(loginSessionID string) bool {
	_, ok := cm.clients[loginSessionID]

	return ok
}

func (cm *ClientManager) getClientByLoginSessionID(loginSessionID string) *Client {
	if !cm.isClientRegistered(loginSessionID) {
		fmt.Println("Client with sessionID", loginSessionID, "does not exist")
		return nil
	}

	return cm.clients[loginSessionID]
}

func (cm *ClientManager) registerClient(client *Client) {
	if cm.isClientRegistered(client.loginSessionID) {
		fmt.Println("Client with sessionID", client.loginSessionID, "already exists")
		return
	}

	cm.clients[client.loginSessionID] = client
}

func (cm *ClientManager) unRegisterClient(sessionID string) {
	if !cm.isClientRegistered(sessionID) {
		fmt.Println("Client with sessionID", sessionID, "does not exist")
		return
	}

	delete(cm.clients, sessionID)
}

func (cm *ClientManager) updateClientID(client *Client, loginSessionID string) {
	for savedID, savedClient := range cm.clients {
		if savedClient == client {
			cm.clients[loginSessionID] = client
			delete(cm.clients, savedID)
			break
		}
	}
}

func (cm *ClientManager) createNewClient(loginSessionID string, conn Conn) *Client {
	if cm.isClientRegistered(loginSessionID) {
		fmt.Println("Client with sessionID", loginSessionID, "already exists")
		return nil
	}

	client := NewClient(loginSessionID, conn)
	return client
}

func (cm *ClientManager) registerNewClient(loginSessionID string, conn Conn) (*Client, error) {
	if cm.isClientRegistered(loginSessionID) {
		fmt.Println("Client with sessionID", loginSessionID, "already exists")
		return cm.clients[loginSessionID], errors.New("client already exists")
	}

	client := cm.createNewClient(loginSessionID, conn)
	if client == nil {
		fmt.Println("Failed to create client with sessionID", loginSessionID)
		return nil, errors.New("failed to create client")
	}

	cm.registerClient(client)

	return client, nil
}

func (cm *ClientManager) getClientCount() int {
	return len(cm.clients)
}

func (cm *ClientManager) emptyClientManager() {
	cm.clients = make(map[string]*Client)
}
