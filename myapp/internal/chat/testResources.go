// myapp/internal/chat/testResources.go
package chat

import (
	"encoding/json"
)

const (
	maxClients                  = 10000
	sampleLoginSessionID        = "123"
	sampleUpdateID              = "456"
	maxRooms                    = 100
	sampleRoomName              = "123"
	expectedClientSessionLength = 1
)

var (
	sampleClient          = newClient(sampleLoginSessionID, &mockConn{})
	sampleRoom            = newRoom(sampleRoomName)
	sampleMessage         = ChatMessage{RoomID: sampleRoomName, UserName: "user", Content: "content"}
	sampleMessageBytes, _ = json.Marshal(sampleMessage)
)

type mockConn struct {
	LastMessageType    int
	LastData           []byte
	WriteMessageCalled bool
	message            []byte
}

func (mc *mockConn) WriteMessage(messageType int, data []byte) error {
	mc.LastMessageType = messageType
	mc.LastData = make([]byte, len(data))
	copy(mc.LastData, data)
	mc.WriteMessageCalled = true
	return nil
}

func (mc *mockConn) ReadMessage() (messageType int, p []byte, err error) {
	return 0, mc.message, nil
}
