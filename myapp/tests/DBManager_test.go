package db_test

import (
	"testing"

	"myapp/internal/db"
)

func TestGetDB(t *testing.T) {
	dbInstance := db.GetDB()
	if dbInstance == nil {
		t.Error("Expected db instance to be created, but it was nil")
	} else {
		t.Log("Successfully created db instance")
	}
}
