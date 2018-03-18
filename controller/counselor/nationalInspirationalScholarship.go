package counselor

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetViewNIS 辅导员查询国家励志奖学金信息
func GetViewNIS(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// PostViewNIS 辅导员筛选国家励志奖学金信息
func PostViewNIS(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}
