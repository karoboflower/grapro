package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// StudentGET 获取学生个人页面
func StudentGET(c *gin.Context) {
	c.HTML(http.StatusOK, "student/profile.tmpl", gin.H{"message": "student profile"})
}
