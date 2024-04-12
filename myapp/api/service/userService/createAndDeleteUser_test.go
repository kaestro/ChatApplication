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

	user := models.User{
		EmailAddress: sampleEmailAddress,
		Password:     samplePassword,
	}

	// Test CreateUser
	err := CreateUser(user, ginContext)
	assert.Nil(t, err)

	// Test DeleteUserBySessionKey
	// Assuming that the session key for the created user is "testSessionKey"
	err = DeleteUserByEmailAddress(sampleEmailAddress)
	assert.Nil(t, err)
}

func TestUserService(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())

	user := models.User{
		EmailAddress: sampleEmailAddress,
		Password:     samplePassword,
	}

	// Test CreateUser
	err := CreateUser(user, ginContext)
	if !assert.Nil(t, err) {
		t.Logf("Failed to create user: %v", err)
		return
	}

	// Test AuthenticateUser
	loginInfo := models.LoginInfo{
		EmailAddress: sampleEmailAddress,
		Password:     samplePassword,
	}

	loginService := NewLoginService()
	sessionKey, err := loginService.AuthenticateUser(loginInfo, sampleSessionkey)
	if !assert.Nil(t, err) {
		t.Logf("Failed to authenticate user: %v", err)
		return
	}

	// Test DeleteUserBySessionKey
	err = DeleteUserBySessionKey(sessionKey, ginContext)
	if !assert.Nil(t, err) {
		t.Logf("Failed to delete user: %v", err)
		return
	}

	t.Logf("Successfully created, authenticated, and deleted user")
}
