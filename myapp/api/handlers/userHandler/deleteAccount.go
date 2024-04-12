// myapp/api/handlers/user/deleteAccount.go
package userHandler

import (
	"fmt"
	"myapp/api/models"
	"myapp/internal/db"
	"myapp/internal/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

// signOut 함수는 사용자를 db에서 삭제합니다.
func DeleteAccount(ginContext *gin.Context) {
	// 세션 키를 요청 헤더에서 읽어옵니다.
	userSessionKey := ginContext.GetHeader("Session-Key")

	// db, 세션 매니저를 가져옵니다.
	sessionManager := session.GetLoginSessionManager()
	dbManager := db.GetDBManager()

	emailAddress, err := sessionManager.GetSession(userSessionKey)
	if err != nil {
		// 세션 키를 가져오는 도중 오류가 발생하면 500 상태 코드를 반환합니다.
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get session key"})
		return
	}

	var user models.User

	err = dbManager.Read(&user, "email_address", emailAddress)
	if err != nil {
		// 사용자 정보를 가져오는 도중 오류가 발생하면 500 상태 코드를 반환합니다.
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user information"})
		return
	}

	// user를 db에서 삭제합니다.
	err = dbManager.Delete(&user)
	if err != nil {
		// 사용자 정보를 삭제하는 도중 오류가 발생하면 500 상태 코드를 반환합니다.
		fmt.Println(err)
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	// 유저 삭제가 성공하면 200 상태 코드를 반환합니다
	ginContext.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
