// myapp/api/handlers/user/logIn.go
package userHandler

import (
	"encoding/json"
	"myapp/api/models"
	"myapp/api/service/userService"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LogIn 함수는 사용자가 제공한 이메일 주소와 비밀번호를 검증하여 로그인합니다.
func LogIn(ginContext *gin.Context) {
	loginInfo, err := getLoginInfo(ginContext)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userSessionKey := ginContext.GetHeader("Session-Key")
	userServiceUtil := userService.NewUserServiceUtil()

	user, err := userServiceUtil.AuthenticateUser(loginInfo, userSessionKey)
	if err != nil {
		userServiceUtil.HandleLoginError(ginContext, err)
		return
	}

	sessionKey, err := userServiceUtil.GenerateSessionKey(user)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusOK, gin.H{"sessionKey": sessionKey})
}

func getLoginInfo(ginContext *gin.Context) (models.LoginInfo, error) {
	var loginInfo models.LoginInfo
	err := json.NewDecoder(ginContext.Request.Body).Decode(&loginInfo)
	return loginInfo, err
}
