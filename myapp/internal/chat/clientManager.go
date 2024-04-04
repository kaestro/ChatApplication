// myapp/internal/chat/clientManager.go
package chat

import (
	"fmt"
	"sync"
)

var (
	once          sync.Once
	clientManager *ClientManager
)

// key: sessionID, value: Client object
// Question: How can I make sure that ClientManager won't be calling
// garbage collection on the Client object?
// Or should I assure it?
// Should I limit the number of clients?
// How does garbage collection work in Go?
// It would be taking care of the memory management, which I'm not sure of
// Question: Is making ClientManager a singleton a good idea?
type ClientManager struct {
	clients map[string]*Client
}

func GetClientManager() *ClientManager {
	once.Do(func() {
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

func (cm *ClientManager) GetClient(sessionID string) *Client {
	if !cm.CheckClient(sessionID) {
		fmt.Println("Client with sessionID", sessionID, "does not exist")
		return nil
	}

	return cm.clients[sessionID]
}

func (cm *ClientManager) AddClient(client *Client) {
	if cm.CheckClient(client.sessionID) {
		fmt.Println("Client with sessionID", client.sessionID, "already exists")
		return
	}
	cm.clients[client.sessionID] = client
}

func (cm *ClientManager) RemoveClient(sessionID string) {
	if !cm.CheckClient(sessionID) {
		fmt.Println("Client with sessionID", sessionID, "does not exist")
		return
	}
	delete(cm.clients, sessionID)
}
