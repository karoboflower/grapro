package studentOffice

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetStudentOffice 获取学生处详细信息页面
func GetStudentOffice(c *gin.Context) {
	c.HTML(http.StatusOK, "studentOffice/profile.tmpl", gin.H{"message": "student office profile"})
}

// PostStudentOffice 提交学生处详细信息页面
func PostStudentOffice(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// PutStudentOffice 修改学生处详细信息页面
func PutStudentOffice(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}
