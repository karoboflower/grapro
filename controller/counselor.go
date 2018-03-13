package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CounselorGET 获取辅导员个人页面
func CounselorGET(c *gin.Context) {
	c.HTML(http.StatusOK, "counselor/profile.tmpl", gin.H{"message": "counselor profile"})
}
