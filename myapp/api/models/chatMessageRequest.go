// myapp/api/models/roomRequest.go
package models

import "myapp/internal/chat"

type ChatMessageRequest struct {
	RoomName       string `json:"roomName"`
	UserName       string `json:"userName"`
	LoginSessionID string `json:"loginSessionID"`
	EmailAddress   string `json:"emailAddress"`
	Password       string `json:"password"`
	Content        string `json:"content"`
}

func NewChatMessageRequest(rr *RoomRequest, cm *chat.ChatMessage) *ChatMessageRequest {
	return &ChatMessageRequest{
		RoomName:       rr.RoomName,
		UserName:       cm.UserName,
		LoginSessionID: rr.LoginSessionID,
		EmailAddress:   rr.EmailAddress,
		Password:       rr.Password,
		Content:        cm.Content,
	}
}
