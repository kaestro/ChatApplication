// myapp/internal/chat/roomClientHandler_test.go
package chat

import (
	"testing"
	"time"
)

func TestRoomClientHandler_sendMessageToClient(t *testing.T) {
	mockConn := &MockConn{}
	roomClientHandler := NewRoomClientHandler(sampleClient, mockConn)

	roomClientHandler.receive <- []byte(sampleMessage)

	// Wait for a short period of time to ensure that the listen goroutine has started
	time.Sleep(100 * time.Millisecond)

	if string(mockConn.LastData) != sampleMessage {
		t.Errorf("Expected message '%s', but got '%s'", sampleMessage, string(mockConn.LastData))
	}
}

func TestRoomClientHandler_listen(t *testing.T) {
	mockConn := &MockConn{}
	roomClientHandler := NewRoomClientHandler(sampleClient, mockConn)

	roomClientHandler.receiveMessageFromRoom([]byte(sampleMessage))

	// Wait for a short period of time to ensure that the listen goroutine has started
	time.Sleep(100 * time.Millisecond)

	if string(mockConn.LastData) != sampleMessage {
		t.Errorf("Expected message '%s', but got '%s'", sampleMessage, string(mockConn.LastData))
	}
}

func TestRoomClientHandler_close(t *testing.T) {
	mockConn := &MockConn{}
	roomClientHandler := NewRoomClientHandler(sampleClient, mockConn)

	roomClientHandler.close()

	// Wait for a short period of time to ensure that the listen goroutine has stopped
	time.Sleep(100 * time.Millisecond)

	select {
	case <-roomClientHandler.done:
		// The done channel is closed, which is expected
	default:
		t.Errorf("Expected roomClientHandler.done to be closed, but it's not")
	}
}
