// myapp/internal/chat/room_test.go
package chat

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewRoom(t *testing.T) {
	room := NewRoom(sampleRoomID)

	assert.NotNil(t, room)
	assert.Equal(t, sampleRoomID, room.roomID)

	t.Logf("TestNewRoom passed")
}

func TestAddClient(t *testing.T) {
	room := NewRoom(sampleRoomID)
	client := NewClient(sampleLoginSessionID, &MockConn{})

	room.AddClient(client)
	time.Sleep(time.Millisecond * 100)

	assert.True(t, room.IsClientInsideRoom(sampleLoginSessionID))
}

func TestRemoveClient(t *testing.T) {
	room := NewRoom(sampleRoomID)
	client := NewClient(sampleLoginSessionID, &MockConn{})

	room.AddClient(client)
	time.Sleep(time.Millisecond * 100)

	room.RemoveClient(sampleLoginSessionID)
	time.Sleep(time.Millisecond * 100)

	assert.False(t, room.IsClientInsideRoom(sampleLoginSessionID))
}

func TestCloseRoom(t *testing.T) {
	room := NewRoom(sampleRoomID)
	room.closeRoom()

	// Check if room is closed by trying to add a client
	client := NewClient(sampleLoginSessionID, &MockConn{})
	room.AddClient(client)
	time.Sleep(time.Millisecond * 100)

	// If room is closed, client should not be added
	assert.False(t, room.IsClientInsideRoom(sampleLoginSessionID))
}

func TestReceiveMessageFromClient(t *testing.T) {
	room := NewRoom(sampleRoomID)

	for i := 0; i < 3; i++ {
		client := NewClient(strconv.Itoa(i), &MockConn{})
		room.AddClient(client)
	}

	message := sampleMessageBytes
	room.ReceiveMessageFromClient("0", message)
	time.Sleep(time.Millisecond * 100)

	for loginSessionID, handler := range room.sessionIDToHandler {
		if !assert.Equal(t, message, handler.conn.(*MockConn).LastData) {
			t.Errorf("TestReceiveMessageFromClient %s failed", loginSessionID)
		} else {
			t.Logf("TestReceiveMessageFromClient %s passed", loginSessionID)
		}
	}
}

func TestGetClients(t *testing.T) {
	room := NewRoom(sampleRoomID)
	numClients := 3

	loginSessionIDs := make(map[string]bool)

	// Add some clients to the room
	for i := 0; i < numClients; i++ {
		conn := &MockConn{}
		client := NewClient(strconv.Itoa(i), conn)
		room.AddClient(client)
		loginSessionIDs[strconv.Itoa(i)] = false
	}

	// Get the clients from the room
	clients := room.GetClients()

	time.Sleep(time.Millisecond * 2000)

	// Check if the correct number of clients was returned
	assert.Equal(t, numClients, len(clients))

	// Check if the correct clients were returned
	for _, client := range clients {
		loginSessionID := client.GetLoginSessionID()
		if _, ok := loginSessionIDs[loginSessionID]; ok {
			loginSessionIDs[loginSessionID] = true
		}
	}

	for loginSessionID, isFound := range loginSessionIDs {
		if !isFound {
			t.Errorf("TestGetClients failed, expected loginSessionID %s to be found", loginSessionID)
			return
		}
	}

	t.Logf("TestGetClients passed")
}
