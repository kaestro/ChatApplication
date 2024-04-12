// myapp/api/handlers/chat/enterRoom.go
package chatHandler

import (
	"myapp/api/service/userService"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleEnterRoom(ginContext *gin.Context, err error) {
	switch err {
	case userService.ErrAlreadyLoggedIn:
		ginContext.JSON(http.StatusConflict, gin.H{"error": "User is already logged in"})
	case userService.ErrUserNotFound:
		ginContext.JSON(http.StatusNotFound, gin.H{"error": "Failed to find user"})
	case userService.ErrInvalidPassword:
		ginContext.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
	case userService.ErrFailedToGenerateSessionKey, userService.ErrFailedToSaveSessionKey:
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process session key"})
	default:
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
}
