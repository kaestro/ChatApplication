// myapp/api/service/userService/deleteUserService.go
package userService

import (
	"errors"
	"myapp/api/models"
	"myapp/internal/db"
	"myapp/internal/session"

	"github.com/gin-gonic/gin"
)

func DeleteUserBySessionKey(userSessionKey string, ginContext *gin.Context) error {
	sessionManager := session.GetLoginSessionManager()
	emailAddress, err := sessionManager.GetSession(userSessionKey)
	if err != nil {
		ginContext.Error(errors.New("user sessionKey = " + userSessionKey + " not found"))
		ginContext.Error(err)
		return err
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
