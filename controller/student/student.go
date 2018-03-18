package student

import (
	"gra-pro/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProfile 获取学生个人信息
func GetProfile(c *gin.Context) {
	var student database.Student
	id := c.Param("id")

	if database.DB.First(&student, id); student.ID != id {
		student = database.Student{
			ID: id,
		}
	}

	c.HTML(http.StatusOK, "student/profile.tmpl", gin.H{"info": student})
}

// PostProfile 提交学生个人信息
func PostProfile(c *gin.Context) {
	// var form database.Student
	// var saltInst user.Salt

	// if err := c.ShouldBindWith(&form, binding.FormPost); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": err.Error()})
	// }

	// if form.Exists(){

	// }
}

// PutProfile 修改学生个人信息
func PutProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}
