package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// StudentOfficeGET 获取学生处详细信息页面
func StudentOfficeGET(c *gin.Context) {
	c.HTML(http.StatusOK, "studentOffice/profile.tmpl", gin.H{"message": "student office profile"})
}
