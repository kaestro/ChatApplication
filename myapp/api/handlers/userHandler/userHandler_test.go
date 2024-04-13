// myapp/api/handlers/user/userHandler_test.go
package userHandler

import (
	"bytes"
	"encoding/json"
	"myapp/api/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler(t *testing.T) {
	sampleEmailAddress := "tuh@gmail.com"
	samplePassword := "testpassword"

	gin.SetMode(gin.TestMode)

	user := models.User{
		EmailAddress: sampleEmailAddress,
		Password:     samplePassword,
	}

	userJson, _ := json.Marshal(user)

	// Run subtests sequentially
	sessionKey := testSignUp(t, userJson)
	sessionKey = testLogIn(t, userJson, sessionKey)
	sessionKey = testLogOut(t, sessionKey)
	sessionKey = testLogIn(t, userJson, sessionKey)
	testSignOut(t, sessionKey)
}

func testSignUp(t *testing.T, userJson []byte) string {
	request, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(userJson))
	response := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(response)
	ginContext.Request = request
	SignUp(ginContext)

	if !assert.Equal(t, http.StatusCreated, response.Code) {
		t.Logf("test Signup failed: %v", response.Body.String())
	}

	var responseBody map[string]string
	json.Unmarshal(response.Body.Bytes(), &responseBody)
	return responseBody["sessionKey"]
}

func testLogIn(t *testing.T, userJson []byte, sessionKey string) string {
	request, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(userJson))
	request.Header.Set("Session-Key", sessionKey)
	response := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(response)
	ginContext.Request = request
	LogIn(ginContext)

	if !assert.Equal(t, http.StatusOK, response.Code) {
		t.Logf("test Login failed: %v", response.Body.String())
	}

	var responseBody map[string]string
	json.Unmarshal(response.Body.Bytes(), &responseBody)
	return responseBody["sessionKey"]
}

func testLogOut(t *testing.T, sessionKey string) string {
	request, _ := http.NewRequest("POST", "/logout", nil)
	request.Header.Set("Session-Key", sessionKey)
	response := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(response)
	ginContext.Request = request
	LogOut(ginContext)

	if !assert.Equal(t, http.StatusOK, response.Code) {
		t.Logf("test Logout failed: %v", response.Body.String())
	}

	var responseBody map[string]string
	json.Unmarshal(response.Body.Bytes(), &responseBody)
	return responseBody["sessionKey"]
}

func testSignOut(t *testing.T, sessionKey string) {
	request, _ := http.NewRequest("POST", "/signout", nil)
	request.Header.Set("Session-Key", sessionKey)
	response := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(response)
	ginContext.Request = request
	SignOut(ginContext)

	if !assert.Equal(t, http.StatusOK, response.Code) {
		t.Logf("test SignOut failed: %v", response.Body.String())
	}
}
