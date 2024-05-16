// myapp/internal/middleware/errorHandling.go:
package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // process request

		// Check for any errors that occurred in the request
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				log.Println(e.Err.Error()) // log error details
			}

			// Return a 500 status code and the errors
			c.JSON(http.StatusInternalServerError, gin.H{"errors": c.Errors.Errors()})
			return
		}
	}
}
