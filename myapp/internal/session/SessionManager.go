// myapp/internal/session/SessionManager.go - Singleton session manager
package session

import (
	"context"
	"sync"

	"github.com/go-redis/redis/v8"
)

type SessionManager struct {
	client *redis.Client
}

func (manager *SessionManager) GetSession(key string) (string, error) {
	val, err := manager.client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func (manager *SessionManager) SetSession(key string, value string) error {
	err := manager.client.Set(context.Background(), key, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (manager *SessionManager) DeleteSession(key string) error {
	err := manager.client.Del(context.Background(), key).Err()
	if err != nil {
		return err
	}

	return nil
}

var (
	once    sync.Once
	manager *SessionManager
)

func GetSessionManager() *SessionManager {
	once.Do(func() {
		manager = &SessionManager{
			client: redis.NewClient(&redis.Options{
				Addr:     "localhost:6379",
				Password: "redisPassword", // no password set
				DB:       0,               // use default DB
			}),
		}
	})

	return manager
}
