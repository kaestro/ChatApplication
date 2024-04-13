// myapp/api/models/loginInfo.go
package models

type LoginInfo struct {
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
}

func NewLoginInfo(emailAddress, password string) LoginInfo {
	return LoginInfo{
		EmailAddress: emailAddress,
		Password:     password,
	}
}
