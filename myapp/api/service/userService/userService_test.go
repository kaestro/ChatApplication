// myapp/api/service/userService/createAndDeleteUser_test.go
package userService

import (
	"myapp/api/models"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
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

	t.Run("TestDeauthenticateUser", func(t *testing.T) {
		loginService := NewLoginService()
		sessionKey, _ := loginService.AuthenticateUser(sampleLoginInfo, "")

		err := DeauthenticateUser(sessionKey)
		if !assert.Nil(t, err) {
			t.Logf("Failed to deauthenticate user: %v", err)
			return
		}

		t.Logf("Successfully deauthenticated user")
	})

	t.Run("TestDeleteUserBySessionKey", func(t *testing.T) {
		loginService := NewLoginService()
		sessionKey, _ := loginService.AuthenticateUser(sampleLoginInfo, "")

		err := DeleteUserBySessionKey(sessionKey, ginContext)
		if !assert.Nil(t, err) {
			t.Logf("Failed to delete user: %v", err)
			return
		}
	})
}
