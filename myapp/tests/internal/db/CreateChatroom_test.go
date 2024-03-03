package db_test

import (
	"testing"

	"myapp/api/models"
	"myapp/internal/db/mongodb"
)

func setupDB() *mongodb.ChatDBManager {
	return mongodb.GetChatDBManager()
}

func TestCreateChatroom(t *testing.T) {
	manager := setupDB()

	// chatroom 생성
	chatroom := models.Chatroom{
		ID:      "testChatroom",
		Members: []string{"testUser"},
	}
	err := manager.CreateChatroom(chatroom)
	if err != nil {
		t.Fatalf("Failed to create chatroom: %v", err)
	}

	// chatroom 확인
	chatrooms, err := manager.GetChatrooms("testUser")
	if err != nil {
		t.Fatalf("Failed to get chatrooms: %v", err)
	}

	found := false
	for _, c := range chatrooms {
		if c.ID == "testChatroom" {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Chatroom 'testChatroom' not found")
	}

	t.Log("Chatroom created successfully")
}
