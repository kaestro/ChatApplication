// myapp/internal/session/SessionManager.go - Singleton session manager
package session

import (
	"context"
	"sync"
	"time"

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

	// 세션 키를 조회할 때마다 만료 시간을 30분으로 갱신합니다.
	manager.client.Expire(context.Background(), key, 30*time.Minute)

	return val, nil
}

func (manager *SessionManager) SetSession(key string, value string) error {
	err := manager.client.Set(context.Background(), key, value, 30*time.Minute).Err()
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

func (manager *SessionManager) IsSessionValid(key string, emailAddress string) bool {
	val, err := manager.GetSession(key)
	if err != nil {
		return false
	}

	return val == emailAddress
}
