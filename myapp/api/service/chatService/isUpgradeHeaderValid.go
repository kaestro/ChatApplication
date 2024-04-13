// myapp/api/service/chatService/isUpgradeHeaderValid.go
package chatService

import (
	"github.com/gin-gonic/gin"
)

func IsUpgradeHeaderValid(c *gin.Context) bool {
	upgradeHeader := c.GetHeader("Upgrade")
	return upgradeHeader == "websocket"
}
