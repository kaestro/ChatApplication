// myapp/api/handlers/user/signup.go

package userHandler

import (
	"encoding/json"
	"myapp/api/models"
	"myapp/internal/db"
	"myapp/internal/password"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignUp 함수는 새로운 사용자를 등록합니다.
func SignUp(ginContext *gin.Context) {
	// 사용자 정보를 담을 User 구조체를 선언합니다.
	var user models.User

	// 요청 본문에서 사용자 정보를 읽어 User 구조체에 저장합니다.
	// 본문을 읽는 도중 오류가 발생하면 400 에러를 반환합니다.
	err := json.NewDecoder(ginContext.Request.Body).Decode(&user)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 사용자가 제공한 비밀번호를 해시합니다.
	// 해싱 도중 오류가 발생하면 500 에러를 반환합니다.
	hashedPassword, err := password.HashPassword(user.Password)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// User 구조체의 비밀번호를 해시로 변경합니다.
	user.Password = hashedPassword

	// 데이터베이스 연결을 가져옵니다.
	dbManager := db.GetDBManager()

	// 사용자 정보를 데이터베이스에 저장합니다.
	// 저장 도중 오류가 발생하면 500 에러를 반환합니다.
	err = dbManager.Create(&user)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 사용자 정보를 반환하며 201 Created 상태 코드를 반환합니다.
	ginContext.JSON(http.StatusCreated, user)
}
