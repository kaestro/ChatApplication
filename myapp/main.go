// main.go

package main

import (
	"fmt"

	"myapp/api/handlers/userHandler"

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

	r.POST("/signup", userHandler.SignUp)
	r.POST("/login", userHandler.LogIn)
	r.POST("/logout", userHandler.LogOut)
	r.POST("/deleteAccount", userHandler.DeleteAccount)

	r.Run()

}
