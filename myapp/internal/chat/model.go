// myapp/internal/chat/models.go
package chat

import "encoding/json"

type ChatMessage struct {
	RoomID   string `json:"roomID"`
	UserName string `json:"userName"`
	Content  string `json:"content"`
}

func (cm *ChatMessage) ToBytes() ([]byte, error) {
	return json.Marshal(cm)
}

func NewChatMessageFromBytes(data []byte) (*ChatMessage, error) {
	chatMessage := &ChatMessage{}
	err := json.Unmarshal(data, chatMessage)
	if err != nil {
		return nil, err
	}

	return chatMessage, nil
}
