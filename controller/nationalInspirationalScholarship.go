// Package controller 国家励志奖学金
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NISGET 学生获取国家励志奖学金申请信息
func NISGET(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// NISPOST 学生提交国家励志奖学金申请信息
func NISPOST(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// NISPUT 学生修改国家励志奖学金申请信息
func NISPUT(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// NISDEL 学生删除国家励志奖学金申请信息
func NISDEL(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}
