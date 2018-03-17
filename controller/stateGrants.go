package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// StateGrantsGET 学生获取国家助学金申请信息
func StateGrantsGET(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// StateGrantsPOST 学生提交国家助学金申请信息
func StateGrantsPOST(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// StateGrantsPUT 学生修改国家助学金申请信息
func StateGrantsPUT(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// StateGrantsDEL 学生删除国家助学金申请信息
func StateGrantsDEL(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}
