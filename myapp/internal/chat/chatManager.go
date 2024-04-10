// myapp/internal/chat/chatManager.go
package chat

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// TODO: Design ChatManager
// 채팅 서버의 전반적인 관리를 담당하는 ChatManager
// 채팅과 관련한 모든 요청은 ChatManager를 통해 이루어진다.
type ChatManager struct {
}

// TODO: enterRoom delegated from client to ChatManager
// Question: Do Client Connection actually need conn object?
func (cm *ChatManager) EnterRoom(room *Room) {
	// Check if the client is already in the room
	// if already in, return

	// if not, add another Connection object to the client
	// Question: How to generate a unique id for the connection?
	// Assuming that each clients won't be having many connections, simple slice can be used
}

// TODO
// session id에 해당하는 client가 이미 있는지 확인하고, 있으면 그 client에다가 room 추가해주고
// 없으면 client 새로 만들고 이 client 저장할 것 만들고.
// 그럼 이 client는 이제 메모리 상에 저장해야한다.
// 그럼 이제 이 client들 크기가 얼마 되는지 고려한 코드 작성 해야할듯?
func (cm *ChatManager) Connect(w http.ResponseWriter, r *http.Request, room *Room, client *Client, sessionID string) (*Client, error) {
	// Upgrade initial GET request to a websocket
	// TODO: 이 부분 따로 함수로 꺼내는 refactoring ~ line 31
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	// TODO: add another call Add Connection method implemented inside Client
	// Register the client to the room
	room.AddClient(client, ws)

	return client, nil
}

// client를 연결에서 끊는다.
func (cm *ChatManager) Disconnect(client *Client, room *Room) {
	// TODO: implement the method inside Client class and call it here
	// Room object will also have to remove the client from the list if room also has the client pointer obj
}
