// myapp/internal/chat/room_test.go
package chat

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewRoom(t *testing.T) {
	room := newRoom(sampleRoomName)

	assert.NotNil(t, room)
	assert.Equal(t, sampleRoomName, room.roomName)

	t.Logf("TestNewRoom passed")
}

func TestAddClient(t *testing.T) {
	room := newRoom(sampleRoomName)
	client := newClient(sampleLoginSessionID, &mockConn{})

	room.addClient(client)
	time.Sleep(time.Millisecond * 100)

	assert.True(t, room.isClientInsideRoom(sampleLoginSessionID))
}

func TestRemoveClient(t *testing.T) {
	room := newRoom(sampleRoomName)
	client := newClient(sampleLoginSessionID, &mockConn{})

	room.addClient(client)
	time.Sleep(time.Millisecond * 100)

	room.removeClient(sampleLoginSessionID)
	time.Sleep(time.Millisecond * 100)

	assert.False(t, room.isClientInsideRoom(sampleLoginSessionID))
}

func TestCloseRoom(t *testing.T) {
	room := newRoom(sampleRoomName)
	room.closeRoom()

	// Check if room is closed by trying to add a client
	client := newClient(sampleLoginSessionID, &mockConn{})
	room.addClient(client)
	time.Sleep(time.Millisecond * 100)

	// If room is closed, client should not be added
	assert.False(t, room.isClientInsideRoom(sampleLoginSessionID))
}

func TestReceiveMessageFromClient(t *testing.T) {
	room := newRoom(sampleRoomName)

	for i := 0; i < 3; i++ {
		client := newClient(strconv.Itoa(i), &mockConn{})
		room.addClient(client)
	}

	message := sampleMessageBytes
	room.receiveMessageFromClient("0", message)
	time.Sleep(time.Millisecond * 100)

	for loginSessionID, handler := range room.sessionIDToHandler {
		if !assert.Equal(t, message, handler.conn.(*mockConn).LastData) {
			t.Errorf("TestReceiveMessageFromClient %s failed", loginSessionID)
		} else {
			t.Logf("TestReceiveMessageFromClient %s passed", loginSessionID)
		}
	}
}

func TestGetClients(t *testing.T) {
	room := newRoom(sampleRoomName)
	numClients := 3

	loginSessionIDs := make(map[string]bool)

	// Add some clients to the room
	for i := 0; i < numClients; i++ {
		conn := &mockConn{}
		client := newClient(strconv.Itoa(i), conn)
		room.addClient(client)
		loginSessionIDs[strconv.Itoa(i)] = false
	}

	// Get the clients from the room
	clients := room.getClients()

	time.Sleep(time.Millisecond * 2500)

	// Check if the correct number of clients was returned
	if !assert.Equal(t, numClients, len(clients)) {
		// 단순히 시간이 오래 걸리는 경우가 있어서, 다시 한번 시도
		time.Sleep(time.Millisecond * 3000)
		assert.Equal(t, numClients, len(clients))
	}

	// Check if the correct clients were returned
	for _, client := range clients {
		loginSessionID := client.getLoginSessionID()
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
