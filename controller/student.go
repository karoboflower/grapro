package controller

import (
	"gra-pro/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// StudentGET 获取学生个人页面
func StudentGET(c *gin.Context) {
	var student database.Student
	id := c.Param("id")

	if database.DB.First(&student, id); student.ID == id {
		c.HTML(http.StatusOK, "student/profile.tmpl", gin.H{"info": student})
	}
	c.HTML(http.StatusOK, "student/profile.tmpl", gin.H{"id": id})
}
