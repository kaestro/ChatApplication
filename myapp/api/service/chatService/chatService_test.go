// myapp/api/service/chatService/chatService_test.go
package chatService

import (
	"myapp/api/models"
	"myapp/api/service/userService"
	"myapp/internal/chat"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func TestValidateUpgradeHeader(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/", nil)
	c.Request.Header.Set("Upgrade", "websocket")

	err := ValidateUpgradeHeader(c)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
		return
	}

	t.Logf("Passed test for ValidateUpgradeHeader with valid Upgrade header")
}

func TestParseAndAuthenticateRequest(t *testing.T) {
	tparRoomName := "123"
	tparLoginSessionID := "456"
	tparEmailAddress := "tpar@gmail.com"
	tparPassword := "password"

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/", strings.NewReader(`{"roomName": "123", "loginSessionID": "456", "emailAddress": "tpar@gmail.com", "password": "password"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	user := models.NewUser(tparEmailAddress, tparEmailAddress, tparPassword)
	userService.CreateUser(user, c)
	req, err := ParseAndAuthenticateRequest(c)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
		return
	}

	if req.RoomName != tparRoomName || req.LoginSessionID != tparLoginSessionID || req.EmailAddress != tparEmailAddress {
		t.Errorf("Expected RoomRequest with RoomName 123, LoginSessionID 456, and EmailAddress test@example.com, but got %v", req)
		return
	}

	t.Logf("Passed test for ParseAndAuthenticateRequest with valid JSON")
}

func TestEnterChatRoom(t *testing.T) {
	// 웹소켓 서버를 시작합니다.
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cm := chat.GetChatManager()
		err := cm.ProvideClientToUser(w, r, "testECRSession")
		if err != nil {
			t.Errorf("Failed to upgrade to websocket: %v", err)
			return
		}
	}))
	defer s.Close()

	// 웹소켓 클라이언트를 생성합니다.
	u := "ws" + strings.TrimPrefix(s.URL, "http")
	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Errorf("Failed to open websocket connection: %v", err)
		return
	}
	defer ws.Close()

	// 채팅방을 생성하고, 클라이언트를 채팅방에 입장시킵니다.
	cm := chat.GetChatManager()
	err = cm.CreateRoom("testECRRoom")
	if err != nil {
		t.Errorf("Failed to create room: %v", err)
		return
	}

	err = cm.ClientEnterRoom("testECRRoom", "testECRSession")
	if err != nil {
		t.Errorf("Failed to enter room: %v", err)
		return
	}

	ws.WriteMessage(websocket.TextMessage, []byte("testECRMessage"))

	t.Logf("Passed test for EnterChatRoom with valid websocket connection")
}
