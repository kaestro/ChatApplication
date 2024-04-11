// myapp/internal/chat/clientManager_test.go
package chat

import (
	"crypto/rand"
	"encoding/hex"
	"testing"
)

func TestClientManager(t *testing.T) {
	cm := getClientManager()

	// Test AddClient
	cm.registerClient(sampleClient)
	if !cm.isClientRegistered(sampleLoginSessionID) {
		t.Errorf("AddClient failed, expected sessionID 123 to exist")
		return
	}

	// Test GetClient
	gotClient := cm.getClientByLoginSessionID(sampleLoginSessionID)
	if gotClient != sampleClient {
		t.Errorf("GetClient failed, expected %v, got %v", sampleClient, gotClient)
		return
	}

	// Test RemoveClient
	cm.unRegisterClient(sampleLoginSessionID)
	if cm.isClientRegistered(sampleLoginSessionID) {
		t.Errorf("RemoveClient failed, expected sessionID %s to be removed", sampleLoginSessionID)
		return
	}

	t.Log("TestClientManager passed")
}

func TestClientManagerCapacity(t *testing.T) {
	cm := getClientManager()

	// Test AddClient
	for i := 0; i < maxClients; i++ {
		client := &client{loginSessionID: string(rune(i))}
		cm.registerClient(client)
	}

	// Test AddClient exceeding capacity
	for i := 0; i < maxClients; i++ {
		if !cm.isClientRegistered(string(rune(i))) {
			t.Errorf("AddClient failed, expected sessionID %d to not exist", i)
			return
		}
	}

	t.Log("TestClientManagerCapacity passed")
}

func TestClientManagerUpdateClientID(t *testing.T) {
	cm := getClientManager()

	cm.registerClient(sampleClient)
	if !cm.isClientRegistered(sampleLoginSessionID) {
		t.Errorf("AddClient failed, expected sessionID %s to exist", sampleLoginSessionID)
		return
	}

	cm.updateClientID(sampleClient, sampleUpdateID)

	// Check if the old sessionID no longer exists
	if cm.isClientRegistered(sampleLoginSessionID) {
		t.Errorf("UpdateClient failed, expected old sessionID %s to be removed", sampleLoginSessionID)
		return
	}

	// Check if the new sessionID exists
	if !cm.isClientRegistered(sampleUpdateID) {
		t.Errorf("UpdateClient failed, expected new sessionID %s to exist", sampleUpdateID)
		return
	}

	// Check if the client associated with the new sessionID is the same as the sampleClient
	if cm.getClientByLoginSessionID(sampleUpdateID) != sampleClient {
		t.Errorf("UpdateClient failed, expected client to be %v, got %v", sampleClient, cm.getClientByLoginSessionID(sampleUpdateID))
		return
	}

	t.Log("TestClientManagerUpdateClient passed")
}

func TestClientManagerCreateClient(t *testing.T) {
	cm := getClientManager()

	client := cm.createNewClient(sampleLoginSessionID, &mockConn{})
	if client == nil {
		t.Errorf("CreateClient failed, expected client to be created")
		return
	}

	if client.getLoginSessionID() != sampleLoginSessionID {
		t.Errorf("CreateClient failed, expected client to have sessionID %s, got %s", sampleLoginSessionID, client.getLoginSessionID())
		return
	}

	t.Log("TestClientManagerCreateClient passed")
}

func TestClientManagerRegisterNewClient(t *testing.T) {
	cm := getClientManager()

	client, _ := cm.registerNewClient(sampleLoginSessionID, &mockConn{})
	if client == nil {
		t.Errorf("RegisterNewClient failed, expected client to be created")
		return
	}

	if !cm.isClientRegistered(client.getLoginSessionID()) {
		t.Errorf("RegisterNewClient failed, expected client to be registered")
		return
	}

	t.Log("TestClientManagerRegisterNewClient passed")
}

func TestClientManagerEmptyClientManager(t *testing.T) {
	cm := getClientManager()
	cm.emptyClientManager()

	if cm.getClientCount() != 0 {
		t.Errorf("EmptyClientManager failed, expected client count to be 0, got %d", cm.getClientCount())
		return
	}

	t.Log("TestClientManagerEmptyClientManager passed")
}

func TestClientManagerGetClientCount(t *testing.T) {
	cm := getClientManager()
	cm.emptyClientManager()

	for i := 0; i < maxClients; i++ {
		cm.registerNewClient(generateUniqueString(), &mockConn{})
	}

	if cm.getClientCount() != maxClients {
		t.Errorf("GetClientCount failed, expected %d, got %d", maxClients, cm.getClientCount())
		return
	}

	t.Log("TestClientManagerGetClientCount passed")
}

func generateUniqueString() string {
	b := make([]byte, 16) // adjust size for your needs
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}
