// myapp/api/handlers/user/deleteAccount.go
package userHandler

import (
	"fmt"
	"myapp/api/service/userService"
	"net/http"

	"github.com/gin-gonic/gin"
)

// signOut 함수는 사용자를 db에서 삭제합니다.
func SignOut(ginContext *gin.Context) {
	// 세션 키를 요청 헤더에서 읽어옵니다.
	userSessionKey := ginContext.GetHeader("Session-Key")

	err := userService.DeleteUserBySessionKey(userSessionKey, ginContext)
	if err != nil {
		fmt.Println(err)
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	ginContext.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
