// myapp/api/service/userService/createUser.go
package userService

import (
	"myapp/api/models"
	"myapp/internal/db"
	"myapp/internal/password"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(user models.User, ginContext *gin.Context) error {
	hashedPassword, err := password.HashPassword(user.Password)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return nil
	}

	user.Password = hashedPassword
	dbManager := db.GetDBManager()
	err = dbManager.Create(&user)
	return err
}
