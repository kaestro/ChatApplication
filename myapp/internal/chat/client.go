// myapp/internal/chat/models.go
package chat

import (
	"fmt"
)

// client는 채팅 서버에 접속한 클라이언트(user)를 나타낸다.
type client struct {
	loginSessionID string           // 어느 user인지 구분하는 id
	clientSessions []*clientSession // room, socket, send channel을 가지고 있는 session slice
	conn           Conn
}

func newClient(loginSessionID string, conn Conn) *client {
	client := &client{
		loginSessionID: loginSessionID,
		clientSessions: make([]*clientSession, 0),
		conn:           conn,
	}

	go client.listen()

	return client
}

// TODO
// 해당 부분이 모든 함수들에서 중복되게 사용되고 있다. 이를 빼내는 middleware 형태로 변경의 필요
func (c *client) isSameClient(loginSessionID string) bool {
	return c.loginSessionID == loginSessionID
}

func (c *client) addClientSession(room *room, loginSessionID string) {
	if !c.isSameClient(loginSessionID) {
		return
	}

	c.clientSessions = append(c.clientSessions, newClientSession(len(c.clientSessions), room))
}

func (c *client) removeClientSession(clientSessionID int, loginSessionID string) {
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

func (c *client) getLoginSessionID() string {
	return c.loginSessionID
}

// TODO: error 발생시 처리
func (c *client) listen() {
	for {
		err := receiveMessageFromConnection(c)
		if err != nil {
			return
		}
	}
}

// TODO
// chatMessage를 변경하는 부분에서 문제 발생중일 가능성 높아보임. 확인 필요
func receiveMessageFromConnection(c *client) error {
	_, message, err := c.conn.ReadMessage()
	if err != nil {
		fmt.Printf("error occurred while reading message: %v", err)
		return err
	}

	chatMessage, err := NewChatMessageFromBytes(message)
	if err != nil {
		//fmt.Printf("error occurred while unmarshalling message: %v", err)
		return err
	}

	c.sendMessageToRoom(message, chatMessage.RoomID)

	return nil
}

// TODO: roomID가 없을 경우 처리
func (c *client) sendMessageToRoom(message []byte, roomID string) {
	for _, clientSession := range c.clientSessions {
		if clientSession.room.roomName == roomID {
			clientSession.sendMessage(message)
			break
		}
	}
}
