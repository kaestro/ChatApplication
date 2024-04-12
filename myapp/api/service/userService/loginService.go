// myapp/api/service/user/loginService.go
package user

import (
	"errors"
	"myapp/api/models"
	"myapp/internal/db"
	"myapp/internal/password"
	"myapp/internal/session"
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

func NewLoginService(dbManager db.DBManagerInterface, sessionManager *session.LoginSessionManager) *LoginService {
	return &LoginService{
		dbManager:      dbManager,
		sessionManager: sessionManager,
	}
}

func (s *LoginService) LogIn(userEmailAddress, userPassword, userSessionKey string) (string, error) {
	// 세션 키가 sessionManager에 저장되어 있는지 확인합니다.
	if s.sessionManager.IsSessionValid(userSessionKey, userEmailAddress) {
		return "", ErrAlreadyLoggedIn
	}

	// 사용자 정보를 담을 User 구조체를 선언합니다.
	var user models.User

	// 사용자가 제공한 이메일 주소로 데이터베이스에서 사용자를 찾습니다.
	err := s.dbManager.Read(&user, "email_address", userEmailAddress)
	if err != nil {
		return "", ErrUserNotFound
	}

	// 사용자가 제공한 비밀번호와 데이터베이스에 저장된 해시된 비밀번호를 비교합니다.
	if !password.CheckPasswordHash(userPassword, user.Password) {
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
