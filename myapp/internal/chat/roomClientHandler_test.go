// myapp/internal/chat/roomClientHandler_test.go
package chat

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestRoomClientHandler_sendMessageToClient(t *testing.T) {
	conn := &MockConn{}
	client := NewClient(sampleLoginSessionID, conn)
	roomClientHandler := NewRoomClientHandler(client)

	roomClientHandler.receive <- sampleMessageBytes

	// Wait for a short period of time to ensure that the listen goroutine has started
	time.Sleep(100 * time.Millisecond)

	var sentMessage, receivedMessage ChatMessage
	err := json.Unmarshal(sampleMessageBytes, &sentMessage)
	if err != nil {
		t.Fatalf("Failed to unmarshal sent message: %v", err)
	}
	fmt.Println("sentMessage: ", sentMessage)

	err = json.Unmarshal(conn.LastData, &receivedMessage)
	if err != nil {
		t.Fatalf("Failed to unmarshal received message: %v", err)
	}

	if !reflect.DeepEqual(sentMessage, receivedMessage) {
		t.Errorf("Expected message '%v', but got '%v'", sentMessage, receivedMessage)
	}
}

func TestRoomClientHandler_listen(t *testing.T) {
	/*
		mockConn := &MockConn{}
		roomClientHandler := NewRoomClientHandler(sampleClient)

		roomClientHandler.receiveMessageFromRoom([]byte(sampleMessage))

		// Wait for a short period of time to ensure that the listen goroutine has started
		time.Sleep(100 * time.Millisecond)

		if string(mockConn.LastData) != sampleMessage {
			t.Errorf("Expected message '%s', but got '%s'", sampleMessage, string(mockConn.LastData))
		}
	*/
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
