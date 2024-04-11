// myapp/internal/chat/chatManager_test.go
// myapp/internal/chat/chatManager_test.go
package chat

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

const (
	chatManagerUserCount = 10
)

func TestChatManager(t *testing.T) {
	cm := NewChatManager()

	// Start a test server
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loginSessionID := r.URL.Query().Get("sessionID")
		err := cm.ProvideClientToUser(w, r, loginSessionID)
		assert.Nil(t, err)
	}))
	defer s.Close()

	// Create a new websocket.Dialer
	dialer := websocket.Dialer{}

	// Send multiple requests
	for i := 0; i < chatManagerUserCount; i++ {
		// Create a new websocket connection
		conn, resp, err := dialer.Dial("ws"+s.URL[4:]+"?sessionID="+strconv.Itoa(i), nil)
		assert.Nil(t, err)
		if !assert.Equal(t, http.StatusSwitchingProtocols, resp.StatusCode) {
			t.Logf("Response status code: %d", resp.StatusCode)
			return
		}

		_ = conn.WriteMessage(websocket.TextMessage, []byte("Hello, World!"))

		// Defer the closing of the connection
		if conn != nil {
			defer conn.Close()
		}
	}

	t.Logf("TestChatManager passed")

	// test removing client
	for i := 0; i < chatManagerUserCount; i++ {
		cm.RemoveClientFromUser(strconv.Itoa(i))
		cmInstance = getClientManager()
		if cmInstance.isClientRegistered(strconv.Itoa(i)) {
			t.Errorf("RemoveClientFromUser failed, expected sessionID %d to be removed", i)
			return
		}
	}

	t.Logf("TestRemoveClientFromUser passed")
}

func TestCreateRoom(t *testing.T) {
	cm := NewChatManager()
	roomName := sampleRoomName

	err := cm.CreateRoom(roomName)
	if err != nil {
		t.Errorf("Failed to create room: %v", err)
		return
	}

	rmInstance := getRoomManager()
	room, ok := rmInstance.rooms[roomName]
	if !ok || room == nil {
		t.Errorf("Room was not created")
		return
	} else if room.roomName != roomName {
		t.Errorf("Room name mismatch, expected %s, got %s", roomName, room.roomName)
		return
	}

	t.Logf("TestCreateRoom passed")
}

func TestRemoveRoom(t *testing.T) {
	cm := NewChatManager()
	cm.CreateRoom(sampleRoomName)

	err := cm.RemoveRoom(sampleRoomName)
	if err != nil {
		t.Errorf("Failed to remove room: %v", err)
		return
	}

	rmInstance := getRoomManager()
	_, ok := rmInstance.rooms[sampleRoomName]
	if ok {
		t.Errorf("Room was not removed")
		return
	}

	t.Logf("TestRemoveRoom passed")
}

func TestClientEnterRoom(t *testing.T) {
	cm := NewChatManager()

	// Create a room and a client
	err := cm.CreateRoom(sampleRoomName)
	if err != nil {
		t.Errorf("Failed to create room: %v", err)
		return
	}
	cm.registerNewClient(sampleLoginSessionID, &mockConn{})

	// Call ClientEnterRoom method
	err = cm.ClientEnterRoom(sampleRoomName, sampleLoginSessionID)
	if err != nil {
		t.Errorf("Failed to enter room: %v", err)
		return
	}

	time.Sleep(100 * time.Millisecond)

	// Check if the client is in the room
	rmInstance := getRoomManager()
	room, ok := rmInstance.rooms[sampleRoomName]
	if !ok || room == nil {
		t.Errorf("Room was not found")
		return
	}
	if !room.isClientInsideRoom(sampleLoginSessionID) {
		t.Errorf("Client was not added to the room")
		return
	}

	t.Logf("TestClientEnterRoom passed")
}
