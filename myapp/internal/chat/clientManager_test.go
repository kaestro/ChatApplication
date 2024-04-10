// myapp/internal/chat/clientManager_test.go
package chat

import (
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
	gotClient := cm.getClient(sampleLoginSessionID)
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
		client := &Client{loginSessionID: string(rune(i))}
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
	if cm.getClient(sampleUpdateID) != sampleClient {
		t.Errorf("UpdateClient failed, expected client to be %v, got %v", sampleClient, cm.getClient(sampleUpdateID))
		return
	}

	t.Log("TestClientManagerUpdateClient passed")
}

func TestClientManagerCreateClient(t *testing.T) {
	cm := getClientManager()

	client := cm.createClient(sampleLoginSessionID, &MockConn{})
	if client == nil {
		t.Errorf("CreateClient failed, expected client to be created")
		return
	}

	if client.GetLoginSessionID() != sampleLoginSessionID {
		t.Errorf("CreateClient failed, expected client to have sessionID %s, got %s", sampleLoginSessionID, client.GetLoginSessionID())
		return
	}

	t.Log("TestClientManagerCreateClient passed")
}

func TestClientManagerRegisterNewClient(t *testing.T) {
	cm := getClientManager()

	client := cm.registerNewClient(sampleLoginSessionID, &MockConn{})
	if client == nil {
		t.Errorf("RegisterNewClient failed, expected client to be created")
		return
	}

	if !cm.isClientRegistered(client.GetLoginSessionID()) {
		t.Errorf("RegisterNewClient failed, expected client to be registered")
		return
	}

	t.Log("TestClientManagerRegisterNewClient passed")
}
