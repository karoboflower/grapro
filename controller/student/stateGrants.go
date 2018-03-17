package student

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetStateGrants 学生获取国家助学金申请信息
func GetStateGrants(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// PostStateGrants 学生提交国家助学金申请信息
func PostStateGrants(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// PutStateGrants 学生修改国家助学金申请信息
func PutStateGrants(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// DeleteStateGrants 学生删除国家助学金申请信息
func DeleteStateGrants(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}
