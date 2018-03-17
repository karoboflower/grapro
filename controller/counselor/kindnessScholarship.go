package counselor

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetKindnessScholarship 辅导员查询应善良助学金
func GetKindnessScholarship(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// PostKindnessScholarship 辅导员筛选应善良助学金
func PostKindnessScholarship(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}
