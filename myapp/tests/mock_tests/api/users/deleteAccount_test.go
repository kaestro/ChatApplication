package tests

import (
	"myapp/api/handlers/user"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDeleteAccountSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/user/deleteAccount", user.DeleteAccount)

	req, _ := http.NewRequest("DELETE", "/user/deleteAccount", nil)
	req.Header.Add("Session-Key", "valid_session_key")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
