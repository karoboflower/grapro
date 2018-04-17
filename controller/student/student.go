package student

import (
	"grapro/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// GetProfile 获取学生个人信息
func GetProfile(c *gin.Context) {
	var student database.Student
	id := c.Param("id")

	reqModify, _ := strconv.ParseBool(c.DefaultQuery("ReqModify", "false"))
	if reqModify {
		c.HTML(http.StatusOK, "student/profileForm.html", gin.H{"ID": id})
		return
	}

	if database.DB.First(&student, id).RecordNotFound() {
		student = database.Student{
			StudentID: id,
		}
	}

	c.HTML(http.StatusOK, "student/profile.tmpl", gin.H{"status": 0, "info": student})
}

// PostProfile 提交学生个人信息
func PostProfile(c *gin.Context) {
	var form database.Student
	var student database.Student

	if err := c.ShouldBindWith(&form, binding.FormPost); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": err.Error()})
		return
	}

	if database.DB.First(&database.Student{}, "student_id = ?", form.StudentID).RecordNotFound() {
		if dbe := database.DB.Create(&form); dbe.Error != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
			return
		}
	} else {
		if dbe := database.DB.First(&student, form.StudentID); dbe.Error != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
			return
		}
		if dbe := database.DB.Model(&student).Updates(form); dbe.Error != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": 0})
}

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
