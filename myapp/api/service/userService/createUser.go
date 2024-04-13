// myapp/api/service/userService/createUser.go
package userService

import (
	"myapp/api/models"
	"myapp/internal/db"
	"myapp/internal/password"
)

func CreateUser(user models.User) error {
	hashedPassword, err := password.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	dbManager := db.GetDBManager()
	err = dbManager.Create(&user)
	return err
}
