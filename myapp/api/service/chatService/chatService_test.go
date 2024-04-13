// myapp/api/service/userService/chatService_test.go
package chatService

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

func TestChatService(t *testing.T) {
	t.Run("TestPublishWebSocket", func(t *testing.T) {
		// 웹소켓 서버 시작
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := PublishWebSocket(w, r, "testSessionKey")
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

}
