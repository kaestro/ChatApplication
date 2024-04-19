// myapp/internal/session/SessionManager.go
package session

import (
	"sync"
)

type LoginSessionManager struct {
	store SessionStore
}

func (manager *LoginSessionManager) GetSession(key string) (string, error) {
	return manager.store.GetSession(key)
}

func (manager *LoginSessionManager) SetSession(key string, value string) error {
	return manager.store.SetSession(key, value)
}

func (manager *LoginSessionManager) DeleteSession(key string) error {
	return manager.store.DeleteSession(key)
}

var (
	once    sync.Once
	manager *LoginSessionManager
)

func GetLoginSessionManager() *LoginSessionManager {
	once.Do(func() {
		manager = &LoginSessionManager{
			store: NewRedisStore(LoginSession),
		}
	})

	return manager
}

func (manager *LoginSessionManager) IsSessionValid(sessionID string, emailAddress string) bool {
	return manager.store.IsSessionValid(sessionID, emailAddress)
}
