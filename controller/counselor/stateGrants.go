package counselor

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetViewStateGrants 辅导员查询国家助学金信息
func GetViewStateGrants(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// PostViewStateGrants 辅导员筛选国家助学金信息
func PostViewStateGrants(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}
