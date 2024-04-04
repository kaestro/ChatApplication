// myapp/internal/chat/models.go
package chat

import (
	"github.com/gorilla/websocket"
)

type Room struct {
	roomID string
	// Registered map of clients to their websocket connections
	client_chan map[*Client]*websocket.Conn

	// Inbound messages from the clients.
	broadcast chan []byte

	register   chan *ClientConn
	unregister chan *ClientConn
}

type ClientConn struct {
	client *Client
	conn   *websocket.Conn
}

// TODO: RoomManager와 상호작용 통해 새로운 RoomID의 Room을 생성하도록 변경
func NewRoom() *Room {
	room := &Room{
		client_chan: make(map[*Client]*websocket.Conn),
		broadcast:   make(chan []byte),
		register:    make(chan *ClientConn),
		unregister:  make(chan *ClientConn),
	}

	go room.run()

	return room
}

func (r *Room) AddClient(client *Client, conn *websocket.Conn) {
	r.register <- &ClientConn{client: client, conn: conn}
}

// client가 room에서 메시지를 읽고 쓰는 전반적인 동작을 수행한다.
// TODO: After implementing Client Object, call the chan returning method from here
// Problem: It seems too much of responsibility on Client Object. That is, it might be better for
// the Room object to have structured data of clients and websocket connections
func (r *Room) run() {
	/*
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
	*/
}
