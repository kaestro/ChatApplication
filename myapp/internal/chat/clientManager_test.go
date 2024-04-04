// myapp/internal/chat/clientManager_test.go
package chat

import (
	"testing"
)

var (
	maxClients      = 10000
	sampleSessionID = "123"
	sampleClient    = &Client{sessionID: sampleSessionID}
)

func TestClientManager(t *testing.T) {
	cm := GetClientManager()

	// Test AddClient
	cm.AddClient(sampleClient)
	if !cm.CheckClient(sampleSessionID) {
		t.Errorf("AddClient failed, expected sessionID 123 to exist")
		return
	}

	// Test GetClient
	gotClient := cm.GetClient(sampleSessionID)
	if gotClient != sampleClient {
		t.Errorf("GetClient failed, expected %v, got %v", sampleClient, gotClient)
		return
	}

	// Test RemoveClient
	cm.RemoveClient(sampleSessionID)
	if cm.CheckClient(sampleSessionID) {
		t.Errorf("RemoveClient failed, expected sessionID %s to be removed", sampleSessionID)
		return
	}

	t.Log("TestClientManager passed")
}

func TestClientManagerCapacity(t *testing.T) {
	cm := GetClientManager()

	// Test AddClient
	for i := 0; i < maxClients; i++ {
		client := &Client{sessionID: string(rune(i))}
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
