// myapp/internal/middleware/setupMiddleware.go
package middleware

import "github.com/gin-gonic/gin"

func SetupMiddleware(r *gin.Engine) {
	r.Use(ErrorHandlingMiddleware())
}
