package counselor

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetNIS 辅导员查询国家励志奖学金信息
func GetNIS(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// PostNIS 辅导员筛选国家励志奖学金信息
func PostNIS(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}
