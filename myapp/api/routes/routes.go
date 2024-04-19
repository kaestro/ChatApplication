// /myapp/api/routes/routes.go
package routes

import (
	handler "myapp/api/handlers"
	"myapp/api/handlers/chatHandler"
	"myapp/api/handlers/userHandler"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Method  string
	Pattern string
	Handler gin.HandlerFunc
}

var routes = []Route{
	{"POST", "/signup", userHandler.SignUp},
	{"POST", "/login", userHandler.LogIn},
	{"POST", "/logout", userHandler.LogOut},
	{"POST", "/deleteAccount", userHandler.SignOut},
	{"GET", "/enterChat", chatHandler.EnterChat},
	{"POST", "/enterRoom", chatHandler.EnterRoom},
	{"POST", "/createRoom", chatHandler.CreateRoom},
	{"GET", "/getRoomList", chatHandler.GetRoomList},
	{"Get", "/ping", handler.HandlePing},
}

func SetupRoutes(r *gin.Engine) {
	for _, route := range routes {
		switch route.Method {
		case "GET":
			r.GET(route.Pattern, route.Handler)
		case "POST":
			r.POST(route.Pattern, route.Handler)
		case "PUT":
			r.PUT(route.Pattern, route.Handler)
		case "DELETE":
			r.DELETE(route.Pattern, route.Handler)
		}
	}
}
