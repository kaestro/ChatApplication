// myapp/internal/chat/models.go
package chat

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	// 어느 user인지 구분할 방법
	sessionID string

	conn []*Connection
}

type Connection struct {
	id   int
	conn *websocket.Conn
	room *Room
	send chan []byte
}
