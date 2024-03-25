package tests

import (
	"bytes"
	"encoding/json"
	"myapp/api/handlers/user"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogIn(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("InvalidRequestBody", func(t *testing.T) {

		router := gin.Default()

		router.POST("/login", user.LogIn)

		requestBody := bytes.NewBufferString(`{invalid json}`)
		req, _ := http.NewRequest(http.MethodPost, "/login", requestBody)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

}

func TestLogIn_UserNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/login", user.LogIn)

	requestBody, _ := json.Marshal(map[string]string{
		"emailAddress": "nonexistent@example.com",
		"password":     "password",
	})
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
