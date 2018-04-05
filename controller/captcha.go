package controller

import (
	"github.com/gin-gonic/gin"
)

// ReloadCaptcha 刷新验证码
func ReloadCaptcha(c *gin.Context) {
	// update := c.Request.FormValue("update")
	// if update == "true" {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status":    "reloadCaptcha",
	// 		"captchaid": captcha.New(),
	// 	})
	// }
}
