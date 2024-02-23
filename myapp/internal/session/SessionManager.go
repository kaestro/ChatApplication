// myapp/internal/session/SessionManager.go
package session

import (
	"sync"
)

type SessionManager struct {
	store Store
}

func (manager *SessionManager) GetSession(key string) (string, error) {
	return manager.store.GetSession(key)
}

func (manager *SessionManager) SetSession(key string, value string) error {
	return manager.store.SetSession(key, value)
}

func (manager *SessionManager) DeleteSession(key string) error {
	return manager.store.DeleteSession(key)
}

var (
	once    sync.Once
	manager *SessionManager
)

func GetSessionManager() *SessionManager {
	once.Do(func() {
		manager = &SessionManager{
			store: NewRedisStore(),
		}
	})

	return manager
}

func (manager *SessionManager) IsSessionValid(key string, emailAddress string) bool {
	return manager.store.IsSessionValid(key, emailAddress)
}
