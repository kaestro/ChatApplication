// myapp/internal/chat/testResources.go
package chat

var (
	maxClients                  = 10000
	sampleLoginSessionID        = "123"
	sampleUpdateID              = "456"
	sampleClient                = NewClient(sampleLoginSessionID)
	maxRooms                    = 100
	sampleRoomID                = "123"
	sampleRoom                  = NewRoom(sampleRoomID)
	sampleMessage               = "test message"
	ExpectedClientSessionLength = 1
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
