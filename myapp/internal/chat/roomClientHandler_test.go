// myapp/internal/chat/roomClientHandler_test.go
package chat

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestRoomClientHandler_sendMessageToClient(t *testing.T) {
	conn, roomClientHandler := setConnClientHandler()

	roomClientHandler.sendMessageToClient(sampleMessageBytes)

	// Wait for a short period of time to ensure that the listen goroutine has started
	time.Sleep(100 * time.Millisecond)

	var sentMessage, receivedMessage ChatMessage
	err := json.Unmarshal(sampleMessageBytes, &sentMessage)
	if err != nil {
		t.Fatalf("Failed to unmarshal sent message: %v", err)
	}

	err = json.Unmarshal(conn.LastData, &receivedMessage)
	if err != nil {
		t.Fatalf("Failed to unmarshal received message: %v", err)
	}

	if !reflect.DeepEqual(sentMessage, receivedMessage) {
		t.Errorf("Expected message '%v', but got '%v'", sentMessage, receivedMessage)
	}
}

func TestRoomClientHandler_receiveMessageFromRoom(t *testing.T) {
	conn, roomClientHandler := setConnClientHandler()

	roomClientHandler.receiveMessageFromRoom(sampleMessageBytes)

	time.Sleep(100 * time.Millisecond)

	var sentMessage, receivedMessage ChatMessage
	err := json.Unmarshal(sampleMessageBytes, &sentMessage)
	if err != nil {
		t.Fatalf("Failed to unmarshal sent message: %v", err)
	}

	err = json.Unmarshal(conn.LastData, &receivedMessage)
	if err != nil {
		t.Fatalf("Failed to unmarshal received message: %v", err)
	}

	if !reflect.DeepEqual(sentMessage, receivedMessage) {
		t.Errorf("Expected message '%v', but got '%v'", sentMessage, receivedMessage)
	}
}

func TestRoomClientHandler_close(t *testing.T) {
	roomClientHandler := NewRoomClientHandler(sampleClient)

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

func setConnClientHandler() (*MockConn, *RoomClientHandler) {
	conn := &MockConn{}
	client := NewClient(sampleLoginSessionID, conn)
	roomClientHandler := NewRoomClientHandler(client)
	return conn, roomClientHandler
}
