// myapp/internal/chat/socket.go
// 이거 이름 근데 socket이 맞을까? 이거는 client랑 room을 연결하는 역할을 하는데
// 맞는 이름 고민해 봐야 할듯?
package chat

// TODO
// session id에 해당하는 client가 이미 있는지 확인하고, 있으면 그 client에다가 room 추가해주고
// 없으면 client 새로 만들고 이 client 저장할 것 만들고.
// 그럼 이 client는 이제 메모리 상에 저장해야한다.
// 그럼 이제 이 client들 크기가 얼마 되는지 고려한 코드 작성 해야할듯?
func connect(room *Room, client *Client, sessionID string) (*Client, error) {
	// Upgrade initial GET request to a websocket
	// TODO: add another call Add Connection method implemented inside Client
	// Register the client to the room
	room.register <- client

	return client, nil
}

// client를 연결에서 끊는다.
func disconnect(client *Client, room *Room) {
	// TODO: implement the method inside Client class and call it here
	// Room object will also have to remove the client from the list if room also has the client pointer obj
}
