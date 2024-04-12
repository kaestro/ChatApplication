// myapp/internal/chat/models.go
package chat

import (
	"encoding/json"
	"fmt"
)

var (
	roomIDNeededError   = "roomID is required"
	userNameNeededError = "userName is required"
	contentNeededError  = "content is required"
)

type ChatMessage struct {
	RoomName string `json:"roomName"`
	UserName string `json:"userName"`
	Content  string `json:"content"`
}

func NewChatMessage(roomName, userName, content string) *ChatMessage {
	return &ChatMessage{
		RoomName: roomName,
		UserName: userName,
		Content:  content,
	}
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

	if chatMessage.RoomName == "" {
		return nil, error(fmt.Errorf(roomIDNeededError))
	} else if chatMessage.UserName == "" {
		return nil, error(fmt.Errorf(userNameNeededError))
	} else if chatMessage.Content == "" {
		return nil, error(fmt.Errorf(contentNeededError))
	}

	return chatMessage, nil
}
