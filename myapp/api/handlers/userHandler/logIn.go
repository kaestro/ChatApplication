// myapp/api/handlers/user/logIn.go
package userHandler

import (
	"log"
	"myapp/api/models"
	"myapp/api/service/generalService"
	"myapp/api/service/userService"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Request Type: POST
// Headers: Session-Key
// Body: LoginInfo { emailAddress, password }
func LogIn(ginContext *gin.Context) {
	loginInfo, err := getLoginInfo(ginContext)
	if err != nil {
		ginContext.Errors = append(ginContext.Errors, &gin.Error{Err: err})
		return
	}

	userSessionKey := generalService.GetSessionKeyFromHeader(ginContext)
	userServiceUtil := userService.NewUserServiceUtil()

	if err = isLoggedIn(userServiceUtil, userSessionKey, loginInfo); err != nil {
		ginContext.JSON(http.StatusOK, gin.H{"message": "Already logged in"})
		log.Print("User is already logged in")
		return
	}

	if err = authenticateUser(userServiceUtil, ginContext, *loginInfo, userSessionKey); err != nil {
		ginContext.Errors = append(ginContext.Errors, &gin.Error{Err: err})
		return
	}

	sessionKey, err := generateOrGetSessionKey(userServiceUtil, *loginInfo)
	if err != nil {
		ginContext.Errors = append(ginContext.Errors, &gin.Error{Err: err})
		return
	}

	ginContext.JSON(http.StatusOK, gin.H{"sessionKey": sessionKey})
}

func getLoginInfo(ginContext *gin.Context) (*models.LoginInfo, error) {
	loginInfo, err := generalService.GetLoginInfoFromBody(ginContext)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	return &loginInfo, err
}

func isLoggedIn(userServiceUtil *userService.UserServiceUtil, userSessionKey string, loginInfo *models.LoginInfo) error {
	_, isLoggedIn := userServiceUtil.CheckUserLoggedIn(userSessionKey, *loginInfo)

	if isLoggedIn {
		return userService.ErrAlreadyLoggedIn
	}

	return nil
}

func authenticateUser(userServiceUtil *userService.UserServiceUtil, ginContext *gin.Context, loginInfo models.LoginInfo, userSessionKey string) error {
	var err error
	_, err = userServiceUtil.AuthenticateUserByLoginInfo(loginInfo, userSessionKey)
	if err != nil {
		userServiceUtil.HandleLoginError(ginContext, err)
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	return err
}

func generateOrGetSessionKey(userServiceUtil *userService.UserServiceUtil, loginInfo models.LoginInfo) (string, error) {
	var sessionKey string
	var err error
	if loginInfo.LoginSessionID == "" {
		sessionKey, err = userServiceUtil.GenerateSessionKey(loginInfo)
		if err != nil {
			return "", err
		}
		loginInfo.LoginSessionID = sessionKey
	} else {
		sessionKey = loginInfo.LoginSessionID
	}
	return sessionKey, nil
}
