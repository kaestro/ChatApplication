// myapp/api/models/user.go
package models

type User struct {
	ID           int    `json:"id"`
	UserName     string `json:"userName" gorm:"type:varchar(255);unique"`
	EmailAddress string `json:"emailAddress" gorm:"type:varchar(255);unique"`
	Password     string `json:"password" gorm:"type:varchar(255)"`
}
