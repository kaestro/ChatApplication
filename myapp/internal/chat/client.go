// myapp/internal/chat/models.go
package chat

import (
	"encoding/json"
	"log"
)

// client는 채팅 서버에 접속한 클라이언트(user)를 나타낸다.
type Client struct {
	loginSessionID string           // 어느 user인지 구분하는 id
	clientSessions []*ClientSession // room, socket, send channel을 가지고 있는 session slice
	conn           Conn
}

func NewClient(loginSessionID string, conn Conn) *Client {
	return &Client{
		loginSessionID: loginSessionID,
		clientSessions: make([]*ClientSession, 0),
		conn:           conn,
	}
}

// TODO
// 해당 부분이 모든 함수들에서 중복되게 사용되고 있다. 이를 빼내는 middleware 형태로 변경의 필요
func (c *Client) isSameClient(loginSessionID string) bool {
	return c.loginSessionID == loginSessionID
}

func (c *Client) AddClientSession(room *Room, loginSessionID string) {
	if !c.isSameClient(loginSessionID) {
		return
	}

	c.clientSessions = append(c.clientSessions, NewClientSession(len(c.clientSessions), room))
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

func (c *Client) GetLoginSessionID() string {
	return c.loginSessionID
}

func (c *Client) listen() {
	for {
		message := receiveMessageFromClient(c)
		if message != nil {
			c.sendMessageToRoom(message)
		}
	}
}

func receiveMessageFromClient(c *Client) []byte {
	_, message, err := c.conn.ReadMessage()
	if err != nil {
		log.Printf("error occurred while reading message: %v", err)
		return nil
	}

	var msg ChatMessage
	err = json.Unmarshal(message, &msg)
	if err != nil {
		log.Printf("error occurred while unmarshalling message: %v", err)
		return nil
	}

	return message
}

func (c *Client) sendMessageToRoom(message []byte) {
	for _, clientSession := range c.clientSessions {
		clientSession.sendMessage(message)
	}
}
