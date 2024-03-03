package db_test

import (
	"testing"
	"time"

	"myapp/api/models"
)

func TestAddMessage(t *testing.T) {
	manager := setupDB()

	// 메시지 추가
	message := models.Message{
		ChatroomID: "testChatroom",
		Sender:     "testUser",
		Content:    "testMessage",
		Timestamp:  time.Now().Unix(),
	}
	err := manager.AddMessage(message)
	if err != nil {
		t.Fatalf("Failed to add message: %v", err)
	}

	// 메시지 확인
	messages, err := manager.GetMessages("testChatroom")
	if err != nil {
		t.Fatalf("Failed to get messages: %v", err)
	}

	found := false
	for _, m := range messages {
		if m.Content == "testMessage" {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Message 'testMessage' not found")
	}

	t.Log("Chat Message added successfully")
}
