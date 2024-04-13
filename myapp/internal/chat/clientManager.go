// myapp/internal/chat/clientManager.go
package chat

import (
	"fmt"
	"sync"
)

var (
	clientOnce                sync.Once
	cmInstance                *clientManager
	ErrorClientExists         = "client already exists"
	ErrorFailedToCreateClient = "failed to create client"
)

// Client에 대한 CRUD를 담당하는 clientManager
// Singleton 객체로 구현되어 있다.
type clientManager struct {
	clients map[string]*client // key: loginSessionID, value: Client object
}

func getClientManager() *clientManager {
	clientOnce.Do(func() {
		cmInstance = &clientManager{
			clients: make(map[string]*client),
		}
	})

	return cmInstance
}

func (cm *clientManager) isClientRegistered(loginSessionID string) bool {
	_, ok := cm.clients[loginSessionID]

	return ok
}

func (cm *clientManager) getClientByLoginSessionID(loginSessionID string) *client {
	if !cm.isClientRegistered(loginSessionID) {
		fmt.Println("Client with sessionID", loginSessionID, "does not exist")
		return nil
	}

	return cm.clients[loginSessionID]
}

func (cm *clientManager) registerClient(client *client) {
	if cm.isClientRegistered(client.loginSessionID) {
		fmt.Println("Client with sessionID", client.loginSessionID, "already exists")
		return
	}

	cm.clients[client.loginSessionID] = client
}

func (cm *clientManager) unRegisterClient(sessionID string) {
	if !cm.isClientRegistered(sessionID) {
		fmt.Println("Client with sessionID", sessionID, "does not exist")
		return
	}

	delete(cm.clients, sessionID)
}

func (cm *clientManager) updateClientID(client *client, loginSessionID string) {
	for savedID, savedClient := range cm.clients {
		if savedClient == client {
			cm.clients[loginSessionID] = client
			delete(cm.clients, savedID)
			break
		}
	}
}

func (cm *clientManager) createNewClient(loginSessionID string, conn Conn) *client {
	if cm.isClientRegistered(loginSessionID) {
		fmt.Println("Client with sessionID", loginSessionID, "already exists")
		return cm.getClientByLoginSessionID(loginSessionID)
	}

	client := newClient(loginSessionID, conn)
	return client
}

func (cm *clientManager) registerNewClient(loginSessionID string, conn Conn) (*client, error) {
	if cm.isClientRegistered(loginSessionID) {
		return cm.clients[loginSessionID], error(fmt.Errorf(ErrorClientExists))
	}

	client := cm.createNewClient(loginSessionID, conn)
	if client == nil {
		return nil, error(fmt.Errorf(ErrorFailedToCreateClient))
	}

	cm.registerClient(client)

	return client, nil
}

func (cm *clientManager) getClientCount() int {
	return len(cm.clients)
}

func (cm *clientManager) clearClientManager() {
	cm.clients = make(map[string]*client)
}
