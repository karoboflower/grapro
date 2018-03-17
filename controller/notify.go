// Package controller 通知管理
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NotifyGET 填写通知页面
func NotifyGET(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// NotifyPOST 提交通知
func NotifyPOST(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// NotifyPUT 修改通知
func NotifyPUT(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// NotifyDEL 删除通知
func NotifyDEL(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}
