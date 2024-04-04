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

// Question: Do Client Connection actually need conn object?
func (c *Client) enterRoom(room *Room) {
	// Check if the client is already in the room
	// if already in, return

	// if not, add another Connection object to the client
	// Question: How to generate a unique id for the connection?
	// Assuming that each clients won't be having many connections, simple slice can be used
}
