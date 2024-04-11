// myapp/internal/chat/chatManager_test.go
// myapp/internal/chat/chatManager_test.go
package chat

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestChatManager(t *testing.T) {
	cm := NewChatManager()

	// Start a test server
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loginSessionID := r.URL.Query().Get("sessionID")
		err := cm.ProvideClientToUser(w, r, loginSessionID)
		assert.Nil(t, err)
	}))
	defer s.Close()

	// Create a new websocket.Dialer
	dialer := websocket.Dialer{}

	// Send multiple requests
	for i := 0; i < 10; i++ {
		// Create a new websocket connection
		conn, resp, err := dialer.Dial("ws"+s.URL[4:]+"?sessionID="+strconv.Itoa(i), nil)
		assert.Nil(t, err)
		if !assert.Equal(t, http.StatusSwitchingProtocols, resp.StatusCode) {
			t.Logf("Response status code: %d", resp.StatusCode)
			return
		}

		_ = conn.WriteMessage(websocket.TextMessage, []byte("Hello, World!"))

		// Defer the closing of the connection
		if conn != nil {
			conn.Close()
		}
	}

	t.Logf("TestChatManager passed")
}
