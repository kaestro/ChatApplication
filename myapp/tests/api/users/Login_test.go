// tests/api/users/Login_test.go
package tests

import (
	"bytes"
	"encoding/json"
	userAPI "myapp/api/handlers/user"
	"myapp/api/models"
	"myapp/internal/db"
	"myapp/internal/password"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogIn(t *testing.T) {
	// 테스트를 위한 사용자 정보를 생성합니다.
	hashedPassword, _ := password.HashPassword("password")
	user := models.User{
		EmailAddress: "test@example.com",
		Password:     hashedPassword,
	}
	db.GetDBManager().Create(&user)

	// 잘못된 비밀번호로 로그인을 시도합니다.
	loginInfo := struct {
		EmailAddress string `json:"emailAddress"`
		Password     string `json:"password"`
	}{
		EmailAddress: "test@example.com",
		Password:     "wrongpassword",
	}
	body, _ := json.Marshal(loginInfo)
	httpRequest, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	responseRecorder := httptest.NewRecorder()

	// Gin 엔진을 생성하고 LogIn 핸들러를 등록합니다.
	ginEngine := gin.Default()
	ginEngine.POST("/login", userAPI.LogIn)

	// HTTP 요청을 처리합니다.
	ginEngine.ServeHTTP(responseRecorder, httpRequest)

	// 응답 상태 코드가 401인지 확인합니다.
	assert.Equal(t, http.StatusUnauthorized, responseRecorder.Code)

	// 올바른 비밀번호로 로그인을 시도합니다.
	loginInfo.Password = "password"
	body, _ = json.Marshal(loginInfo)
	httpRequest, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	responseRecorder = httptest.NewRecorder()

	// HTTP 요청을 처리합니다.
	ginEngine.ServeHTTP(responseRecorder, httpRequest)

	// 응답 상태 코드가 200인지 확인합니다.
	if assert.Equal(t, http.StatusOK, responseRecorder.Code) {
		t.Log("LogIn 테스트 통과")
	}

	// 해당 계정을 삭제합니다.
	db.GetDBManager().Delete(&user)
}
