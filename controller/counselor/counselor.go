package counselor

import (
	"gra-pro/database"
	"net/http"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
)

// GetCounselor 获取辅导员个人信息页面
func GetCounselor(c *gin.Context) {
	var counselor database.Counselor
	id := c.Param("id")

	if database.DB.First(&counselor, id); counselor.ID != id {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": "还没有填写个人信息！"})
		return
	}

	c.HTML(http.StatusOK, "counselor/profile.tmpl", gin.H{"staus": 0, "info": counselor})
}

// PostCounselor 提交辅导员个人信息页面
func PostCounselor(c *gin.Context) {
	var form database.Counselor

	if err := c.ShouldBindWith(&form, binding.FormPost); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": err.Error()})
		return
	}

	if form.Exists() {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": "已录入信息，无需重复提交！"})
		return
	}

	if dbe := database.DB.Create(&form); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 0})
}

// PutCounselor 修改辅导员个人信息页面
func PutCounselor(c *gin.Context) {
	var form database.Counselor

	if err := c.ShouldBindWith(&form, binding.FormPost); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": err.Error()})
		return
	}

	if dbe := database.DB.Save(&form); dbe != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 0})
}
