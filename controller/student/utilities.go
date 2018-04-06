package student

import (
	"gra-pro/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCounselor 获取辅导员信息
func GetCounselor(c *gin.Context) {
	grade := c.DefaultQuery("grade", "17级")
	college := c.DefaultQuery("college", "新闻学院")
	var counselors []database.Counselor

	if dbe := database.DB.Where("grade = ? AND college = ?", grade, college).Find(&counselors); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 0, "counselors": counselors})
}
