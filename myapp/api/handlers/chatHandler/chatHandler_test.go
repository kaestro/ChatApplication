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

	// Test enter chat
	resp := GetEnterChat(loginInfo)
	if !assert.Equal(t, http.StatusSwitchingProtocols, resp.StatusCode) {
		t.Logf("Failed to reponse on Request to enter chat: %v", resp)
		return
	}

	// Test create room
	resp = PostCreateRoom(roomName, loginSessionID, emailAddress, password)
	if !assert.Equal(t, http.StatusOK, resp.StatusCode) {
		t.Logf("Failed to reponse on Request to create room: %v", resp)
		return
	}

	// Test enter room
	secondUserName := "tch2User"
	secondEmailAddress := "tch2@example.com"
	secondPassword := "password"
	secondUser := models.NewUser(secondUserName, secondEmailAddress, secondPassword)
	secondLoginInfo := models.NewLoginInfo(secondEmailAddress, secondPassword, loginSessionID)

	userService.CreateUser(secondUser)
	secondLoginInfo, err = userServiceUtil.AuthenticateUser(secondLoginInfo, loginSessionID)
	if err != nil {
		t.Fatalf("Failed to authenticate user: %v", err)
	}

	secondLoginSessionID := secondLoginInfo.LoginSessionID

	GetEnterChat(secondLoginInfo)

	resp = PostEnterRoom(roomName, secondLoginSessionID, secondEmailAddress, secondPassword)
	if !assert.Equal(t, http.StatusOK, resp.StatusCode) {
		t.Logf("Failed to reponse on Request to enter room: %v", resp)
		return
	}

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
