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

	t.Run("SignUp", func(t *testing.T) {
		request, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(userJson))
		response := httptest.NewRecorder()
		ginContext, _ := gin.CreateTestContext(response)
		ginContext.Request = request
		SignUp(ginContext)

		if !assert.Equal(t, http.StatusCreated, response.Code) {
			t.Logf("test Signup failed: %v", response.Body.String())
			return
		}
	})

	t.Run("LogIn", func(t *testing.T) {
		request, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(userJson))
		response := httptest.NewRecorder()
		ginContext, _ := gin.CreateTestContext(response)
		ginContext.Request = request
		LogIn(ginContext)

		if !assert.Equal(t, http.StatusOK, response.Code) {
			t.Logf("test Login failed: %v", response.Body.String())
			return
		}
	})

	t.Run("LogOut", func(t *testing.T) {
		// LogIn before LogOut
		request, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(userJson))
		response := httptest.NewRecorder()
		ginContext, _ := gin.CreateTestContext(response)
		ginContext.Request = request
		LogIn(ginContext)

		// Parse the response body to get the sessionKey
		var responseBody map[string]string
		json.Unmarshal(response.Body.Bytes(), &responseBody)
		sessionKey := responseBody["sessionKey"]

		// LogOut with the sessionKey from the LogIn response
		request, _ = http.NewRequest("POST", "/logout", nil)
		request.Header.Set("Session-Key", sessionKey)
		response = httptest.NewRecorder()
		ginContext, _ = gin.CreateTestContext(response)
		ginContext.Request = request
		LogOut(ginContext)

		if !assert.Equal(t, http.StatusOK, response.Code) {
			t.Logf("test Logout failed: %v", response.Body.String())
			return
		}
	})

	t.Run("SignOut", func(t *testing.T) {
		// LogIn before SignOut
		request, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(userJson))
		response := httptest.NewRecorder()
		ginContext, _ := gin.CreateTestContext(response)
		ginContext.Request = request
		LogIn(ginContext)

		// Parse the response body to get the sessionKey
		var responseBody map[string]string
		json.Unmarshal(response.Body.Bytes(), &responseBody)
		sessionKey := responseBody["sessionKey"]

		// SignOut with the sessionKey from the LogIn response
		request, _ = http.NewRequest("POST", "/signout", nil)
		request.Header.Set("Session-Key", sessionKey)
		response = httptest.NewRecorder()
		ginContext, _ = gin.CreateTestContext(response)
		ginContext.Request = request
		SignOut(ginContext)

		if !assert.Equal(t, http.StatusOK, response.Code) {
			t.Logf("test SignOut failed: %v", response.Body.String())
			return
		}
	})
}
