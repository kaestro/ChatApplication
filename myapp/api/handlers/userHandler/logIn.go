// myapp/api/handlers/user/logIn.go
package userHandler

import (
	"myapp/api/service/generalService"
	"myapp/api/service/userService"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Request Type: POST
// Headers: Session-Key
// Body: LoginInfo { emailAddress, password }
func LogIn(ginContext *gin.Context) {
	loginInfo, err := generalService.GetLoginInfoFromBody(ginContext)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userSessionKey := generalService.GetSessionKeyFromHeader(ginContext)
	userServiceUtil := userService.NewUserServiceUtil()
	_, isLoggedIn := userServiceUtil.CheckUserLoggedIn(userSessionKey, loginInfo)
	if isLoggedIn {
		ginContext.JSON(http.StatusOK, gin.H{"message": "Already logged in"})
		return
	}

	loginInfo, err = userServiceUtil.AuthenticateUserByLoginInfo(loginInfo, userSessionKey)
	if err != nil {
		userServiceUtil.HandleLoginError(ginContext, err)
		return
	}

	var sessionKey string
	if loginInfo.LoginSessionID == "" {
		sessionKey, err := userServiceUtil.GenerateSessionKey(loginInfo)
		if err != nil {
			ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		loginInfo.LoginSessionID = sessionKey
	} else {
		sessionKey = loginInfo.LoginSessionID
	}

	ginContext.JSON(http.StatusOK, gin.H{"sessionKey": sessionKey})
}
