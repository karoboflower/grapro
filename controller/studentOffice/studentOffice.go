package studentOffice

import (
	"gra-pro/database"
	"net/http"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
)

// GetStudentOffice 获取学生处详细信息页面
func GetStudentOffice(c *gin.Context) {
	var so database.StudentOffice
	id := c.Param("id")

	if database.DB.First(&so, id); so.ID != id {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": "还没有填写信息！"})
		return
	}

	c.HTML(http.StatusOK, "studentOffice/profile.tmpl", gin.H{"status": 0, "info": so})
}

// PostStudentOffice 提交学生处详细信息页面
func PostStudentOffice(c *gin.Context) {
	var form database.StudentOffice

	if err := c.ShouldBindWith(&form, binding.FormPost); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": err.Error()})
		return
	}

	if !database.DB.First(&database.StudentOffice{}, "id = ?", form.ID).RecordNotFound() {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": "已录入信息，无需重复提交！"})
		return
	}

	if dbe := database.DB.Create(&form); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 0})
}

// PutStudentOffice 修改学生处详细信息页面
func PutStudentOffice(c *gin.Context) {
	var form database.StudentOffice

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
