// myapp/internal/models/chat.go
package models

type Chatroom struct {
	ID      string
	Members []string
}

type Message struct {
	ChatroomID string
	Sender     string
	Content    string
	Timestamp  int64
}
