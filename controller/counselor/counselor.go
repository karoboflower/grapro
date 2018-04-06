package counselor

import (
	"gra-pro/database"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
)

// GetCounselor 获取辅导员个人信息页面
func GetCounselor(c *gin.Context) {
	var counselor database.Counselor
	id := c.Param("id")

	reqModify, _ := strconv.ParseBool(c.DefaultQuery("ReqModify", "false"))
	if reqModify {
		c.HTML(http.StatusOK, "counselor/profileForm.tmpl", gin.H{"ID": id})
		return
	}

	if database.DB.First(&counselor, id); counselor.ID != id {
		counselor = database.Counselor{
			ID: id,
		}
		c.HTML(http.StatusOK, "counselor/profile.tmpl", gin.H{"status": 0, "info": counselor})
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

	if !database.DB.First(&database.Counselor{}, "id = ?", form.ID).RecordNotFound() {
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
	var counselor database.Counselor

	var err error
	var dbe *gorm.DB

	if err = c.ShouldBindWith(&form, binding.FormPost); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": err.Error()})
		return
	}

	if dbe = database.DB.First(&counselor, form.ID); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}

	if dbe := database.DB.Model(&counselor).Updates(form); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 0})
}
