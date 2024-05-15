package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestErrorHandlingMiddleware(t *testing.T) {
	// Set up Gin engine with the middleware
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(ErrorHandlingMiddleware())

	// Add a route that always causes an error
	router.GET("/error", func(c *gin.Context) {
		c.Error(gin.Error{
			Err:  http.ErrBodyNotAllowed,
			Type: gin.ErrorTypePublic,
		})
	})

	// Create a request to the route
	req, _ := http.NewRequest(http.MethodGet, "/error", nil)
	resp := httptest.NewRecorder()

	// Process the request
	router.ServeHTTP(resp, req)

	// Check that the middleware returned a 500 status code and the correct error message
	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	assert.Contains(t, resp.Body.String(), http.ErrBodyNotAllowed.Error())

	t.Logf("test errorHandlingMiddleware passed")
}
