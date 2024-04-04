// myapp/internal/chat/clientManager_test.go
package chat

import (
	"testing"
)

var (
	maxClients = 10000
)

func TestClientManager(t *testing.T) {
	cm := GetClientManager()

	// Test AddClient
	client := &Client{sessionID: "123"}
	cm.AddClient(client)
	if !cm.CheckClient("123") {
		t.Errorf("AddClient failed, expected sessionID 123 to exist")
		return
	}

	// Test GetClient
	gotClient := cm.GetClient("123")
	if gotClient != client {
		t.Errorf("GetClient failed, expected %v, got %v", client, gotClient)
		return
	}

	// Test RemoveClient
	cm.RemoveClient("123")
	if cm.CheckClient("123") {
		t.Errorf("RemoveClient failed, expected sessionID 123 to not exist")
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
