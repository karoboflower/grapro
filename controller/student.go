package controller

import (
	"fmt"
	"gra-pro/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// StudentGET 获取学生个人页面
func StudentGET(c *gin.Context) {
	fmt.Println("Student get")
	var student database.Student
	id := c.Param("id")

	if database.DB.First(&student, id); student.ID != id {
		student = database.Student{
			ID: id,
		}
	}
	c.HTML(http.StatusOK, "student/profile.tmpl", gin.H{"info": student})
}

// StudentPOST 提交学生个人信息
func StudentPOST(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// StudentPUT 修改学生个人信息
func StudentPUT(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}
