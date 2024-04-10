// myapp/internal/chat/clientManager_test.go
package chat

import (
	"testing"
)

var (
	maxClients           = 10000
	sampleLoginSessionID = "123"
	sampleUpdateID       = "456"
	sampleClient         = NewClient(sampleLoginSessionID)
)

func TestClientManager(t *testing.T) {
	cm := GetClientManager()

	// Test AddClient
	cm.AddClient(sampleClient)
	if !cm.CheckClient(sampleLoginSessionID) {
		t.Errorf("AddClient failed, expected sessionID 123 to exist")
		return
	}

	// Test GetClient
	gotClient := cm.GetClient(sampleLoginSessionID)
	if gotClient != sampleClient {
		t.Errorf("GetClient failed, expected %v, got %v", sampleClient, gotClient)
		return
	}

	// Test RemoveClient
	cm.RemoveClient(sampleLoginSessionID)
	if cm.CheckClient(sampleLoginSessionID) {
		t.Errorf("RemoveClient failed, expected sessionID %s to be removed", sampleLoginSessionID)
		return
	}

	t.Log("TestClientManager passed")
}

func TestClientManagerCapacity(t *testing.T) {
	cm := GetClientManager()

	// Test AddClient
	for i := 0; i < maxClients; i++ {
		client := &Client{loginSessionID: string(rune(i))}
		cm.AddClient(client)
	}

	// Test AddClient exceeding capacity
	for i := 0; i < maxClients; i++ {
		if !cm.CheckClient(string(rune(i))) {
			t.Errorf("AddClient failed, expected sessionID %d to not exist", i)
			return
		}
	}

	t.Log("TestClientManagerCapacity passed")
}

func TestClientManagerUpdateClientID(t *testing.T) {
	cm := GetClientManager()

	cm.AddClient(sampleClient)
	if !cm.CheckClient(sampleLoginSessionID) {
		t.Errorf("AddClient failed, expected sessionID %s to exist", sampleLoginSessionID)
		return
	}

	cm.UpdateClientID(sampleClient, sampleUpdateID)

	// Check if the old sessionID no longer exists
	if cm.CheckClient(sampleLoginSessionID) {
		t.Errorf("UpdateClient failed, expected old sessionID %s to be removed", sampleLoginSessionID)
		return
	}

	// Check if the new sessionID exists
	if !cm.CheckClient(sampleUpdateID) {
		t.Errorf("UpdateClient failed, expected new sessionID %s to exist", sampleUpdateID)
		return
	}

	// Check if the client associated with the new sessionID is the same as the sampleClient
	if cm.GetClient(sampleUpdateID) != sampleClient {
		t.Errorf("UpdateClient failed, expected client to be %v, got %v", sampleClient, cm.GetClient(sampleUpdateID))
		return
	}

	t.Log("TestClientManagerUpdateClient passed")
}
