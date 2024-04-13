// myapp/api/service/userService/chatService_test.go
package chatService

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func TestPublishAndCheckConnection(t *testing.T) {
	t.Run("TestPublishWebSocket", func(t *testing.T) {
		testSessionKey := "tpwsSessionKey"
		// 웹소켓 서버 시작
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := PublishWebSocket(w, r, testSessionKey)
			if err != nil {
				t.Errorf("Expected no error, but got %v", err)
			}
		}))
		defer server.Close()

		// 웹소켓 클라이언트로 연결 시도
		_, _, err := websocket.DefaultDialer.Dial(strings.Replace(server.URL, "http", "ws", 1), nil)
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}
	})

	t.Run("TestCheckSocketConnection", func(t *testing.T) {
		testSessionKey := "tcscSessionKey"
		// 웹소켓 서버 시작
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := PublishWebSocket(w, r, testSessionKey)
			if err != nil {
				t.Errorf("Expected error, but got %v", err)
			}
		}))
		defer server.Close()

		// 웹소켓 클라이언트로 연결 시도
		_, _, err := websocket.DefaultDialer.Dial(strings.Replace(server.URL, "http", "ws", 1), nil)
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}

		// CheckSocketConnection 함수를 사용하여 연결 확인
		err = CheckSocketConnection(testSessionKey)
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}
	})
}

func TestIsUpgradeHeaderValid(t *testing.T) {
	t.Run("returns true and 200 OK when Upgrade header is websocket", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("GET", "/", nil) // Add this line
		c.Request.Header.Set("Upgrade", "websocket")

		result := IsUpgradeHeaderValid(c)

		if !result || w.Code != http.StatusOK {
			t.Errorf("Expected true and 200 OK, but got %v and %d", result, w.Code)
		}
	})

	t.Run("returns false and 400 Bad Request when Upgrade header is not websocket", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("GET", "/", nil) // Add this line
		c.Request.Header.Set("Upgrade", "not-websocket")

		result := IsUpgradeHeaderValid(c)

		if result || w.Code != http.StatusBadRequest {
			t.Errorf("Expected false and 400 Bad Request, but got %v and %d", result, w.Code)
		}
	})
}
