package tests

import (
	"testing"

	"myapp/api/models"
	"myapp/internal/db"
)

func TestDBManager(t *testing.T) {
	manager := db.GetDBManager()

	// Create
	user := models.User{
		UserName:     "testuser",
		EmailAddress: "testuser@example.com",
		Password:     "password",
	}
	if err := manager.Create(&user); err != nil {
		t.Errorf("Error creating user: %v", err)
	}

	// Read
	var readUser models.User
	if err := manager.Read(&readUser, "email_address", user.EmailAddress); err != nil {
		t.Errorf("Error reading user: %v", err)
	} else if readUser.EmailAddress != user.EmailAddress {
		t.Errorf("Read user does not match created user")
	}

	// Update
	user.Password = "newpassword"
	if err := manager.Update(&user); err != nil {
		t.Errorf("Error updating user: %v", err)
	}

	// Delete
	if err := manager.Delete(&user); err != nil {
		t.Errorf("Error deleting user: %v", err)
	}

	t.Log("DBManager tests passed")
}
