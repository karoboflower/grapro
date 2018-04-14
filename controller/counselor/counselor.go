package counselor

import (
	"gra-pro/database"
	"net/http"
	"strconv"

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

	if database.DB.First(&counselor, id).RecordNotFound() {
		counselor = database.Counselor{
			CounselorID: id,
		}
	}

	c.HTML(http.StatusOK, "counselor/profile.tmpl", gin.H{"staus": 0, "info": counselor})
}

// PostCounselor 提交辅导员个人信息页面
func PostCounselor(c *gin.Context) {
	var form database.Counselor
	var counselor database.Counselor

	if err := c.ShouldBindWith(&form, binding.FormPost); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": err.Error()})
		return
	}

	if database.DB.First(&database.Counselor{}, "counselor_id = ?", form.CounselorID).RecordNotFound() {
		if dbe := database.DB.Create(&form); dbe.Error != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
			return
		}
	} else {
		if dbe := database.DB.First(&counselor, form.CounselorID); dbe.Error != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
			return
		}
		if dbe := database.DB.Model(&counselor).Updates(form); dbe.Error != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": 0})
}
