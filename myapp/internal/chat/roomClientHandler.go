// myapp/internal/chat/roomClientHandler.go
package chat

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type RoomClientHandler struct {
	client  *Client
	conn    Conn
	receive chan []byte
	done    chan struct{}
}

func NewRoomClientHandler(client *Client, conn Conn) *RoomClientHandler {
	roomClientHandler := &RoomClientHandler{
		client:  client,
		conn:    conn,
		receive: make(chan []byte),
		done:    make(chan struct{}),
	}

	go roomClientHandler.listen()

	return roomClientHandler
}

func (rch *RoomClientHandler) getLoginSessionID() string {
	return rch.client.loginSessionID
}

func (rch *RoomClientHandler) receiveMessageFromRoom(message []byte) {
	rch.receive <- message
}

func (rch *RoomClientHandler) close() {
	close(rch.done)
}

func (rch *RoomClientHandler) listen() {
	for {
		select {
		case message := <-rch.receive:
			rch.sendMessageToClient(message)
		case <-rch.done:
			return
		}
	}
}

// TODO: fmt 대신 별개의 로거를 사용하도록 변경
func (rch *RoomClientHandler) sendMessageToClient(message []byte) {
	err := rch.conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		fmt.Println("Error writing message to websocket:", err)
	}
}
