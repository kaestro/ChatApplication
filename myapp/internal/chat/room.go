// myapp/internal/chat/models.go
package chat

import (
	"fmt"
)

// room은 클라이언트들이 메시지를 주고받을 수 있는 공간을 나타낸다.
// User가 방에 들어오고 나갈 수 있으며, 방에 있는 User들에게 메시지를 전달할 수 있다.
type room struct {
	roomName string
	// Registered map of clients to their websocket connections
	sessionIDToHandler map[string]*roomClientHandler

	// Inbound messages from the clients.
	broadcast chan []byte

	register   chan *roomClientHandler
	unregister chan *roomClientHandler

	done chan struct{}
}

// TODO: RoomManager와 상호작용 통해 새로운 RoomID의 Room을 생성하도록 변경
func newRoom(roomName string) *room {
	room := &room{
		roomName:           roomName,
		sessionIDToHandler: make(map[string]*roomClientHandler),
		broadcast:          make(chan []byte),
		register:           make(chan *roomClientHandler),
		unregister:         make(chan *roomClientHandler),
		done:               make(chan struct{}),
	}

	go room.run()

	return room
}

func (r *room) isClientInsideRoom(loginSessionID string) bool {
	_, ok := r.sessionIDToHandler[loginSessionID]
	return ok
}

func (r *room) closeRoom() {
	close(r.done)
}

// TODO: client가 있을 경우 충돌 처리
// TODO: line 58 ~ 63의 select 구문을 사용하여 room이 닫힌 경우를 middleware로 처리
func (r *room) addClient(client *client) error {
	loginSessionID := client.getLoginSessionID()
	if r.isClientInsideRoom(loginSessionID) {
		fmt.Printf("Client with sessionID %s is already inside room %s\n", loginSessionID, r.roomName)
		return nil
	}

	select {
	case <-r.done:
		return error(fmt.Errorf("room %s is closed", r.roomName))
	default:
		r.register <- newRoomClientHandler(client)
	}

	return nil
}

// TODO: room이 닫힌 경우를 middleware로 처리
func (r *room) removeClient(loginSessionID string) {
	if !r.isClientInsideRoom(loginSessionID) {
		fmt.Println("Client with sessionID", loginSessionID, "does not exist")
		return
	}

	r.unregister <- r.sessionIDToHandler[loginSessionID]
}

// TODO: Set debugging messages to be printed only when debugging is enabled
func (r *room) receiveMessageFromClient(loginSessionID string, message []byte) {
	if !r.isClientInsideRoom(loginSessionID) {
		// Debugging message
		// fmt.Println("Client with sessionID", loginSessionID, "does not exist")
		return
	}

	r.broadcast <- message
}

// client가 room에서 메시지를 읽고 쓰는 전반적인 동작을 수행한다.
// TODO: After implementing Client Object, call the chan returning method from here
// Problem: It seems too much of responsibility on Client Object. That is, it might be better for
// the Room object to have structured data of clients and websocket connections
func (r *room) run() {
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

func (r *room) registerClientHandler(clientHandler *roomClientHandler) {
	r.sessionIDToHandler[clientHandler.getLoginSessionID()] = clientHandler
}

func (r *room) unregisterClientHandler(clientHandler *roomClientHandler) {
	clientHandler.close()
	delete(r.sessionIDToHandler, clientHandler.getLoginSessionID())
}

func (r *room) broadcastMessage(message []byte) {
	for _, clientHandler := range r.sessionIDToHandler {
		clientHandler.receiveMessageFromRoom(message)
	}
}

func (r *room) getClients() []*client {
	clients := make([]*client, 0, len(r.sessionIDToHandler))
	for _, clientHandler := range r.sessionIDToHandler {
		clients = append(clients, clientHandler.client)
	}
	return clients
}
