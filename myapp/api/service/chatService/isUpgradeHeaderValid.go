// myapp/api/service/chatService/isUpgradeHeaderValid.go
package chatService

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsUpgradeHeaderValid(c *gin.Context) bool {
	upgradeHeader := c.GetHeader("Upgrade")
	if upgradeHeader != "websocket" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Upgrade header"})
		return false
	}

	return true
}
