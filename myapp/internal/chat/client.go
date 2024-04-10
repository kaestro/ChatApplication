// myapp/internal/chat/models.go
package chat

type Client struct {
	loginSessionID string           // 어느 user인지 구분하는 id
	clientSessions []*ClientSession // room, socket, send channel을 가지고 있는 session slice
}

func NewClient(loginSessionID string) *Client {
	return &Client{
		loginSessionID: loginSessionID,
		clientSessions: make([]*ClientSession, 0),
	}
}

// TODO
// 해당 부분이 모든 함수들에서 중복되게 사용되고 있다. 이를 빼내는 middleware 형태로 변경의 필요
func (c *Client) isSameClient(loginSessionID string) bool {
	return c.loginSessionID == loginSessionID
}

func (c *Client) AddClientSession(conn Conn, room *Room, loginSessionID string) {
	if !c.isSameClient(loginSessionID) {
		return
	}

	c.clientSessions = append(c.clientSessions, NewClientSession(len(c.clientSessions), conn, room))
}

func (c *Client) RemoveClientSession(clientSessionID int, loginSessionID string) {
	if !c.isSameClient(loginSessionID) {
		return
	}

	for i, clientSession := range c.clientSessions {
		if clientSession.id == clientSessionID {
			c.clientSessions = append(c.clientSessions[:i], c.clientSessions[i+1:]...)
			break
		}
	}
}

func (c *Client) GetLoginSessionID() string {
	return c.loginSessionID
}
