package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// StudentOfficeGET 获取学生处详细信息页面
func StudentOfficeGET(c *gin.Context) {
	c.HTML(http.StatusOK, "studentOffice/profile.tmpl", gin.H{"message": "student office profile"})
}

// StudentOfficePOST 提交学生处详细信息页面
func StudentOfficePOST(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// StudentOfficePUT 修改学生处详细信息页面
func StudentOfficePUT(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}
