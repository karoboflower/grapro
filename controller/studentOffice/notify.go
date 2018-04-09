// Package studentOffice 通知管理
package studentOffice

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PostNotify 提交通知
func PostNotify(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// DeleteNotify 删除通知
func DeleteNotify(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}
