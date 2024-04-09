// myapp/internal/chat/room_test.go
package chat

import (
	"testing"
	"time"

	"github.com/gorilla/websocket"
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
	client := NewClient(sampleLoginSessionID)
	conn := &websocket.Conn{}

	room.AddClient(client, conn)
	time.Sleep(time.Millisecond * 100)

	assert.True(t, room.IsClientInsideRoom(sampleLoginSessionID))
}

func TestRemoveClient(t *testing.T) {
	room := NewRoom(sampleRoomID)
	client := NewClient(sampleLoginSessionID)
	conn := &websocket.Conn{}

	room.AddClient(client, conn)
	time.Sleep(time.Millisecond * 100)

	room.RemoveClient(sampleLoginSessionID)
	time.Sleep(time.Millisecond * 100)

	assert.False(t, room.IsClientInsideRoom(sampleLoginSessionID))
}

func TestCloseRoom(t *testing.T) {
	room := NewRoom(sampleRoomID)
	room.closeRoom()

	// Check if room is closed by trying to add a client
	client := NewClient(sampleLoginSessionID)
	conn := &websocket.Conn{}
	room.AddClient(client, conn)
	time.Sleep(time.Millisecond * 100)

	// If room is closed, client should not be added
	assert.False(t, room.IsClientInsideRoom(sampleLoginSessionID))
}
