// myapp/api/service/userService/userServiceUtil.go
package userService

import (
	"errors"
	"fmt"
	"myapp/api/models"
	"myapp/internal/db"
	"myapp/internal/password"
	"myapp/internal/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrAlreadyLoggedIn            = errors.New("user is already logged in")
	ErrUserNotFound               = errors.New("failed to find user")
	ErrInvalidPassword            = errors.New("invalid password")
	ErrFailedToGenerateSessionKey = errors.New("failed to generate session key")
	ErrFailedToSaveSessionKey     = errors.New("failed to save session key")
	dbColumnUserIdentifier        = "email_address"
)

type UserServiceUtil struct {
	dbManager      db.DBManagerInterface
	sessionManager *session.LoginSessionManager
}

func NewUserServiceUtil() *UserServiceUtil {
	return &UserServiceUtil{
		dbManager:      db.GetDBManager(),
		sessionManager: session.GetLoginSessionManager(),
	}
}

// 이미 로그인 된 경우 loginInfo의 세션 키를 userSessionKey로 설정하고,
// 아닐 경우 세션 키를 생성해서 지급한 loginInfo를 반환합니다.
func (usu *UserServiceUtil) AuthenticateUser(loginInfo models.LoginInfo, userSessionKey string) (models.LoginInfo, error) {
	_, isLoggedIn := usu.CheckUserLoggedIn(userSessionKey, loginInfo)

	var user models.User
	err := usu.dbManager.Read(&user, dbColumnUserIdentifier, loginInfo.EmailAddress)
	if err != nil {
		return models.LoginInfo{}, ErrUserNotFound
	} else if isLoggedIn {
		loginInfo.LoginSessionID = userSessionKey
		return loginInfo, nil
	}

	if !password.CheckPasswordHash(loginInfo.Password, user.Password) {
		return models.LoginInfo{}, ErrInvalidPassword
	}

	loginSession := session.GetLoginSessionManager()
	sessionKey, err := session.GenerateRandomSessionKey()
	if err != nil {
		return models.LoginInfo{}, ErrFailedToGenerateSessionKey
	}
	loginInfo.LoginSessionID = sessionKey
	err = loginSession.SetSession(sessionKey, loginInfo.EmailAddress)
	if err != nil {
		return models.LoginInfo{}, ErrFailedToSaveSessionKey
	}

	return loginInfo, nil
}

func (usu *UserServiceUtil) GenerateSessionKey(user models.LoginInfo) (string, error) {
	sessionKey, err := session.GenerateRandomSessionKey()
	if err != nil {
		return "", ErrFailedToGenerateSessionKey
	}

	// 세션 키를 캐시에 저장합니다.
	err = usu.sessionManager.SetSession(sessionKey, user.EmailAddress)
	if err != nil {
		return "", ErrFailedToSaveSessionKey
	}

	return sessionKey, nil
}

func (usu *UserServiceUtil) CheckUserLoggedIn(userSessionKey string, loginInfo models.LoginInfo) (string, bool) {
	if usu.sessionManager.IsSessionValid(userSessionKey, loginInfo.EmailAddress) {
		fmt.Println("User is already logged in")

		_, err := usu.sessionManager.GetSession(loginInfo.LoginSessionID)
		if err != nil {
			return "", false
		}

		return userSessionKey, true
	}
	return "", false
}

// handleLoginError 함수는 로그인 과정에서 발생한 오류를 처리합니다.
// 오류 유형에 따라 적절한 HTTP 상태 코드를 반환합니다.
func (usu *UserServiceUtil) HandleLoginError(ginContext *gin.Context, err error) {
	switch err {
	case ErrAlreadyLoggedIn:
		ginContext.JSON(http.StatusConflict, gin.H{"error": "User is already logged in"})
	case ErrUserNotFound:
		ginContext.JSON(http.StatusNotFound, gin.H{"error": "Failed to find user"})
	case ErrInvalidPassword:
		ginContext.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
	case ErrFailedToGenerateSessionKey, ErrFailedToSaveSessionKey:
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process session key"})
	default:
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
}
