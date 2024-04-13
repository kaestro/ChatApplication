// myapp/api/service/generalService.go
package service

import (
	"encoding/json"
	"myapp/api/models"

	"github.com/gin-gonic/gin"
)

func ParseLoginInfo(c *gin.Context) (models.LoginInfo, error) {
	var loginInfo models.LoginInfo
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		return models.LoginInfo{}, err
	}
	return loginInfo, nil // Parse login info
}

func GetSessionKeyFromHeader(c *gin.Context) string {
	userSessionKey := c.GetHeader("Session-Key")
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
