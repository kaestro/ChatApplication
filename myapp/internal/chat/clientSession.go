// myapp/internal/chat/clientSession.go
package chat

import (
	"log"
	"strconv"
)

type ClientSession struct {
	id               int
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

		cs.room.ReceiveMessageFromClient(strconv.Itoa(cs.id), message)
	}
}
