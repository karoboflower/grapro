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
	"github.com/gin-gonic/gin/binding"
)

// GetRegister 返回用户注册视图
func GetRegister(c *gin.Context) {
	captcha := struct {
		CaptchaID string
	}{
		captcha.New(),
	}

	c.HTML(http.StatusOK, "common/register.tmpl", gin.H{"captcha": captcha})
}

// PostRegister 用户注册处理
func PostRegister(c *gin.Context) {
	var form database.User
	var saltInst user.Salt

	if !captcha.VerifyString(c.PostForm("captchaID"), c.PostForm("captchaSolution")) {
		captcha := struct {
			CaptchaID string
		}{
			captcha.New(),
		}
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": "验证码错误", "captchaID": captcha.CaptchaID})
		return
	}

	if err := c.ShouldBindWith(&form, binding.FormPost); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": err.Error()})
		return
	}

	if form.Exists() {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": "您已经注册，请登录！"})
		return
	}

	var secret auth.Secret
	var result bool
	if secret, result = auth.GetSignKey(); !(result && user.GetSalt(&saltInst)) {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": "获取用户生成信息失败！"})
		return
	}

	h := md5.New()
	io.WriteString(h, form.Password)
	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))
	io.WriteString(h, saltInst.Salt1)
	io.WriteString(h, form.ID)
	io.WriteString(h, saltInst.Salt2)
	io.WriteString(h, pwmd5)
	last := fmt.Sprintf("%x", h.Sum(nil))
	form.Password = last

	if dbe := database.DB.Create(&form); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}

	database.AuthEnforcer.AddPermissionForUser(form.ID, "/auth/"+form.Role+"/"+form.ID+"/*", "(GET)|(POST)|(PUT)|(PATCH)|(DELETE)|(OPTIONS)")
	database.AuthEnforcer.LoadPolicy()

	j := middleware.JWT{
		SigningKey: []byte(secret.SignKey),
	}

	claims := middleware.CustomClaims{
		ID:    form.ID,
		Email: form.Email,
		Role:  form.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
			Issuer:    "ThePupilOfTheOcean",
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 0, "Authrization": token, "Redirect": "/auth/" + form.Role + "/" + form.ID + "/profile"})
}
