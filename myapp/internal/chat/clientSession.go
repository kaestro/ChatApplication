// myapp/internal/chat/clientSession.go
package chat

import (
	"log"
)

// ClientSession은 방과 클라이언트의 중개자 역할을 한다.
// 클라이언트와 방 사이의 연결을 유지하고, 클라이언트의 메시지를 받아 해당 방으로 전달한다.
type ClientSession struct {
	id               int
	loginSessionID   string
	socketConnection Conn
	room             *Room
}

func NewClientSession(id int, conn Conn, room *Room) *ClientSession {
	clientSession := &ClientSession{
		id:               id,
		socketConnection: conn,
		room:             room,
	}

	go clientSession.listen()

	return clientSession
}

func (cs *ClientSession) listen() {
	for {
		_, message, err := cs.socketConnection.ReadMessage()
		if err != nil {
			log.Printf("error occurred while reading message: %v", err)
			break
		}

		if message != nil {
			cs.room.ReceiveMessageFromClient(cs.loginSessionID, message)
		}
	}
}
