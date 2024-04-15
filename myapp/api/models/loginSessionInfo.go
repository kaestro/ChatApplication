// myapp/api/models/loginSessionInfo.go
package models

import "myapp/types"

type LoginSessionInfo struct {
	EmailAddress   string `json:"emailAddress"`
	LoginSessionID string `json:"loginSessionID"`
}

func NewLoginSessionInfo(emailAddress string, loginSessionID types.LoginSessionID) LoginSessionInfo {
	return LoginSessionInfo{
		EmailAddress:   emailAddress,
		LoginSessionID: string(loginSessionID),
	}
}
