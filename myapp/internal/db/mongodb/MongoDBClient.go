// myapp/internal/db/mongodb/MongoDBClient.go
package mongodb

import (
	"context"
	"time"

	"myapp/api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbURI               = "mongodb://localhost:27017"
	dbName              = "test"
	chatroomsCollection = "chatrooms"
	messagesCollection  = "messages"
)

type MongoDBClient struct {
	client *mongo.Client
}

func NewMongoDBClient() (*MongoDBClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		return nil, err
	}

	return &MongoDBClient{client: client}, nil
}

func (c *MongoDBClient) InsertChatroom(ctx context.Context, chatroom models.Chatroom) (*mongo.InsertOneResult, error) {
	collection := c.client.Database(dbName).Collection(chatroomsCollection)
	return collection.InsertOne(ctx, chatroom)
}

func (c *MongoDBClient) InsertMessage(ctx context.Context, message models.Message) (*mongo.InsertOneResult, error) {
	collection := c.client.Database(dbName).Collection(messagesCollection)
	return collection.InsertOne(ctx, message)
}

func (c *MongoDBClient) FindMessages(ctx context.Context, filter bson.M) (*mongo.Cursor, error) {
	collection := c.client.Database(dbName).Collection(messagesCollection)
	return collection.Find(ctx, filter)
}

func (c *MongoDBClient) FindChatrooms(ctx context.Context, filter bson.M) (*mongo.Cursor, error) {
	collection := c.client.Database(dbName).Collection(chatroomsCollection)
	return collection.Find(ctx, filter)
}
