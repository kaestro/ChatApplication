// myapp/api/handlers/user/logout.go
package userHandler

import (
	"myapp/api/service/generalService"
	"myapp/api/service/userService"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Request Type: POST
// Headers: Session-Key
// Body: None
func LogOut(ginContext *gin.Context) {
	userSessionKey := generalService.GetSessionKeyFromHeader(ginContext)

	err := userService.DeauthenticateUser(userSessionKey)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete session key"})
		return
	}

	ginContext.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
