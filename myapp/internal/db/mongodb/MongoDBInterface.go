// myapp/internal/db/mongodb/MongoDBInterface.go
package mongodb

import (
	"context"

	"myapp/api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBInterface interface {
	InsertChatroom(ctx context.Context, chatroom models.Chatroom) (*mongo.InsertOneResult, error)
	InsertMessage(ctx context.Context, message models.Message) (*mongo.InsertOneResult, error)
	FindMessages(ctx context.Context, filter bson.M) (*mongo.Cursor, error)
	FindChatrooms(ctx context.Context, filter bson.M) (*mongo.Cursor, error)
}
