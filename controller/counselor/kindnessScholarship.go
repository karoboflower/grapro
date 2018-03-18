package counselor

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetViewKindnessScholarship 辅导员查询应善良助学金
func GetViewKindnessScholarship(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// PostViewKindnessScholarship 辅导员筛选应善良助学金
func PostViewKindnessScholarship(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}
