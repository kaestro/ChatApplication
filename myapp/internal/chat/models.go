// myapp/internal/chat/models.go
package chat

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// TODO
// Client랑 Room 작성할만큼 해 둔 다음에, 두개 분리된 go파일로 만드는게 맞을지 고민
// 근데 그냥 하는게 맞긴 할 것 같은듯?
// Client는 괜찮은데 Room은 더 일반적인 명사를 통해 추상화된 객체 이름으로 작성하는게 나을수도 있을듯?

type Client struct {
	// 웹소켓 커넥션
	conn *websocket.Conn

	// 어느 user인지 구분할 방법. 세션id?
	sessionID string

	// Buffered channel of outbound messages.
	// 어느 방에서 보내는지에 따라 메시지가 따로 가야하는데 이건 어떻게 해야하는가?
	send chan []byte

	// client가 속한 방들
	// rooms를 pointer로 만들어야하는가?
	rooms []Room
}

func (c *Client) enterRoom(room *Room) {
	c.rooms = append(c.rooms, *room)
}

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
