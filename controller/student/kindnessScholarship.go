// Package student 应善良助学金
package student

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetKindnessScholarship 学生获取应善良助学金申请信息
func GetKindnessScholarship(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// PostKindnessScholarship 学生获提交善良助学金申请信息
func PostKindnessScholarship(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// PutKindnessScholarship 学生修改应善良助学金申请信息
func PutKindnessScholarship(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// DeleteKindnessScholarship 学生删除应善良助学金申请信息
func DeleteKindnessScholarship(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}
