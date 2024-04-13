// myapp/api/handlers/chatHandler/enterChat_test.go
package chatHandler

import (
	"bytes"
	"encoding/json"
	"myapp/api/models"
	"myapp/api/service/chatService"
	"myapp/api/service/userService"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestEnterChat(t *testing.T) {
	emailAddress := "tec@example.com"
	password := "password"
	loginSessionID := "testSessionKey"

	router := gin.Default()
	router.GET("/enterChat", EnterChat)

	loginInfo := models.NewLoginInfo(emailAddress, password, loginSessionID)
	user := models.NewUser("testUser", emailAddress, password)

	userService.CreateUser(user)

	socketKey, _ := chatService.GenerateRandomSocketKey()

	loginInfoBytes, _ := json.Marshal(loginInfo)
	req, _ := http.NewRequest("GET", "/enterChat", bytes.NewBuffer(loginInfoBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Upgrade", "websocket")
	req.Header.Set("Connection", "upgrade")
	req.Header.Set("Sec-WebSocket-Version", "13")
	req.Header.Set("Sec-WebSocket-Key", socketKey)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	userService.DeleteUserByEmailAddress(emailAddress)
}
