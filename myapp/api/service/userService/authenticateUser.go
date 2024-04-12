// myapp/api/service/userService/authenticateUser.go
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
)

type LoginService struct {
	dbManager      db.DBManagerInterface
	sessionManager *session.LoginSessionManager
}

func NewLoginService() *LoginService {
	return &LoginService{
		dbManager:      db.GetDBManager(),
		sessionManager: session.GetLoginSessionManager(),
	}
}

func (s *LoginService) AuthenticateUser(loginInfo models.LoginInfo, userSessionKey string) (string, error) {
	// 세션 키가 sessionManager에 저장되어 있는지 확인합니다.
	if s.sessionManager.IsSessionValid(userSessionKey, loginInfo.EmailAddress) {
		fmt.Println("User is already logged in")

		sessionKey, err := s.sessionManager.GetSession(loginInfo.EmailAddress)
		if err != nil {
			return "", err
		}

		return sessionKey, ErrAlreadyLoggedIn
	}

	// 사용자 정보를 담을 User 구조체를 선언합니다.
	var user models.User

	// 사용자가 제공한 이메일 주소로 데이터베이스에서 사용자를 찾습니다.
	err := s.dbManager.Read(&user, "email_address", loginInfo.EmailAddress)
	if err != nil {
		return "", ErrUserNotFound
	}

	// 사용자가 제공한 비밀번호와 데이터베이스에 저장된 해시된 비밀번호를 비교합니다.
	if !password.CheckPasswordHash(loginInfo.Password, user.Password) {
		return "", ErrInvalidPassword
	}

	// 세션 키를 생성합니다.
	sessionKey, err := session.GenerateRandomSessionKey()
	if err != nil {
		return "", ErrFailedToGenerateSessionKey
	}

	// 세션 키를 캐시에 저장합니다.
	err = s.sessionManager.SetSession(sessionKey, user.EmailAddress)
	if err != nil {
		return "", ErrFailedToSaveSessionKey
	}

	return sessionKey, nil
}

// handleLoginError 함수는 로그인 과정에서 발생한 오류를 처리합니다.
// 오류 유형에 따라 적절한 HTTP 상태 코드를 반환합니다.
func HandleLoginError(ginContext *gin.Context, err error) {
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
