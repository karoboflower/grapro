// Package studentOffice 通知管理
package studentOffice

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetNotify 填写通知页面
func GetNotify(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// PostNotify 提交通知
func PostNotify(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// PutNotify 修改通知
func PutNotify(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// DeleteNotify 删除通知
func DeleteNotify(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}
