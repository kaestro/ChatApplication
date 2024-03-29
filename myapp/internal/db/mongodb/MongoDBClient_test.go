package mongodb

import (
	"context"
	"os"
	"testing"
)

const testMongoURL = "mongodb://localhost:27017"

func TestNewMongoDBClient_Success(t *testing.T) {
	_ = os.Setenv("MONGO_URL", testMongoURL)

	client, err := NewMongoDBClient()

	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	if client == nil {
		t.Fatal("Expected client to be not nil")
	}

	err = client.client.Ping(context.Background(), nil)
	if err != nil {
		t.Fatalf("Failed to ping MongoDB: %v", err)
	}
}
