package controller

import (
	"net/http"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

// ReloadCaptcha 刷新验证码
func ReloadCaptcha(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"captchaID": captcha.New()})
}
