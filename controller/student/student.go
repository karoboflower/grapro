package student

import (
	"gra-pro/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// GetProfile 获取学生个人信息
func GetProfile(c *gin.Context) {
	var student database.Student
	id := c.Param("id")

	if database.DB.First(&student, id); student.ID != id {
		student = database.Student{
			ID: id,
		}
		// 学生尚未填写个人信息，个人信息只有ID
		c.HTML(http.StatusOK, "student/profile.tmpl", gin.H{"wrote": false, "info": student})
		return
	}

	c.HTML(http.StatusOK, "student/profile.tmpl", gin.H{"wrote": true, "info": student})
}

// PostProfile 提交学生个人信息
func PostProfile(c *gin.Context) {
	var form database.Student

	if err := c.ShouldBindWith(&form, binding.FormPost); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": err.Error()})
		return
	}

	if form.Exists() {
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
