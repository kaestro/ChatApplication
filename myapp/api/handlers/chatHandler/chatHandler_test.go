// myapp/api/handlers/chatHandler/chatHandler_test.go
package chatHandler

import (
	"bytes"
	"encoding/json"
	"myapp/api/models"
	"myapp/api/service/chatService"
	"myapp/api/service/userService"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TODO: 딴놈 새로 계정 만들어서 enter room 한번 더 테스트
// 현재는 동일한 놈이 두번 접근중
func TestChatHandler(t *testing.T) {
	emailAddress := "tec@example.com"
	password := "password"
	loginSessionID := "testSessionID"
	roomName := "tecRoom"
	userName := "tecUser"

	router := gin.Default()
	router.GET("/enterChat", EnterChat)
	router.POST("/createRoom", CreateRoom)
	router.POST("/enterRoom", EnterRoom)

	loginInfo := models.NewLoginInfo(emailAddress, password, loginSessionID)
	user := models.NewUser(userName, emailAddress, password)

	userService.CreateUser(user)
	userServiceUtil := userService.NewUserServiceUtil()
	loginInfo, err := userServiceUtil.AuthenticateUser(loginInfo, loginSessionID)
	if err != nil {
		t.Fatalf("Failed to authenticate user: %v", err)
	}

	userServiceUtil = userService.NewUserServiceUtil()
	loginInfo, err = userServiceUtil.AuthenticateUser(loginInfo, loginSessionID)
	if err != nil {
		t.Fatalf("Failed to authenticate user: %v", err)
	}
	loginSessionID = loginInfo.LoginSessionID

	go func() {
		router.Run(":8085")
	}()

	// Give the server a second to start
	time.Sleep(time.Second * 3)

	resp := GetEnterChat(loginInfo)

	assert.Equal(t, http.StatusSwitchingProtocols, resp.StatusCode)

	resp = PostCreateRoom(roomName, loginSessionID, emailAddress, password)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	resp = PostEnterRoom(roomName, loginSessionID, emailAddress, password)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	userService.DeleteUserByEmailAddress(emailAddress)
}

func GetEnterChat(loginInfo models.LoginInfo) *http.Response {
	socketKey, _ := chatService.GenerateRandomSocketKey()

	loginInfoBytes, _ := json.Marshal(loginInfo)
	req, _ := http.NewRequest("GET", "http://localhost:8085/enterChat", bytes.NewBuffer(loginInfoBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Upgrade", "websocket")
	req.Header.Set("Connection", "Upgrade")
	req.Header.Set("Sec-WebSocket-Key", socketKey)
	req.Header.Set("Sec-WebSocket-Version", "13")
	req.Header.Set("Session-Key", loginInfo.LoginSessionID)

	client := &http.Client{}
	resp, _ := client.Do(req)
	return resp
}

func PostCreateRoom(roomName string, loginSessionID string, emailAddress string, password string) *http.Response {
	createRoomRequest := models.NewRoomRequest(roomName, loginSessionID, emailAddress, password)
	roomRequestBytes, _ := json.Marshal(createRoomRequest)

	req, _ := http.NewRequest("POST", "http://localhost:8085/createRoom", bytes.NewBuffer(roomRequestBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Session-Key", loginSessionID)

	client := &http.Client{}
	resp, _ := client.Do(req)
	return resp
}

func PostEnterRoom(roomName string, loginSessionID string, emailAddress string, password string) *http.Response {
	enterRoomRequest := models.NewRoomRequest(roomName, loginSessionID, emailAddress, password)
	roomRequestBytes, _ := json.Marshal(enterRoomRequest)

	req, _ := http.NewRequest("POST", "http://localhost:8085/enterRoom", bytes.NewBuffer(roomRequestBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Session-Key", loginSessionID)

	client := &http.Client{}
	resp, _ := client.Do(req)
	return resp
}
