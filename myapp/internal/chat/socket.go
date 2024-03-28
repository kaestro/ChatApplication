// myapp/internal/chat/socket.go
// 이거 이름 근데 socket이 맞을까? 이거는 client랑 room을 연결하는 역할을 하는데
// 맞는 이름 고민해 봐야 할듯?
package chat

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// TODO
// session id에 해당하는 client가 이미 있는지 확인하고, 있으면 그 client에다가 room 추가해주고
// 없으면 client 새로 만들고 이 client 저장할 것 만들고.
// 그럼 이 client는 이제 메모리 상에 저장해야한다.
// 그럼 이제 이 client들 크기가 얼마 되는지 고려한 코드 작성 해야할듯?
func connect(w http.ResponseWriter, r *http.Request, room *Room, client *Client, sessionID string) (*Client, error) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	// Register the client to the room
	room.register <- client

	return client, nil
}

// client를 연결에서 끊는다.
func disconnect(client *Client) {
	client.room.unregister <- client

	client.conn.Close()
}
