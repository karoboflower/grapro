package utilities

import (
	"gra-pro/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetIndex 首页
func GetIndex(c *gin.Context) {
	var notifies []database.Notify

	if dbe := database.DB.Find(&notifies); dbe != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}

	c.HTML(http.StatusOK, "common/index.tmpl", gin.H{"notifies": notifies})
}
