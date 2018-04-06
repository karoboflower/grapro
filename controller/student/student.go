package student

import (
	"gra-pro/database"
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// GetProfile 获取学生个人信息
func GetProfile(c *gin.Context) {
	id := c.Param("id")

	c.HTML(http.StatusOK, "student/profile.tmpl", gin.H{"ID": id})
}

// PostProfile 提交学生个人信息
func PostProfile(c *gin.Context) {
	var form database.Student

	if err := c.ShouldBindWith(&form, binding.FormPost); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": err.Error()})
		return
	}

	if !database.DB.First(&database.Student{}, "id = ?", form.ID).RecordNotFound() {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": "已录入个人信息，无需重复提交！"})
		return
	}

	if dbe := database.DB.Create(&form); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 0})
}

// PutProfile 修改学生个人信息
func PutProfile(c *gin.Context) {
	var form database.Student
	var student database.Student

	var err error
	var dbe *gorm.DB
	if err = c.ShouldBindWith(&form, binding.FormPost); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": err.Error()})
		return
	}

	if dbe = database.DB.First(&student, form.ID); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}

	if dbe = database.DB.Model(&student).Updates(form); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 0})
}
