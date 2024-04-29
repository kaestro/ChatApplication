// myapp/api/service/generalService.go
package generalService

import (
	"bytes"
	"encoding/json"
	"io"
	"myapp/api/models"
	"myapp/internal/chat"

	jsonProperties "myapp/jsonproperties"

	"github.com/gin-gonic/gin"
)

func readAndRestoreBody(c *gin.Context) []byte {
	bodyBytes, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	return bodyBytes
}

func ParseLoginInfo(c *gin.Context) (models.LoginInfo, error) {
	var loginInfo models.LoginInfo
	bodyBytes := readAndRestoreBody(c)
	if err := json.Unmarshal(bodyBytes, &loginInfo); err != nil {
		return models.LoginInfo{}, err
	}

	sessionKey := GetSessionKeyFromHeader(c)
	loginInfo.LoginSessionID = sessionKey

	return loginInfo, nil // Parse login info
}

func GetSessionKeyFromHeader(c *gin.Context) string {
	userSessionKey := c.GetHeader(jsonProperties.SessionKey)
	return userSessionKey
}

func DecodeUserFromBody(ginContext *gin.Context, user *models.User) error {
	err := json.NewDecoder(ginContext.Request.Body).Decode(&user)
	return err
}

func GetLoginInfoFromBody(ginContext *gin.Context) (models.LoginInfo, error) {
	var loginInfo models.LoginInfo
	err := json.NewDecoder(ginContext.Request.Body).Decode(&loginInfo)
	return loginInfo, err
}

// RequestBody에서 RoomRequest를 파싱하고, Header에서 Session-Key를 가져와서 RoomRequest에 넣어준다.
func ParseRoomRequest(c *gin.Context) (models.RoomRequest, error) {
	bodyBytes := readAndRestoreBody(c)
	var req models.RoomRequest
	if err := json.Unmarshal(bodyBytes, &req); err != nil {
		return models.RoomRequest{}, err
	}

	loginSessionID := GetSessionKeyFromHeader(c)
	req.LoginSessionID = loginSessionID

	return req, nil
}

func ParseChatMessage(c *gin.Context) (chat.ChatMessage, error) {
	bodyBytes := readAndRestoreBody(c)
	var chatMessage chat.ChatMessage
	if err := json.Unmarshal(bodyBytes, &chatMessage); err != nil {
		return chat.ChatMessage{}, err
	}

	return chatMessage, nil
}
