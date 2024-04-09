// myapp/internal/chat/client_test.go
package chat

import (
	"testing"

	"github.com/gorilla/websocket"
)

const (
	ExpectedConnLength = 1
)

func TestClientAddConnection(t *testing.T) {
	client := &Client{sessionID: sampleSessionID}

	// Test AddConnection
	conn := &websocket.Conn{}
	room := &Room{roomID: sampleRoomID}
	client.AddConnection(conn, room)

	if len(client.conn) != ExpectedConnLength {
		t.Errorf("AddConnection failed, expected length %d, got %v", ExpectedConnLength, len(client.conn))
		return
	}

	if client.conn[0].conn != conn || client.conn[0].room != room {
		t.Errorf("AddConnection failed, expected conn and room to match")
		return
	}

	t.Logf("AddConnection passed")
}
