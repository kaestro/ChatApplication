// myapp/api/handlers/chatHandler/chatHandler_test.go
package chatHandler

import (
	"bytes"
	"encoding/json"
	"io"
	"myapp/api/models"
	"myapp/api/service/userService"
	"myapp/internal/chat"
	"myapp/jsonProperties"
	"myapp/types"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

var (
	firstRoomName = "tecRoom"
	roomNames     = []string{"room1", "room2", "room3"}

	firstUserEmailAddress = "tec@example.com"
	firstUserPassword     = "password"
	firstUserSessionID    = "tecSessionID"
	firstUserName         = "tecUser"

	secondUserName      = "tch2User"
	secondEmailAddress  = "tch2@example.com"
	secondPassword      = "password"
	secondUserSessionID = "tch2SessionID"

	jsonResponseRoomNamesKey = "roomNames"

	firstChatMessageContent = "Hello, World"
)

func TestChatHandler(t *testing.T) {
	router := gin.Default()
	router.GET("/enterChat", EnterChat)
	router.POST("/createRoom", CreateRoom)
	router.POST("/enterRoom", EnterRoom)
	router.GET("/getRoomList", GetRoomList)
	router.POST("/sendMessage", SendMessage)

	firstLoginInfo := models.NewLoginInfo(firstUserEmailAddress, firstUserPassword, firstUserSessionID)
	firstUser := models.NewUser(firstUserName, firstUserEmailAddress, firstUserPassword)

	userService.CreateUser(firstUser)
	userServiceUtil := userService.NewUserServiceUtil()
	firstLoginInfo, err := userServiceUtil.AuthenticateUserByLoginInfo(firstLoginInfo, firstUserSessionID)
	if err != nil {
		t.Fatalf("Failed to authenticate user: %v", err)
	}

	userServiceUtil = userService.NewUserServiceUtil()
	firstLoginInfo, err = userServiceUtil.AuthenticateUserByLoginInfo(firstLoginInfo, firstUserSessionID)
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
	firstUserConn, resp, err := GETEnterChat(firstLoginInfo)

	if err != nil {
		t.Fatalf("Failed to enter chat: %v", err)
	}

	if !assert.Equal(t, http.StatusSwitchingProtocols, resp.StatusCode) {
		t.Logf("Failed to reponse on Request to enter chat: %v", resp)
		return
	}

	defer firstUserConn.Close()

	// Test create room
	resp = POSTCreateRoom(firstRoomName, firstUserSessionID, firstUserEmailAddress, firstUserPassword)
	if !assert.Equal(t, http.StatusOK, resp.StatusCode) {
		t.Logf("Failed to reponse on Request to create room: %v", resp)
		return
	}

	// Test enter room
	secondUser := models.NewUser(secondUserName, secondEmailAddress, secondPassword)
	secondLoginInfo := models.NewLoginInfo(secondEmailAddress, secondPassword, firstUserSessionID)

	userService.CreateUser(secondUser)
	secondLoginInfo, err = userServiceUtil.AuthenticateUserByLoginInfo(secondLoginInfo, firstUserSessionID)
	if err != nil {
		t.Fatalf("Failed to authenticate user: %v", err)
	}

	secondUserSessionID = secondLoginInfo.LoginSessionID

	secondUserConn, _, err := GETEnterChat(secondLoginInfo)

	if err != nil {
		t.Fatalf("Failed to enter chat: %v", err)
	}

	defer secondUserConn.Close()

	resp = POSTEnterRoom(firstRoomName, secondUserSessionID, secondEmailAddress, secondPassword)
	if !assert.Equal(t, http.StatusOK, resp.StatusCode) {
		t.Logf("Failed to reponse on Request to enter room: %v", resp)
		return
	}

	// Test get room list
	for _, roomName := range roomNames {
		POSTCreateRoom(roomName, firstUserSessionID, firstUserEmailAddress, firstUserPassword)
	}

	resp = GETGetRoomList(firstUserSessionID, firstUserEmailAddress, firstUserPassword)

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

	// Test send message
	roomRequest := models.NewRoomRequest(firstRoomName, "", firstUserEmailAddress, firstUserPassword)
	chatMessage := chat.NewChatMessage(firstRoomName, firstUserName, firstChatMessageContent)
	resp = POSTSendMessage(chatMessage, roomRequest, types.LoginSessionID(firstUserSessionID))

	if !assert.Equal(t, http.StatusOK, resp.StatusCode) {
		t.Logf("Failed to reponse on Request to send message: %v", resp)
		return
	}

	// Test receive message
	secondUserConn.SetReadDeadline(time.Now().Add(time.Second * 2)) // Timeout after 2 seconds
	_, message, err := secondUserConn.ReadMessage()
	if err != nil {
		t.Fatalf("Failed to read message from WebSocket: %v", err)
	}

	readChatMessage, err := chat.NewChatMessageFromBytes(message)
	if err != nil {
		t.Fatalf("Failed to unmarshal chat message: %v", err)
	}

	if !assert.Equal(t, firstRoomName, readChatMessage.RoomName) {
		t.Logf("Failed to receive the correct room name")
		return
	}

	if !assert.Equal(t, firstUserName, readChatMessage.UserName) {
		t.Logf("Failed to receive the correct user name")
		return
	}

	if !assert.Equal(t, firstChatMessageContent, readChatMessage.Content) {
		t.Logf("Failed to receive the correct message")
		return
	}

	finish()
	t.Logf("ChatHandler test passed")
}

func GETEnterChat(loginInfo models.LoginInfo) (*websocket.Conn, *http.Response, error) {
	// socketKey, _ := chatService.GenerateRandomSocketKey()

	loginInfoBytes, _ := json.Marshal(loginInfo)
	req, _ := http.NewRequest("GET", "ws://localhost:8085/enterChat", bytes.NewBuffer(loginInfoBytes))
	req.Header.Set("Content-Type", "application/json")

	// handshake header는 dialer가 해결해준다.
	req.Header.Set("Session-Key", loginInfo.LoginSessionID)
	req.Header.Set(jsonProperties.EmailAddress, loginInfo.EmailAddress)

	dialer := &websocket.Dialer{}
	conn, resp, err := dialer.Dial(req.URL.String(), req.Header)

	return conn, resp, err
}

func POSTCreateRoom(roomName string, loginSessionID string, emailAddress string, password string) *http.Response {
	createRoomRequest := models.NewRoomRequest(roomName, loginSessionID, emailAddress, password)
	roomRequestBytes, _ := json.Marshal(createRoomRequest)

	req, _ := http.NewRequest("POST", "http://localhost:8085/createRoom", bytes.NewBuffer(roomRequestBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Session-Key", loginSessionID)

	client := &http.Client{}
	resp, _ := client.Do(req)
	return resp
}

func POSTEnterRoom(roomName string, loginSessionID string, emailAddress string, password string) *http.Response {
	enterRoomRequest := models.NewRoomRequest(roomName, loginSessionID, emailAddress, password)
	roomRequestBytes, _ := json.Marshal(enterRoomRequest)

	req, _ := http.NewRequest("POST", "http://localhost:8085/enterRoom", bytes.NewBuffer(roomRequestBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Session-Key", loginSessionID)

	client := &http.Client{}
	resp, _ := client.Do(req)
	return resp
}

func GETGetRoomList(loginSessionID string, emailAddress string, password string) *http.Response {
	loginInfo := models.NewLoginInfo(emailAddress, password, "")
	loginInfoBytes, _ := json.Marshal(loginInfo)

	req, _ := http.NewRequest("GET", "http://localhost:8085/getRoomList", bytes.NewBuffer(loginInfoBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Session-Key", loginSessionID)

	client := &http.Client{}
	resp, _ := client.Do(req)

	return resp
}

func POSTSendMessage(chatMessage *chat.ChatMessage, roomRequest *models.RoomRequest, loginSessionID types.LoginSessionID) *http.Response {
	chatMessageRequest := models.NewChatMessageRequest(roomRequest, chatMessage)
	chatMessageRequestBytes, _ := json.Marshal(chatMessageRequest)

	req, _ := http.NewRequest("POST", "http://localhost:8085/sendMessage", bytes.NewBuffer(chatMessageRequestBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(jsonProperties.SessionKey, string(loginSessionID))

	client := &http.Client{}
	resp, _ := client.Do(req)

	return resp
}

func finish() {
	userService.DeleteUserByEmailAddress(firstUserEmailAddress)
	userService.DeleteUserByEmailAddress(secondEmailAddress)
}
