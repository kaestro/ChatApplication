// myapp/internal/chat/clientSession.go
package chat

// ClientSession은 방과 클라이언트의 중개자 역할을 한다.
// 클라이언트와 방 사이의 연결을 유지하고, 클라이언트의 메시지를 받아 해당 방으로 전달한다.
type ClientSession struct {
	id   int
	room *room
	send chan []byte
	done chan struct{}
}

func NewClientSession(id int, room *room) *ClientSession {
	clientSession := &ClientSession{
		id:   id,
		room: room,
		send: make(chan []byte),
		done: make(chan struct{}),
	}

	go clientSession.listen()

	return clientSession
}

func (cs *ClientSession) sendMessage(message []byte) {
	cs.send <- message
}

func (cs *ClientSession) listen() {
	for {
		select {
		case message := <-cs.send:
			cs.room.broadcastMessage(message)
		case <-cs.done:
			return
		}
	}
}
