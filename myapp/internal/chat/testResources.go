// myapp/internal/chat/testResources.go
package chat

import "encoding/json"

var (
	maxClients                  = 10000
	sampleLoginSessionID        = "123"
	sampleUpdateID              = "456"
	sampleClient                = NewClient(sampleLoginSessionID, &MockConn{})
	maxRooms                    = 100
	sampleRoomID                = "123"
	sampleRoom                  = NewRoom(sampleRoomID)
	sampleMessage               = ChatMessage{RoomID: sampleRoomID, UserName: "user", Content: "content"}
	sampleMessageBytes, _       = json.Marshal(sampleMessage)
	ExpectedClientSessionLength = 1
)

type MockConn struct {
	LastMessageType    int
	LastData           []byte
	WriteMessageCalled bool
}

func (mc *MockConn) WriteMessage(messageType int, data []byte) error {
	mc.LastMessageType = messageType
	mc.LastData = make([]byte, len(data))
	copy(mc.LastData, data)
	mc.WriteMessageCalled = true
	return nil
}

func (mc *MockConn) GetLastData() []byte {
	return mc.LastData
}

func (mc *MockConn) ReadMessage() (messageType int, p []byte, err error) {
	return 0, nil, nil
}
