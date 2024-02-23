// myapp/api/handlers/user/logIn.go
package user

import (
	"encoding/json"
	"myapp/api/service/user"
	"myapp/internal/db"
	"myapp/internal/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	// 세션 키를 요청 헤더에서 읽어옵니다.
	userSessionKey := ginContext.GetHeader("Session-Key")

	// 데이터베이스 연결을 가져옵니다.
	dbManager := db.GetDBManager()

	// 세션 매니저를 가져옵니다.
	sessionManager := session.GetSessionManager()

	// LoginService를 생성합니다.
	loginService := user.NewLoginService(dbManager, sessionManager)

	// LoginService의 LogIn 메서드를 호출합니다.
	sessionKey, err := loginService.LogIn(loginInfo.EmailAddress, loginInfo.Password, userSessionKey)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 로그인이 성공했으므로 유저에게 세션 키를 반환합니다.
	ginContext.JSON(http.StatusOK, gin.H{"sessionKey": sessionKey})
}
