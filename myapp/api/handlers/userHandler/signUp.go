// myapp/api/handlers/user/signup.go

package userHandler

import (
	"myapp/api/models"
	"myapp/api/service/generalService"
	"myapp/api/service/userService"
	"net/http"

	"github.com/gin-gonic/gin"
)

// request type: POST
// Headers: None
// Body: User { userName, emailAddress, password }
func SignUp(ginContext *gin.Context) {
	var user models.User

	err := generalService.DecodeUserFromBody(ginContext, &user)
	if err != nil {
		ginContext.Error(err)
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = userService.CreateUser(user)
	if err != nil {
		ginContext.Error(err)
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 사용자 정보를 반환하며 201 Created 상태 코드를 반환합니다.
	ginContext.JSON(http.StatusCreated, user)
}
