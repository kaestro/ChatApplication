// myapp/internal/chat/conn.go
package chat

type Conn interface {
	WriteMessage(messageType int, data []byte) error
	ReadMessage() (messageType int, p []byte, err error)
}
