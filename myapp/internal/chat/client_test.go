// myapp/internal/chat/client_test.go
package chat

import (
	"testing"
)

func TestIsSameClient(t *testing.T) {
	client := NewClient(sampleLoginSessionID)

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
	client := NewClient(sampleLoginSessionID)

	// Test AddClientSession
	socketConn := &MockConn{}
	room := NewRoom(sampleRoomID)
	client.AddClientSession(socketConn, room, sampleLoginSessionID)

	if len(client.clientSessions) != ExpectedClientSessionLength {
		t.Errorf("AddClientSession failed, expected length %d, got %v", ExpectedClientSessionLength, len(client.clientSessions))
		return
	}

	if client.clientSessions[0].socketConnection != socketConn || client.clientSessions[0].room != room {
		t.Errorf("AddClientSession failed, expected socketConn and room to match")
		return
	}

	t.Logf("AddClientSession passed")
}

func TestClientRemoveClientSession(t *testing.T) {
	client := NewClient(sampleLoginSessionID)

	// Test RemoveClientSession
	socketConn := &MockConn{}
	room := NewRoom(sampleRoomID)
	client.AddClientSession(socketConn, room, sampleLoginSessionID)
	client.RemoveClientSession(0, sampleLoginSessionID)

	if len(client.clientSessions) != 0 {
		t.Errorf("RemoveClientSession failed, expected length 0, got %v", len(client.clientSessions))
		return
	}

	t.Logf("RemoveClientSession passed")
}

func TestGetClientGetLoginSessionID(t *testing.T) {
	client := NewClient(sampleLoginSessionID)

	if client.GetLoginSessionID() != sampleLoginSessionID {
		t.Errorf("GetLoginSessionID failed, expected %s, got %s", sampleLoginSessionID, client.GetLoginSessionID())
	}

	t.Logf("GetLoginSessionID passed")
}
