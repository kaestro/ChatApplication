// myapp/api/service/chatService/chatService_test.go
package chatService

import (
	"crypto/rand"
	"encoding/base64"
	"myapp/api/models"
	"myapp/api/service/userService"
	"myapp/internal/chat"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestValidateUpgradeHeader(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/", nil)
	c.Request.Header.Set("Upgrade", "websocket")
	c.Request.Header.Set("Connection", "upgrade")
	c.Request.Header.Set("Sec-WebSocket-Version", "13")
	c.Request.Header.Set("Sec-WebSocket-Key", "test")

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
	userService.DeleteUserByEmailAddress(user.EmailAddress)
}

func TestEnterChatRoom(t *testing.T) {
	sessionKey := "testECRSession"
	roomName := "testECRRoom"
	password := "testECRPassword"
	emailAddress := "testECR@example.com"
	roomRequest := models.NewRoomRequest(roomName, sessionKey, emailAddress, password)
	loginInfo := models.NewLoginInfo(emailAddress, password)
	loginInfo.LoginSessionID = sessionKey

	// gin.Engine을 생성합니다.
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		err := EnterChat(context, loginInfo)
		if err != nil {
			t.Errorf("Failed to enter chat: %v", err)
			return
		}

		err = EnterChatRoom(context, roomRequest)
		if err != nil {
			t.Errorf("Failed to enter chat room: %v", err)
			return
		}

		t.Logf("Passed test for EnterChatRoom with valid websocket connection")
	})

	// httptest.Server를 생성합니다.
	server := httptest.NewServer(router)

	// 16바이트 길이의 랜덤한 바이트를 생성합니다.
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		t.Fatal(err)
	}

	// 채팅방을 생성합니다.
	cm := chat.GetChatManager()
	err = cm.CreateRoom(roomName)
	if err != nil {
		t.Errorf("Failed to create room: %v", err)
		return
	}

	// 웹소켓 핸드셰이크를 위한 헤더를 추가합니다.
	req, err := http.NewRequest("GET", server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Connection", "upgrade")
	req.Header.Add("Upgrade", "websocket")
	req.Header.Add("Sec-WebSocket-Version", "13")
	req.Header.Add("Sec-WebSocket-Key", base64.StdEncoding.EncodeToString(randomBytes))

	// Make a new HTTP request to the server
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

}
