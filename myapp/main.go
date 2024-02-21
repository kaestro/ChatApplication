// main.go

package main

import (
	"fmt"

	"myapp/api/handlers/user"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/signup", user.SignUp)
	r.POST("/login", user.LogIn)

	r.Run()

}
