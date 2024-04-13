// myapp/api/models/loginInfo.go
package models

type LoginInfo struct {
	EmailAddress   string `json:"emailAddress"`
	Password       string `json:"password"`
	LoginSessionID string `json:"loginSessionID"`
}

func NewLoginInfo(emailAddress, password string) LoginInfo {
	return LoginInfo{
		EmailAddress:   emailAddress,
		Password:       password,
		LoginSessionID: "",
	}
}
