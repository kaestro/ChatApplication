// myapp/api/service/userService/testResources.go
package userService

import "myapp/api/models"

var (
	sampleEmailAddress = "sample@gmail.com"
	samplePassword     = "password"
	sampleLoginInfo    = models.LoginInfo{
		EmailAddress: sampleEmailAddress,
		Password:     samplePassword,
	}
	sampleUser = models.User{
		EmailAddress: sampleEmailAddress,
		Password:     samplePassword,
	}
)
