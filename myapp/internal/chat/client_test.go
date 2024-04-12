// myapp/internal/chat/client_test.go
package chat

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestIsSameClient(t *testing.T) {
	conn := &mockConn{}
	client := newClient(sampleLoginSessionID, conn)

	// Test isSameClient with the same session ID
	if !client.isSameClient(sampleLoginSessionID) {
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
	conn := &mockConn{}
	client := newClient(sampleLoginSessionID, conn)

	// Test AddClientSession
	room := newRoom(sampleRoomName)
	client.addClientSession(room, sampleLoginSessionID)

	if len(client.clientSessions) != expectedClientSessionLength {
		t.Errorf("AddClientSession failed, expected length %d, got %v", expectedClientSessionLength, len(client.clientSessions))
		return
	}

	if client.clientSessions[0].id != len(client.clientSessions)-1 {
		t.Errorf("AddClientSession failed, expected id 0, got %v", client.clientSessions[0].id)
		return
	}

	t.Logf("AddClientSession passed")
}

func TestClientRemoveClientSession(t *testing.T) {
	client := newClient(sampleLoginSessionID, &mockConn{})

	// Test RemoveClientSession
	room := newRoom(sampleRoomName)
	client.addClientSession(room, sampleLoginSessionID)
	client.removeClientSession(0, sampleLoginSessionID)

	if len(client.clientSessions) != 0 {
		t.Errorf("RemoveClientSession failed, expected length 0, got %v", len(client.clientSessions))
		return
	}

	t.Logf("RemoveClientSession passed")
}

func TestGetClientGetLoginSessionID(t *testing.T) {
	client := newClient(sampleLoginSessionID, &mockConn{})

	if client.getLoginSessionID() != sampleLoginSessionID {
		t.Errorf("GetLoginSessionID failed, expected %s, got %s", sampleLoginSessionID, client.getLoginSessionID())
	}

	t.Logf("GetLoginSessionID passed")
}

func TestListen(t *testing.T) {
	conn := &mockConn{}
	client := &client{
		loginSessionID: sampleLoginSessionID,
		clientSessions: make([]*clientSession, 0),
		conn:           conn,
	}

	go client.listen()

	// Wait for listen to process the message
	time.Sleep(time.Second)

	conn.WriteMessage(0, sampleMessageBytes)

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

func TestSendMessageToRoom(t *testing.T) {
	conn := &mockConn{}
	conn.WriteMessage(0, sampleMessageBytes)
	client := newClient(sampleLoginSessionID, conn)

	// Add a client session with a room
	room := newRoom(sampleRoomName)
	client.addClientSession(room, sampleLoginSessionID)

	// Send a message to the room
	client.sendMessageToRoom(sampleMessageBytes, sampleRoomName)

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
