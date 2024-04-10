// myapp/internal/chat/room_test.go
package chat

import (
	"strconv"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

type MockConn struct {
	LastMessageType int
	LastData        []byte
}

func (mc *MockConn) WriteMessage(messageType int, data []byte) error {
	mc.LastMessageType = messageType
	mc.LastData = make([]byte, len(data))
	copy(mc.LastData, data)
	return nil
}

func (mc *MockConn) ReadMessage() (messageType int, p []byte, err error) {
	return 0, nil, nil
}

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

func TestReceiveMessageFromClient(t *testing.T) {
	room := NewRoom(sampleRoomID)

	for i := 0; i < 3; i++ {
		client := NewClient(strconv.Itoa(i))
		conn := &MockConn{}
		room.AddClient(client, conn)
	}

	message := []byte(sampleMessage)
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
