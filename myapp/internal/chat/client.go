// myapp/internal/chat/models.go
package chat

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	loginSessionID string           // 어느 user인지 구분하는 id
	clientSessions []*ClientSession // room, socket, send channel을 가지고 있는 session slice
}

type ClientSession struct {
	id               int
	socketConnection *websocket.Conn
	room             *Room
	send             chan []byte // 메시지를 보내는 채널. 동시에 여러 클라이언트에서 메시지를 보낼 수 있도록 하기 위해 사용
}

func NewClient(loginSessionID string) *Client {
	return &Client{
		loginSessionID: loginSessionID,
		clientSessions: make([]*ClientSession, 0),
	}
}

// TODO
// 해당 부분이 모든 함수들에서 중복되게 사용되고 있다. 이를 빼내는 middleware 형태로 변경의 필요
func (c *Client) isSameClient(loginSessionID string) bool {
	return c.loginSessionID == loginSessionID
}

func (c *Client) AddClientSession(conn *websocket.Conn, room *Room, loginSessionID string) {
	if !c.isSameClient(loginSessionID) {
		return
	}

	c.clientSessions = append(c.clientSessions, &ClientSession{
		id:               len(c.clientSessions),
		socketConnection: conn,
		room:             room,
		send:             make(chan []byte),
	})
}

func (c *Client) RemoveClientSession(clientSessionID int, loginSessionID string) {
	if !c.isSameClient(loginSessionID) {
		return
	}

	for i, clientSession := range c.clientSessions {
		if clientSession.id == clientSessionID {
			c.clientSessions = append(c.clientSessions[:i], c.clientSessions[i+1:]...)
			break
		}
	}
}

func (c *Client) SendMessageToClientSession(clientSessionID int, message []byte, loginSessionID string) {
	if !c.isSameClient(loginSessionID) {
		return
	}

	for _, clientSession := range c.clientSessions {
		if clientSession.id == clientSessionID {
			clientSession.send <- message
			break
		}
	}
}

func (c *Client) GetLoginSessionID() string {
	return c.loginSessionID
}
