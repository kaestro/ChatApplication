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
// 로그인이 안 돼있지만, 가능한 경우 세션 키를 생성해서 지급한 loginInfo를 반환합니다.
// 로그인이 불가능한 경우 빈 loginInfo와 오류를 반환합니다.
func (usu *UserServiceUtil) AuthenticateUserByLoginInfo(loginInfo models.LoginInfo, userSessionKey string) (models.LoginInfo, error) {
	// Step 1: Check if user is already logged in
	isLoggedIn, err := usu.isUserLoggedIn(userSessionKey, loginInfo)
	if err != nil {
		return models.LoginInfo{}, err
	} else if isLoggedIn {
		loginInfo.LoginSessionID = userSessionKey
		return loginInfo, nil
	}

	// Step 2: Validate user credentials
	err = usu.validateUserCredentials(loginInfo)
	if err != nil {
		return models.LoginInfo{}, err
	}

	// Step 3: Create new session for user
	loginInfo, err = usu.createNewUserSession(loginInfo)
	if err != nil {
		return models.LoginInfo{}, err
	}

	return loginInfo, nil
}

func (usu *UserServiceUtil) isUserLoggedIn(userSessionKey string, loginInfo models.LoginInfo) (bool, error) {
	_, isLoggedIn := usu.CheckUserLoggedIn(userSessionKey, loginInfo)
	return isLoggedIn, nil
}

func (usu *UserServiceUtil) validateUserCredentials(loginInfo models.LoginInfo) error {
	var user models.User
	err := usu.dbManager.Read(&user, dbColumnUserIdentifier, loginInfo.EmailAddress)
	if err != nil {
		return ErrUserNotFound
	}

	if !password.CheckPasswordHash(loginInfo.Password, user.Password) {
		return ErrInvalidPassword
	}

	return nil
}

func (usu *UserServiceUtil) createNewUserSession(loginInfo models.LoginInfo) (models.LoginInfo, error) {
	loginSession := session.GetLoginSessionManager()

	sessionKey, err := session.GenerateRandomSessionKey()
	if err != nil {
		return models.LoginInfo{}, ErrFailedToGenerateSessionKey
	}

	err = loginSession.SetSession(sessionKey, loginInfo.EmailAddress)
	if err != nil {
		return models.LoginInfo{}, ErrFailedToSaveSessionKey
	}

	loginInfo.LoginSessionID = sessionKey
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

func (usu *UserServiceUtil) AuthenticateUserByLoginSessionInfo(loginSessionInfo models.LoginSessionInfo) error {
	if !usu.sessionManager.IsSessionValid(loginSessionInfo.LoginSessionID, loginSessionInfo.EmailAddress) {
		return errors.New("user is not logged in")
	}
	return nil

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
