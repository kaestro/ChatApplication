package db_test

import (
	"testing"

	"github.com/kaestro/ChatApplication/internal/db"
)

func TestGetDB(t *testing.T) {
	conn := db.GetDB()
	err := conn.Ping()
	if err != nil {
		t.Fatalf("Could not connect to the database: %v", err)
	}
}
