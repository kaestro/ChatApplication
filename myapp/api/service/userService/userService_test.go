// myapp/api/service/userService/createAndDeleteUser_test.go
package userService

import (
	"myapp/api/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndDeleteUserByEmailAddress(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())

	sampleEmailAddress := "tcduea@gmail.com"
	samplePassword := "testpassword"
	sampleUserName := "tcduea"

	user := models.User{
		EmailAddress: sampleEmailAddress,
		Password:     samplePassword,
		UserName:     sampleUserName,
	}

	err := CreateUser(user, ginContext)
	if !assert.Nil(t, err) {
		t.Logf("Failed to create user: %v", err)
		return
	}

	err = DeleteUserByEmailAddress(sampleEmailAddress)
	if !assert.Nil(t, err) {
		t.Logf("Failed to delete user: %v", err)
		return
	}

	t.Logf("Successfully created and deleted user by email address")
}

func TestUserService(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())

	sampleEmailAddress := "tus@gmail.com"
	samplePassword := "testpassword"
	sampleUserName := "tus"
	sampleUser := models.User{
		EmailAddress: sampleEmailAddress,
		Password:     samplePassword,
		UserName:     sampleUserName,
	}
	sampleLoginInfo := models.LoginInfo{
		EmailAddress: sampleEmailAddress,
		Password:     samplePassword,
	}

	t.Run("TestCreateUser", func(t *testing.T) {
		err := CreateUser(sampleUser, ginContext)
		if !assert.Nil(t, err) {
			t.Logf("Failed to create user: %v", err)
			return
		}
		t.Logf("Successfully created user")
	})

	t.Run("TestAuthenticateUser", func(t *testing.T) {
		loginService := NewLoginService()
		_, err := loginService.AuthenticateUser(sampleLoginInfo, "")
		if !assert.Nil(t, err) {
			t.Logf("Failed to authenticate user: %v", err)
			return
		}

		t.Logf("Successfully authenticated user")
	})

	t.Run("TestPublishWebSocket", func(t *testing.T) {
		// 웹소켓 서버 시작
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			loginService := NewLoginService()
			err := loginService.PublishWebSocket(w, r, "testSessionKey")
			if err != nil {
				t.Errorf("Expected no error, but got %v", err)
			}
		}))
		defer server.Close()

		// 웹소켓 클라이언트로 연결 시도
		_, _, err := websocket.DefaultDialer.Dial(strings.Replace(server.URL, "http", "ws", 1), nil)
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}
	})

	t.Run("TestDeauthenticateUser", func(t *testing.T) {
		loginService := NewLoginService()
		user, _ := loginService.AuthenticateUser(sampleLoginInfo, "")
		sessionKey, _ := loginService.GenerateSessionKey(user)

		err := DeauthenticateUser(sessionKey)
		if !assert.Nil(t, err) {
			t.Logf("Failed to deauthenticate user: %v", err)
			return
		}

		t.Logf("Successfully deauthenticated user")
	})

	t.Run("TestDeleteUserBySessionKey", func(t *testing.T) {
		loginService := NewLoginService()
		user_model, _ := loginService.AuthenticateUser(sampleLoginInfo, "")
		sessionKey, _ := loginService.GenerateSessionKey(user_model)

		err := DeleteUserBySessionKey(sessionKey, ginContext)
		if !assert.Nil(t, err) {
			t.Logf("Failed to delete user: %v", err)
			return
		}
	})
}
