package utilities

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetNotify 获取通知信息
func GetNotify(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}
