// Package studentOffice 通知管理
package studentOffice

import (
	"grapro/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PostNotify 提交通知
func PostNotify(c *gin.Context) {
	var notify database.Notify
	notify.Content = c.PostForm("Content")

	if c.PostForm("NotifyID") == "" {
		if dbe := database.DB.Create(&notify); dbe.Error != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
			return
		}
	} else {
		var temp database.Notify
		NotifyID, _ := strconv.ParseUint(c.PostForm("NotifyID"), 10, 64)
		if dbe := database.DB.First(&temp, NotifyID); dbe.Error != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
			return
		}
		if dbe := database.DB.Model(&temp).Updates(notify); dbe.Error != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": 0})
}

// DeleteNotify 删除通知
func DeleteNotify(c *gin.Context) {
	NotifyID, _ := strconv.ParseUint(c.PostForm("NotifyID"), 10, 64)
	if dbe := database.DB.Where("notify_id = ?", NotifyID).Delete(&database.Notify{}); dbe != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 0})
}
