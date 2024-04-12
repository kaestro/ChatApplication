// myapp/api/models/loginInfo.go
package models

type LoginInfo struct {
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
}
