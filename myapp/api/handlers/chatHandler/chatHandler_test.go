// myapp/api/handlers/chatHandler/chatHandler_test.go
package chatHandler

import (
	"bytes"
	"encoding/json"
	"io"
	"myapp/api/models"
	"myapp/api/service/chatService"
	"myapp/api/service/userService"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	roomNames = []string{"room1", "room2", "room3"}

	firstUserEmailAddress = "tec@example.com"
	firstUserPassword     = "password"
	firstUserSessionID    = "tecSessionID"
	firstRoomName         = "tecRoom"
	firstUserName         = "tecUser"

	secondUserName     = "tch2User"
	secondEmailAddress = "tch2@example.com"
	secondPassword     = "password"

	jsonResponseRoomNamesKey = "roomNames"
)

func TestChatHandler(t *testing.T) {
	router := gin.Default()
	router.GET("/enterChat", EnterChat)
	router.POST("/createRoom", CreateRoom)
	router.POST("/enterRoom", EnterRoom)
	router.GET("/getRoomList", GetRoomList)

	firstLoginInfo := models.NewLoginInfo(firstUserEmailAddress, firstUserPassword, firstUserSessionID)
	firstUser := models.NewUser(firstUserName, firstUserEmailAddress, firstUserPassword)

	userService.CreateUser(firstUser)
	userServiceUtil := userService.NewUserServiceUtil()
	firstLoginInfo, err := userServiceUtil.AuthenticateUser(firstLoginInfo, firstUserSessionID)
	if err != nil {
		t.Fatalf("Failed to authenticate user: %v", err)
	}

	userServiceUtil = userService.NewUserServiceUtil()
	firstLoginInfo, err = userServiceUtil.AuthenticateUser(firstLoginInfo, firstUserSessionID)
	if err != nil {
		t.Fatalf("Failed to authenticate user: %v", err)
	}
	firstUserSessionID = firstLoginInfo.LoginSessionID

	go func() {
		router.Run(":8085")
	}()

	// Give the server a second to start
	time.Sleep(time.Second * 3)

	// Test enter chat
	resp := GetEnterChat(firstLoginInfo)
	if !assert.Equal(t, http.StatusSwitchingProtocols, resp.StatusCode) {
		t.Logf("Failed to reponse on Request to enter chat: %v", resp)
		return
	}

	// Test create room
	resp = PostCreateRoom(firstRoomName, firstUserSessionID, firstUserEmailAddress, firstUserPassword)
	if !assert.Equal(t, http.StatusOK, resp.StatusCode) {
		t.Logf("Failed to reponse on Request to create room: %v", resp)
		return
	}

	// Test enter room
	secondUser := models.NewUser(secondUserName, secondEmailAddress, secondPassword)
	secondLoginInfo := models.NewLoginInfo(secondEmailAddress, secondPassword, firstUserSessionID)

	userService.CreateUser(secondUser)
	secondLoginInfo, err = userServiceUtil.AuthenticateUser(secondLoginInfo, firstUserSessionID)
	if err != nil {
		t.Fatalf("Failed to authenticate user: %v", err)
	}

	secondLoginSessionID := secondLoginInfo.LoginSessionID

	GetEnterChat(secondLoginInfo)

	resp = PostEnterRoom(firstRoomName, secondLoginSessionID, secondEmailAddress, secondPassword)
	if !assert.Equal(t, http.StatusOK, resp.StatusCode) {
		t.Logf("Failed to reponse on Request to enter room: %v", resp)
		return
	}

	// Test get room list
	for _, roomName := range roomNames {
		PostCreateRoom(roomName, firstUserSessionID, firstUserEmailAddress, firstUserPassword)
	}

	resp = GetGetRoomList(firstUserSessionID, firstUserEmailAddress, firstUserPassword)

	if !assert.Equal(t, http.StatusOK, resp.StatusCode) {
		t.Logf("Failed to reponse on Request to get room list: %v", resp)
		return
	}

	// Read the response body
	body, _ := io.ReadAll(resp.Body)
	var returnedRoomNames map[string][]string
	json.Unmarshal(body, &returnedRoomNames)

	roomNamesToCheck := append(roomNames, firstRoomName)

	// Compare the returned room names with the expected room names
	if !assert.ElementsMatch(t, roomNamesToCheck, returnedRoomNames[jsonResponseRoomNamesKey]) {
		t.Logf("Returned room names do not match the expected room names")
		return
	}

	finish()
	t.Logf("ChatHandler test passed")
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

func GetGetRoomList(loginSessionID string, emailAddress string, password string) *http.Response {
	loginInfo := models.NewLoginInfo(emailAddress, password, "")
	loginInfoBytes, _ := json.Marshal(loginInfo)

	req, _ := http.NewRequest("GET", "http://localhost:8085/getRoomList", bytes.NewBuffer(loginInfoBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Session-Key", loginSessionID)

	client := &http.Client{}
	resp, _ := client.Do(req)

	return resp
}

func finish() {
	userService.DeleteUserByEmailAddress(firstUserEmailAddress)
	userService.DeleteUserByEmailAddress(secondEmailAddress)
}
