// session/RedisStore.go
package session

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisStore struct {
	client *redis.Client
}

type RedisStoreFactory struct{}

func (factory *RedisStoreFactory) Create(sessionTypeNum SessionType) SessionStore {
	var store SessionStore
	if sessionTypeNum == LoginSession {
		redisAddr := os.Getenv("REDIS_ADDR")
		if redisAddr == "" {
			redisAddr = "localhost:6379" // default value
		}

		store = &RedisStore{
			client: redis.NewClient(&redis.Options{
				Addr:     redisAddr,
				Password: "redisPassword", // no password set
				DB:       0,               // use default DB
			}),
		}
	} else if sessionTypeNum == OtherSession {
		panic("Unauthorized session type number given to RedisStoreFactory.")
	}
	return store
}

func NewRedisStore(sessionTypeNum SessionType) SessionStore {
	factory := &RedisStoreFactory{}
	return factory.Create(sessionTypeNum)
}

func (store *RedisStore) GetSession(key string) (string, error) {
	val, err := store.client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}

	// 세션 키를 조회할 때마다 만료 시간을 30분으로 갱신합니다.
	store.client.Expire(context.Background(), key, 30*time.Minute)

	return val, nil
}

func (store *RedisStore) SetSession(key string, value string) error {
	err := store.client.Set(context.Background(), key, value, 30*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

func (store *RedisStore) DeleteSession(key string) error {
	err := store.client.Del(context.Background(), key).Err()
	if err != nil {
		return err
	}

	return nil
}

func (store *RedisStore) IsSessionValid(key string, emailAddress string) bool {
	val, err := store.GetSession(key)
	if err != nil {
		return false
	}

	return val == emailAddress
}
