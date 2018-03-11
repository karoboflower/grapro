package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserGET 获取用户个人信息
func UserGET(c *gin.Context) {
	c.HTML(http.StatusOK, "student/profile.tmpl", gin.H{"message": "student profile"})
}
