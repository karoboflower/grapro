package utilities

import (
	"gra-pro/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetNotify 获取通知信息
func GetNotify(c *gin.Context) {
	var notify database.Notify
	if dbe := database.DB.First(&notify, c.PostForm("NotifyID")); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 1, "notify": notify})
}
