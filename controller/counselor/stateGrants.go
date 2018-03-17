package counselor

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetStateGrants 辅导员查询国家助学金信息
func GetStateGrants(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// PostStateGrants 辅导员筛选国家助学金信息
func PostStateGrants(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}
