package tests

import (
	"myapp/api/handlers/user"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogOut(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("LogoutSuccess", func(t *testing.T) {
		router := gin.Default()

		router.GET("/logout", user.LogOut)

		req, _ := http.NewRequest("GET", "/logout", nil)
		req.Header.Add("Session-Key", "existingSessionKey")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

}
