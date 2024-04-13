// myapp/api/service/chatService/parseEnterRoomRequest_test.go
package chatService

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestIsUpgradeHeaderValid(t *testing.T) {
	t.Run("returns true and 200 OK when Upgrade header is websocket", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("GET", "/", nil) // Add this line
		c.Request.Header.Set("Upgrade", "websocket")

		isValid := IsUpgradeHeaderValid(c)

		if !isValid {
			t.Error("Expected true but got false")
			return
		}
		t.Logf("Passed test for IsUpgradeHeaderValid with Upgrade header websocket")
	})

	t.Run("returns false and 400 Bad Request when Upgrade header is not websocket", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("GET", "/", nil) // Add this line
		c.Request.Header.Set("Upgrade", "not-websocket")

		isValid := IsUpgradeHeaderValid(c)

		if isValid {
			t.Errorf("Expected false but got %v", isValid)
			return
		}
		t.Logf("Passed test for IsUpgradeHeaderValid with Upgrade header not websocket")
	})
}
func TestParseEnterRoomRequest(t *testing.T) {
	t.Run("returns RoomRequest and no error when JSON is valid", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("POST", "/", strings.NewReader(`{"roomName": "123"}`))
		c.Request.Header.Set("Content-Type", "application/json")

		req, err := ParseEnterRoomRequest(c)

		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
			return
		}

		if req.RoomName != "123" {
			t.Errorf("Expected RoomRequest with RoomName 123, but got %v", req)
			return
		}

		t.Logf("Passed test for ParseEnterRoomRequest with valid JSON")
	})

	t.Run("returns error when JSON is invalid", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("POST", "/", strings.NewReader(`{"invalid": "json"`))
		c.Request.Header.Set("Content-Type", "application/json")

		_, err := ParseEnterRoomRequest(c)

		if err == nil {
			t.Error("Expected error, but got nil")
			return
		}
		t.Logf("Passed test for ParseEnterRoomRequest with invalid JSON")
	})
}
