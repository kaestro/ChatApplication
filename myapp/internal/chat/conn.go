// myapp/internal/chat/conn.go
package chat

// Conn은 websocket.Conn의 인터페이스를 정의한다.
type Conn interface {
	WriteMessage(messageType int, data []byte) error
	ReadMessage() (messageType int, p []byte, err error)
}
