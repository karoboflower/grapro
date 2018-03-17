// Package controller 应善良助学金
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// KindnessScholarshipGET 学生获取应善良助学金申请信息
func KindnessScholarshipGET(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// KindnessScholarshipPOST 学生获提交善良助学金申请信息
func KindnessScholarshipPOST(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// KindnessScholarshipPUT 学生修改应善良助学金申请信息
func KindnessScholarshipPUT(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// KindnessScholarshipDEL 学生删除应善良助学金申请信息
func KindnessScholarshipDEL(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}
