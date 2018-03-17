// Package student 国家励志奖学金
package student

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetNIS 学生获取国家励志奖学金申请信息
func GetNIS(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// PostNIS 学生提交国家励志奖学金申请信息
func PostNIS(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// PutNIS 学生修改国家励志奖学金申请信息
func PutNIS(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// DeleteNIS 学生删除国家励志奖学金申请信息
func DeleteNIS(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}
