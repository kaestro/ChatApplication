// myapp/internal/session/LoginSessionManager_test.go
package session

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LoginSessionManagerSuite struct{
	suite.Suite
	redisClient *redis.Client
	loginSession SessionStore
}

func TestLoginSeesionManagerSuite(t *testing.T) {
	suite.Run(t, new(LoginSessionManagerSuite))
}

func (s *LoginSessionManagerSuite) SetupSuite() {
	s.redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "redisPassword", // no password set
		DB:       0,               // use default DB
	})

	ctx := context.Background()

	_,err := s.redisClient.Ping(ctx).Result()
	if err != nil{
		s.T().Fatal(fmt.Errorf("failed to ping redis: %v", err))
	}

	fmt.Println("Redis client is ready")

	// 각 테스트 진행시 redis 에 저장된 데이터 초기화
	if err := s.redisClient.FlushAll(ctx).Err(); err != nil {
		s.T().Fatal(fmt.Errorf("failed to flush redis: %v", err))
	}

	s.loginSession = &RedisStore{
		client: s.redisClient,
	}
}

func (s *LoginSessionManagerSuite) TestGetSession() {
		t := s.T()
		
		// 세션 키가 없는 상황을 테스트합니다.
		t.Run("조회시 key 가 없는 경우 테스트", func(t *testing.T) {		
			key := "no_existing_key"

			result, err := s.loginSession.GetSession(key)

			assert.EqualError(t, err, "redis: nil")
			assert.Empty(t, result, "result should be empty")
		})

		// 세션 키가 있는 상황을 테스트합니다.
		t.Run("조회시 key 가 있는 경우", func(t *testing.T) {
			key := "existing_key"
			value := "existing_value"

			err := s.loginSession.SetSession(key, value)

			assert.Nil(t, err) 

			result, err := s.loginSession.GetSession("existing_key")

			assert.Nil(t, err)
			assert.Equal(t, "existing_value", result)
		})
}

func (s *LoginSessionManagerSuite) TestSetSession() {
	t := s.T()

	t.Run("데이터 저장 테스트",func(t *testing.T) {
		key := "insert_key"
		value := "insert_test_value"

		err := s.loginSession.SetSession(key, value)

		assert.Nil(t,err)
	})

	t.Run("데이터 저장후 조회 테스트",func(t *testing.T) {
		key := "insert_key"

		result, err := s.loginSession.GetSession(key)

		assert.Nil(t, err)
		assert.Equal(t, "insert_test_value", result)
	})
}

func (s *LoginSessionManagerSuite) TestDeleteSession() {
	t := s.T()

	t.Run("삭제시 key 가 없는 경우 테스트",func(t *testing.T) {
		key := "no_existing_key"

		err := s.loginSession.DeleteSession(key)
		assert.Nil(t, err, "error deleting session")
	})

	t.Run("삭제시 key 가 있는 경우 테스트",func(t *testing.T) {
			key := "existing_key"
			value := "existing_value"

			err := s.loginSession.SetSession(key, value)
			assert.Nil(t, err, "error setting session") 

			err = s.loginSession.DeleteSession(key)
			assert.Nil(t, err, "error deleting session")
	})
}

func (s *LoginSessionManagerSuite) TestIsSessionValid() {
	t := s.T()

	t.Run("유효한 세션인 경우", func(t *testing.T) {
		// 유효한 세션 키와 이메일 주소를 사용하여 세션이 유효한지 테스트합니다.
		key := "valid_key"
		emailAddress := "valid@example.com"

		// 세션을 유효하게 설정합니다.
		err := s.loginSession.SetSession(key, emailAddress)
		assert.Nil(t, err, "error setting session")

		// 세션이 유효한지 확인합니다.
		valid := s.loginSession.IsSessionValid(key, emailAddress)
		assert.True(t, valid, "session should be valid")
	})

	t.Run("잘못된 세션 키 또는 이메일 주소인 경우", func(t *testing.T) {
		// 잘못된 세션 키와 이메일 주소를 사용하여 세션이 유효한지 테스트합니다.
		key := "invalid_key"
		emailAddress := "invalid@example.com"

		// 세션이 유효하지 않도록 설정합니다.
		err := s.loginSession.DeleteSession(key)
		assert.Nil(t, err, "error deleting session")

		// 세션이 유효한지 확인합니다.
		valid := s.loginSession.IsSessionValid(key, emailAddress)
		assert.False(t, valid, "session should not be valid")
	})

	t.Run("세션 키가 존재하지 않는 경우", func(t *testing.T) {
		// 세션 키가 존재하지 않는 경우를 테스트합니다.
		key := "nonexistent_key"
		emailAddress := "valid@example.com"

		// 세션이 유효한지 확인합니다.
		valid := s.loginSession.IsSessionValid(key, emailAddress)
		assert.False(t, valid, "session should not be valid")
	})

	t.Run("세션 키가 존재하나 이메일 주소가 일치하지 않는 경우", func(t *testing.T) {
		// 세션 키가 존재하지만 이메일 주소가 일치하지 않는 경우를 테스트합니다.
		key := "existing_key"
		emailAddress := "valid@example.com"
		invalidEmailAddress := "invalid@example.com"

		// 세션을 설정합니다.
		err := s.loginSession.SetSession(key, emailAddress)
		assert.Nil(t, err, "error setting session")

		// 세션이 유효한지 확인합니다.
		valid := s.loginSession.IsSessionValid(key, invalidEmailAddress)
		assert.False(t, valid, "session should not be valid")
	})
}
