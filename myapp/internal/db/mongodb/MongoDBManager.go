// myapp/internal/db/mongodb/MongoDBManager.go
package mongodb

import (
	"context"

	"myapp/api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBManager struct {
	db MongoDBInterface
}

func NewMongoDBManager(db MongoDBInterface) *MongoDBManager {
	return &MongoDBManager{
		db: db,
	}
}

func (m *MongoDBManager) InsertChatroom(ctx context.Context, chatroom models.Chatroom) (*mongo.InsertOneResult, error) {
	return m.db.InsertChatroom(ctx, chatroom)
}

func (m *MongoDBManager) InsertMessage(ctx context.Context, message models.Message) (*mongo.InsertOneResult, error) {
	return m.db.InsertMessage(ctx, message)
}

func (m *MongoDBManager) FindMessages(ctx context.Context, filter bson.M) (*mongo.Cursor, error) {
	return m.db.FindMessages(ctx, filter)
}

func (m *MongoDBManager) FindChatrooms(ctx context.Context, filter bson.M) (*mongo.Cursor, error) {
	return m.db.FindChatrooms(ctx, filter)
}
