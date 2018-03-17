package counselor

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCounselor 获取辅导员个人信息页面
func GetCounselor(c *gin.Context) {
	c.HTML(http.StatusOK, "counselor/profile.tmpl", gin.H{"message": "counselor profile"})
}

// PostCounselor 提交辅导员个人信息页面
func PostCounselor(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}

// PutCounselor 修改辅导员个人信息页面
func PutCounselor(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": c.Request.URL.Path})
}
