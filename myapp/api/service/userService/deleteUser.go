// myapp/api/service/userService/deleteUserService.go
package userService

import (
	"myapp/api/models"
	"myapp/internal/db"
	"myapp/internal/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUserBySessionKey(userSessionKey string, ginContext *gin.Context) error {
	sessionManager := session.GetLoginSessionManager()
	emailAddress, err := sessionManager.GetSession(userSessionKey)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get session key"})
		return nil
	}

	err = DeleteUserByEmailAddress(emailAddress)
	return err
}

func DeleteUserByEmailAddress(email_address string) error {
	dbManager := db.GetDBManager()

	var user models.User

	err := dbManager.Read(&user, "email_address", email_address)
	if err != nil {
		return err
	}

	err = dbManager.Delete(&user)
	return err
}
