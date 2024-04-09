// myapp/internal/chat/models.go
package chat

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Room struct {
	roomID string
	// Registered map of clients to their websocket connections
	sessionIDToHandler map[string]*RoomClientHandler

	// Inbound messages from the clients.
	broadcast chan []byte

	register   chan *RoomClientHandler
	unregister chan *RoomClientHandler

	done chan struct{}
}

// TODO: RoomManager와 상호작용 통해 새로운 RoomID의 Room을 생성하도록 변경
func NewRoom(roomId string) *Room {
	room := &Room{
		roomID:             roomId,
		sessionIDToHandler: make(map[string]*RoomClientHandler),
		broadcast:          make(chan []byte),
		register:           make(chan *RoomClientHandler),
		unregister:         make(chan *RoomClientHandler),
		done:               make(chan struct{}),
	}

	go room.run()

	return room
}

func (r *Room) IsClientInsideRoom(loginSessionID string) bool {
	_, ok := r.sessionIDToHandler[loginSessionID]
	return ok
}

func (r *Room) closeRoom() {
	close(r.done)
}

// TODO: client가 있을 경우 충돌 처리
func (r *Room) AddClient(client *Client, conn *websocket.Conn) {
	loginSessionID := client.GetLoginSessionID()
	if r.IsClientInsideRoom(loginSessionID) {
		fmt.Println("Client with sessionID", loginSessionID, "already exists")
		return
	}

	select {
	case <-r.done:
		fmt.Println("Room is closed, can't add client")
		return
	default:
		r.register <- NewRoomClientHandler(client, conn)
	}
}

func (r *Room) RemoveClient(loginSessionID string) {
	if !r.IsClientInsideRoom(loginSessionID) {
		fmt.Println("Client with sessionID", loginSessionID, "does not exist")
		return
	}

	r.unregister <- r.sessionIDToHandler[loginSessionID]
}

// client가 room에서 메시지를 읽고 쓰는 전반적인 동작을 수행한다.
// TODO: After implementing Client Object, call the chan returning method from here
// Problem: It seems too much of responsibility on Client Object. That is, it might be better for
// the Room object to have structured data of clients and websocket connections
func (r *Room) run() {
	for {
		select {
		case clientHandler := <-r.register:
			r.registerClientHandler(clientHandler)
		case clientHandler := <-r.unregister:
			r.unregisterClientHandler(clientHandler)
		case message := <-r.broadcast:
			r.broadcastMessage(message)
		case <-r.done:
			return
		}
	}
}

func (r *Room) registerClientHandler(clientHandler *RoomClientHandler) {
	r.sessionIDToHandler[clientHandler.getLoginSessionID()] = clientHandler
}

func (r *Room) unregisterClientHandler(clientHandler *RoomClientHandler) {
	clientHandler.close()
	delete(r.sessionIDToHandler, clientHandler.getLoginSessionID())
}

func (r *Room) broadcastMessage(message []byte) {
	for _, clientHandler := range r.sessionIDToHandler {
		clientHandler.receiveMessageFromRoom(message)
	}
}
