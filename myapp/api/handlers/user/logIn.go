// myapp/api/handlers/user/logIn.go
package user

import (
	"encoding/json"
	"myapp/api/models"
	"myapp/internal/db"
	"myapp/internal/password"
	"myapp/internal/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO
// Login 성공시에 session key를 생성하고, 이를 반환하도록 수정하세요.

// LogIn 함수는 사용자가 제공한 이메일 주소와 비밀번호를 검증하여 로그인합니다.
func LogIn(ginContext *gin.Context) {
	// 사용자가 제공한 로그인 정보를 담을 LoginInfo 구조체를 선언합니다.
	var loginInfo struct {
		EmailAddress string `json:"emailAddress"`
		Password     string `json:"password"`
	}

	// 요청 본문에서 로그인 정보를 읽어 LoginInfo 구조체에 저장합니다.
	// 본문을 읽는 도중 오류가 발생하면 400 에러를 반환합니다.
	err := json.NewDecoder(ginContext.Request.Body).Decode(&loginInfo)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 데이터베이스 연결을 가져옵니다.
	dbManager := db.GetDBManager()

	// 사용자 정보를 담을 User 구조체를 선언합니다.
	var user models.User

	// 사용자가 제공한 이메일 주소로 데이터베이스에서 사용자를 찾습니다.
	// 사용자를 찾는 도중 오류가 발생하면 500 에러를 반환합니다.
	err = dbManager.Read(&user, "email_address", loginInfo.EmailAddress)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
		return
	}

	// 사용자가 제공한 비밀번호와 데이터베이스에 저장된 해시된 비밀번호를 비교합니다.
	// 비밀번호가 일치하지 않으면 401 에러를 반환합니다.
	if !password.CheckPasswordHash(loginInfo.Password, user.Password) {
		ginContext.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// 세션 키를 생성합니다.
	sessionKey, err := session.GenerateRandomSessionKey()
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate session key"})
		return
	}

	// 세션 키를 캐시에 저장합니다.
	sessionManager := session.GetSessionManager()
	err = sessionManager.SetSession(sessionKey, user.EmailAddress)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session key"})
		return
	}

	// 로그인이 성공했으므로 유저에게 세션 키를 반환합니다.
	ginContext.JSON(http.StatusOK, gin.H{"sessionKey": sessionKey})
}
