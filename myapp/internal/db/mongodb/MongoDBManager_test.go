// myapp/internal/db/mongodb/MongoDBManager_test.go
package mongodb

import (
	"context"
	"testing"

	"myapp/api/models"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockMongoDBInterface struct {
	mock.Mock
}

func (m *MockMongoDBInterface) InsertChatroom(ctx context.Context, chatroom models.Chatroom) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, chatroom)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockMongoDBInterface) InsertMessage(ctx context.Context, message models.Message) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, message)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockMongoDBInterface) FindMessages(ctx context.Context, filter bson.M) (*mongo.Cursor, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.Cursor), args.Error(1)
}

func (m *MockMongoDBInterface) FindChatrooms(ctx context.Context, filter bson.M) (*mongo.Cursor, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.Cursor), args.Error(1)
}

func TestMongoDBManager(t *testing.T) {
	mockDB := new(MockMongoDBInterface)
	manager := NewMongoDBManager(mockDB)

	chatroom := models.Chatroom{
		ID:      "123",
		Members: []string{"user1", "user2"},
	}
	message := models.Message{
		ChatroomID: "123",
		Sender:     "user1",
		Content:    "Test Message",
		Timestamp:  1633029442,
	}
	filter := bson.M{"chatroom_id": "123"}

	mockDB.On("InsertChatroom", mock.Anything, chatroom).Return(&mongo.InsertOneResult{}, nil)
	mockDB.On("InsertMessage", mock.Anything, message).Return(&mongo.InsertOneResult{}, nil)
	mockDB.On("FindMessages", mock.Anything, filter).Return(&mongo.Cursor{}, nil)
	mockDB.On("FindChatrooms", mock.Anything, filter).Return(&mongo.Cursor{}, nil)

	_, err := manager.InsertChatroom(context.Background(), chatroom)
	require.NoError(t, err)

	_, err = manager.InsertMessage(context.Background(), message)
	require.NoError(t, err)

	_, err = manager.FindMessages(context.Background(), filter)
	require.NoError(t, err)

	_, err = manager.FindChatrooms(context.Background(), filter)
	require.NoError(t, err)

	mockDB.AssertExpectations(t)
}
