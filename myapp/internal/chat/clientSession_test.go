// myapp/internal/chat/clientSession_test.go
package chat

import (
	"testing"
	"time"
)

func TestClientSession_listen(t *testing.T) {
	room := newRoom(sampleRoomName)
	clientSession := newClientSession(0, room)

	go clientSession.listen()

	// Wait for a short period of time to ensure that the listen goroutine has started
	time.Sleep(100 * time.Millisecond)

	select {
	case <-room.broadcast:
		t.Errorf("Expected no messages in room.broadcast, but got one")
	default:
		// No message in room.broadcast, which is expected
	}
}
