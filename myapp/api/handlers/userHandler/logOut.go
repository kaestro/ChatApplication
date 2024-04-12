// myapp/api/handlers/user/logout.go
package userHandler

import (
	"myapp/internal/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Logout 함수는 사용자의 세션을 종료합니다.
func LogOut(ginContext *gin.Context) {
	userSessionKey := ginContext.GetHeader("Session-Key")

	sessionManager := session.GetLoginSessionManager()
	err := sessionManager.DeleteSession(userSessionKey)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete session key"})
		return
	}

	ginContext.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
