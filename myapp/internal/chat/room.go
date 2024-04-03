// myapp/internal/chat/models.go
package chat

import (
	"fmt"
)

type Room struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

// 새로운 room을 생성한다.
func NewRoom() *Room {
	room := &Room{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}

	go room.run()

	return room
}

// client가 room에서 메시지를 읽고 쓰는 전반적인 동작을 수행한다.
// TODO: After implementing Client Object, call the chan returning method from here
// Problem: It seems too much of responsibility on Client Object. That is, it might be better for
// the Room object to have structured data of clients and websocket connections
func (r *Room) run() {
	for {
		select {
		case client := <-r.register:
			r.clients[client] = true
		case client := <-r.unregister:
			if _, ok := r.clients[client]; ok {
				delete(r.clients, client)
				close(client.send)
			}
		case message := <-r.broadcast:
			for client := range r.clients {
				for i := 0; i < 3; i++ { // 재전송을 최대 3번 시도
					select {
					case client.send <- message:
						break // 메시지 전송 성공, 재전송 시도 중단
					default:
						if i == 2 { // 재전송 횟수 초과, 로그 출력
							fmt.Println("Message delivery failed after 3 attempts")
						}
					}
				}
			}
		}
	}
}
