package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CounselorGET 获取辅导员个人页面
func CounselorGET(c *gin.Context) {
	c.HTML(http.StatusOK, "counselor/profile.tmpl", gin.H{"message": "counselor profile"})
}

// CounselorPOST 提交辅导员个人页面
func CounselorPOST(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// CounselorPUT 修改辅导员个人页面
func CounselorPUT(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}
