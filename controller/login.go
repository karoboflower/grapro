package controller

import (
	"crypto/md5"
	"fmt"
	"gra-pro/config/auth"
	"gra-pro/config/user"
	"gra-pro/database"
	"gra-pro/middleware"
	"io"
	"net/http"
	"time"

	"github.com/dchest/captcha"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// GetLogin 返回用户登录视图
func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "common/login.tmpl", gin.H{"captcha": captcha.New()})
}

// PostLogin 用户登录处理
func PostLogin(c *gin.Context) {
	var json database.User
	var saltInst user.Salt
	var secret auth.Secret
	var result bool

	if !captcha.VerifyString(c.PostForm("captchaID"), c.PostForm("captchaSolution")) {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": "验证码错误", "captcha": captcha.New()})
		return
	}

	if dbe := database.DB.First(&json, c.PostForm("id")); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error(), "captcha": captcha.New()})
		return
	}

	if secret, result = auth.GetSignKey(); !(result && user.GetSalt(&saltInst)) {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": "获取用户生成信息失败！", "captcha": captcha.New()})
		return
	}

	h := md5.New()
	io.WriteString(h, c.PostForm("password"))
	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))
	io.WriteString(h, saltInst.Salt1)
	io.WriteString(h, c.PostForm("id"))
	io.WriteString(h, saltInst.Salt2)
	io.WriteString(h, pwmd5)
	last := fmt.Sprintf("%x", h.Sum(nil))

	if json.Password != last {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": "密码错误！", "captcha": captcha.New()})
		return
	}

	j := middleware.JWT{
		SigningKey: []byte(secret.SignKey),
	}

	claims := middleware.CustomClaims{
		ID:    json.UserID,
		Email: json.Email,
		Role:  json.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
			Issuer:    "ThePupilOfTheOcean",
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": err.Error(), "captcha": captcha.New()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 0, "Authrization": token, "Redirect": "/auth/" + json.Role + "/" + json.UserID + "/profile"})
}
