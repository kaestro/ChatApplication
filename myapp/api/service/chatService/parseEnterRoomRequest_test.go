// myapp/api/service/chatService/parseEnterRoomRequest_test.go
package chatService

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"myapp/api/service"

	"github.com/gin-gonic/gin"
)

func TestIsUpgradeHeaderValid(t *testing.T) {
	t.Run("returns true when all headers are valid", func(t *testing.T) {
		socketKey, _ := GenerateRandomSocketKey()
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("GET", "/", nil)
		c.Request.Header.Set("Upgrade", "websocket")
		c.Request.Header.Set("Connection", "upgrade")
		c.Request.Header.Set("Sec-WebSocket-Version", "13")
		c.Request.Header.Set("Sec-WebSocket-Key", socketKey)

		isValid := IsUpgradeHeaderValid(c)

		if !isValid {
			t.Error("Expected true but got false")
			return
		}
		t.Logf("Passed test for IsUpgradeHeaderValid with valid headers")
	})

	t.Run("returns false when any header is invalid", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("GET", "/", nil)
		c.Request.Header.Set("Upgrade", "not-websocket")
		c.Request.Header.Set("Connection", "not-upgrade")
		c.Request.Header.Set("Sec-WebSocket-Version", "not-13")
		c.Request.Header.Set("Sec-WebSocket-Key", "")

		isValid := IsUpgradeHeaderValid(c)

		if isValid {
			t.Errorf("Expected false but got %v", isValid)
			return
		}
		t.Logf("Passed test for IsUpgradeHeaderValid with invalid headers")
	})
}
func TestParseEnterRoomRequest(t *testing.T) {
	t.Run("returns RoomRequest and no error when JSON is valid", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("POST", "/", strings.NewReader(`{"roomName": "123", "emailAddress": "test@example.com"}`))
		c.Request.Header.Set("Content-Type", "application/json")
		c.Request.Header.Set("Session-Key", "456")

		req, err := service.ParseRoomRequest(c)

		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
			return
		}

		if req.RoomName != "123" || req.LoginSessionID != "456" || req.EmailAddress != "test@example.com" {
			t.Errorf("Expected RoomRequest with RoomName 123, LoginSessionID 456, and EmailAddress test@example.com, but got %v", req)
			return
		}

		t.Logf("Passed test for ParseEnterRoomRequest with valid JSON")
	})

	t.Run("returns error when JSON is invalid", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("POST", "/", strings.NewReader(`{"invalid": "json"`))
		c.Request.Header.Set("Content-Type", "application/json")

		_, err := service.ParseRoomRequest(c)

		if err == nil {
			t.Error("Expected error, but got nil")
			return
		}
		t.Logf("Passed test for ParseEnterRoomRequest with invalid JSON")
	})
}
