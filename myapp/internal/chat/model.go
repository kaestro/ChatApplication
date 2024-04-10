// myapp/internal/chat/models.go
package chat

type ChatMessage struct {
	RoomID   string `json:"roomID"`
	UserName string `json:"userName"`
	Content  string `json:"content"`
}
