package studentOffice

import (
	"gra-pro/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
)

// GetStudentOffice 获取学生处详细信息页面
func GetStudentOffice(c *gin.Context) {
	var so database.StudentOffice
	id := c.Param("id")

	reqModify, _ := strconv.ParseBool(c.DefaultQuery("ReqModify", "false"))
	if reqModify {
		c.HTML(http.StatusOK, "counselor/profileForm.tmpl", gin.H{"ID": id})
		return
	}

	if database.DB.First(&so, id).RecordNotFound() {
		so = database.StudentOffice{
			StudentOfficeID: id,
		}
	}

	c.HTML(http.StatusOK, "studentOffice/profile.tmpl", gin.H{"status": 0, "info": so})
}

// PostStudentOffice 提交学生处详细信息页面
func PostStudentOffice(c *gin.Context) {
	var form database.StudentOffice
	var so database.StudentOffice

	if err := c.ShouldBindWith(&form, binding.FormPost); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": err.Error()})
		return
	}

	if database.DB.First(&database.StudentOffice{}, "id = ?", form.StudentOfficeID).RecordNotFound() {
		if dbe := database.DB.Create(&form); dbe.Error != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
			return
		}
	} else {
		if dbe := database.DB.First(&so, form.StudentOfficeID); dbe.Error != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
			return
		}
		if dbe := database.DB.Model(&so).Updates(form); dbe.Error != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": 0})
}
