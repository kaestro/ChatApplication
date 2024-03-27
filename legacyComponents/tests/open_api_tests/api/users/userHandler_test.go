// tests/api/users/Login_test.go
package tests

import (
	"bytes"
	"encoding/json"
	userAPI "myapp/api/handlers/user"
	"myapp/api/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler(t *testing.T) {
	// 테스트를 위한 사용자 정보를 생성합니다.
	sampleUser := models.User{
		UserName:     "testuser",
		EmailAddress: "test@example.com",
		Password:     "password",
	}

	// Gin 엔진을 생성하고 핸들러들을 등록합니다.
	ginEngine := gin.Default()
	ginEngine.POST("/login", userAPI.LogIn)
	ginEngine.POST("/logout", userAPI.LogOut)
	ginEngine.POST("/signup", userAPI.SignUp)
	ginEngine.POST("/deleteAccount", userAPI.DeleteAccount)

	// signup HTTP 요청을 처리합니다.
	body, _ := json.Marshal(sampleUser)
	httpRequest, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(body))
	responseRecorder := httptest.NewRecorder()
	ginEngine.ServeHTTP(responseRecorder, httpRequest)

	t.Log(responseRecorder.Body.String())

	// 응답 상태 코드가 201인지 확인합니다.
	if assert.Equal(t, http.StatusCreated, responseRecorder.Code) {
		t.Log("SignUp 테스트 통과")
	} else {
		t.Log("SignUp 테스트 실패")
	}

	// login HTTP 요청을 처리합니다.
	loginInfo := struct {
		EmailAddress string `json:"emailAddress"`
		Password     string `json:"password"`
	}{
		EmailAddress: sampleUser.EmailAddress,
		Password:     sampleUser.Password,
	}
	body, _ = json.Marshal(loginInfo)
	httpRequest, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	responseRecorder = httptest.NewRecorder()
	ginEngine.ServeHTTP(responseRecorder, httpRequest)

	// 응답 상태 코드가 200인지 확인합니다.
	if assert.Equal(t, http.StatusOK, responseRecorder.Code) {
		t.Log("LogIn 테스트 통과")
	} else {
		t.Log("LogIn 테스트 실패")
	}

	// logout HTTP 요청을 처리합니다.
	httpRequest, _ = http.NewRequest("POST", "/logout", nil)
	httpRequest.Header.Set("Session-Key", responseRecorder.Body.String())
	responseRecorder = httptest.NewRecorder()
	ginEngine.ServeHTTP(responseRecorder, httpRequest)

	// 응답 상태 코드가 200인지 확인합니다.
	if assert.Equal(t, http.StatusOK, responseRecorder.Code) {
		t.Log("LogOut 테스트 통과")
	} else {
		t.Log("LogOut 테스트 실패")
	}

	// deleteAccount HTTP 요청을 처리합니다.
	body, _ = json.Marshal(loginInfo)
	httpRequest, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	responseRecorder = httptest.NewRecorder()
	ginEngine.ServeHTTP(responseRecorder, httpRequest)

	// 응답 상태 코드가 200인지 확인합니다.
	if assert.Equal(t, http.StatusOK, responseRecorder.Code) {
		t.Log("deleteAccount 전 LogIn 테스트 통과")
	} else {
		t.Log("deleteAccount 전 LogIn 테스트 실패")
	}

	t.Log("Session-Key:" + responseRecorder.Body.String())
	var responseBody map[string]string
	json.Unmarshal(responseRecorder.Body.Bytes(), &responseBody)
	sessionKey := responseBody["sessionKey"]

	httpRequest, _ = http.NewRequest("POST", "/deleteAccount", nil)
	httpRequest.Header.Set("Session-Key", sessionKey)
	responseRecorder = httptest.NewRecorder()
	ginEngine.ServeHTTP(responseRecorder, httpRequest)

	// 응답 상태 코드가 200인지 확인합니다.
	if assert.Equal(t, http.StatusOK, responseRecorder.Code) {
		t.Log("deleteAccount 테스트 통과")
	} else {
		t.Log("deleteAccount 테스트 실패")
	}
}
