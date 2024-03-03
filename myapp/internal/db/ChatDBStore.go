// myapp/internal/db/ChatDBStore.go
package db

import (
	"myapp/api/models"
)

type ChatDBStore interface {
	CreateChatroom(chatroom models.Chatroom) error
	AddMessage(message models.Message) error
	GetMessages(chatroomID string) ([]models.Message, error)
	GetChatrooms(userID string) ([]models.Chatroom, error)
}
