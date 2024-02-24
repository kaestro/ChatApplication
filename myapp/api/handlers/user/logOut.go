// myapp/api/handlers/user/logout.go
package user

import (
	"myapp/internal/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Logout 함수는 사용자의 세션을 종료합니다.
func LogOut(ginContext *gin.Context) {
	// 세션 키를 요청 헤더에서 읽어옵니다.
	userSessionKey := ginContext.GetHeader("Session-Key")

	// 세션 매니저를 가져옵니다.
	sessionManager := session.GetSessionManager()

	// 세션 매니저의 DeleteSession 메서드를 호출하여 세션 키를 삭제합니다.
	err := sessionManager.DeleteSession(userSessionKey)
	if err != nil {
		// 세션 키 삭제에 실패하면 500 에러를 반환합니다.
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete session key"})
		return
	}

	// 세션 키 삭제에 성공하면 200 상태 코드를 반환합니다.
	ginContext.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
