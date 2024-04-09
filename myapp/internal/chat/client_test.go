// myapp/internal/chat/client_test.go
package chat

import (
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

const (
	ExpectedClientSessionLength = 1
	testMessage                 = "test message"
)

func TestIsSameClient(t *testing.T) {
	client := NewClient(sampleSessionID)

	// Test isSameClient with the same session ID
	if !client.isSameClient(sampleSessionID) {
		t.Errorf("isSameClient failed, expected true, got false")
		return
	}

	// Test isSameClient with a different session ID
	if client.isSameClient("differentSessionID") {
		t.Errorf("isSameClient failed, expected false, got true")
		return
	}

	t.Logf("isSameClient passed")
}

func TestClientAddClientSession(t *testing.T) {
	client := NewClient(sampleSessionID)

	// Test AddClientSession
	socketConn := &websocket.Conn{}
	room := NewRoom(sampleRoomID)
	client.AddClientSession(socketConn, room, sampleSessionID)

	if len(client.clientSessions) != ExpectedClientSessionLength {
		t.Errorf("AddClientSession failed, expected length %d, got %v", ExpectedClientSessionLength, len(client.clientSessions))
		return
	}

	if client.clientSessions[0].socketConn != socketConn || client.clientSessions[0].room != room {
		t.Errorf("AddClientSession failed, expected socketConn and room to match")
		return
	}

	t.Logf("AddClientSession passed")
}

func TestClientRemoveClientSession(t *testing.T) {
	client := NewClient(sampleSessionID)

	// Test RemoveClientSession
	socketConn := &websocket.Conn{}
	room := NewRoom(sampleRoomID)
	client.AddClientSession(socketConn, room, sampleSessionID)
	client.RemoveClientSession(0, sampleSessionID)

	if len(client.clientSessions) != 0 {
		t.Errorf("RemoveClientSession failed, expected length 0, got %v", len(client.clientSessions))
		return
	}

	t.Logf("RemoveClientSession passed")
}

func TestClientSendMessageToClientSession(t *testing.T) {
	client := NewClient(sampleSessionID)

	// Test SendMessageToClientSession
	socketConn := &websocket.Conn{}
	room := NewRoom(sampleRoomID)
	client.AddClientSession(socketConn, room, sampleSessionID)
	message := []byte(testMessage)

	// Start a goroutine to read from the send channel
	go func() {
		select {
		case msg := <-client.clientSessions[0].send:
			if string(msg) != string(message) {
				t.Errorf("SendMessageToClientSession failed, expected message %s, got %s", string(message), string(msg))
			}
		case <-time.After(time.Second * 1):
			t.Errorf("SendMessageToClientSession failed, no message sent")
		}
	}()

	client.SendMessageToClientSession(0, message, sampleSessionID)

	t.Logf("SendMessageToClientSession passed")
}
