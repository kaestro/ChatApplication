// myapp/api/service/userService/deauthenticateUser.go
package userService

import (
	"myapp/internal/session"
)

func DeauthenticateUser(userSessionKey string) error {
	sessionManager := session.GetLoginSessionManager()
	err := sessionManager.DeleteSession(userSessionKey)
	return err
}
