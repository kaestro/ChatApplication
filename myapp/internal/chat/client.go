// myapp/internal/chat/models.go
package chat

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	// 어느 user인지 구분할 방법
	loginSessionID string
	clientSessions []*ClientSession
}

type ClientSession struct {
	id         int
	socketConn *websocket.Conn
	room       *Room
	send       chan []byte
}

func NewClient(loginSessionID string) *Client {
	return &Client{
		loginSessionID: loginSessionID,
	}
}

// TODO
// 해당 부분이 모든 함수들에서 중복되게 사용되고 있다. 이를 빼내는 middleware 형태로 변경의 필요
func (c *Client) isSameClient(loginSessionID string) bool {
	return c.loginSessionID == loginSessionID
}

func (c *Client) AddClientSession(socketConn *websocket.Conn, room *Room, loginSessionID string) {
	if !c.isSameClient(loginSessionID) {
		return
	}

	c.clientSessions = append(c.clientSessions, &ClientSession{
		id:         len(c.clientSessions),
		socketConn: socketConn,
		room:       room,
		send:       make(chan []byte),
	})
}

func (c *Client) RemoveClientSession(id int, loginSessionID string) {
	if !c.isSameClient(loginSessionID) {
		return
	}

	for i, conn := range c.clientSessions {
		if conn.id == id {
			c.clientSessions = append(c.clientSessions[:i], c.clientSessions[i+1:]...)
			break
		}
	}
}

func (c *Client) SendMessageToClientSession(id int, message []byte, loginSessionID string) {
	if !c.isSameClient(loginSessionID) {
		return
	}

	for _, session := range c.clientSessions {
		if session.id == id {
			session.send <- message
			break
		}
	}
}
