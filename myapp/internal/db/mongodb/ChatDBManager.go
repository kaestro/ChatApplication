// myapp/internal/db/ChatDBManager.go
package mongodb

import (
	"context"
	"os"
	"sync"

	"myapp/api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ChatDBManager struct {
	client    *mongo.Client
	chatrooms *mongo.Collection
	messages  *mongo.Collection
}

var (
	once sync.Once

	manager *ChatDBManager
)

func GetChatDBManager() *ChatDBManager {
	once.Do(func() {
		var err error
		mongoURL := os.Getenv("MONGO_URL")
		if mongoURL == "" {
			mongoURL = "mongodb://localhost:27017" // default value
		}
		client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURL))
		if err != nil {
			panic(err)
		}

		db := client.Database("chat")
		manager = &ChatDBManager{
			client:    client,
			chatrooms: db.Collection("chatrooms"),
			messages:  db.Collection("messages"),
		}
	})

	return manager
}

func (m *ChatDBManager) CreateChatroom(chatroom models.Chatroom) error {
	_, err := m.chatrooms.InsertOne(context.Background(), chatroom)
	return err
}

func (m *ChatDBManager) AddMessage(message models.Message) error {
	_, err := m.messages.InsertOne(context.Background(), message)
	return err
}

// chatroom에 입장했을 때, chatroom상에 있는 session 모든 데이터를 가져온다.
func (m *ChatDBManager) GetMessages(chatroomID string) ([]models.Message, error) {
	filter := bson.M{"chatroomid": chatroomID}
	cursor, err := m.messages.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var messages []models.Message
	if err = cursor.All(context.Background(), &messages); err != nil {
		return nil, err
	}

	return messages, nil
}

// user가 login했을 때, 접속 가능한 모든 Chatroom을 가져온다.
func (m *ChatDBManager) GetChatrooms(userID string) ([]models.Chatroom, error) {
	filter := bson.M{"members": userID}
	cursor, err := m.chatrooms.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var chatrooms []models.Chatroom
	if err = cursor.All(context.Background(), &chatrooms); err != nil {
		return nil, err
	}

	return chatrooms, nil
}
