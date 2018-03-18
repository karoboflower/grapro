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

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// GetLogin 返回用户登录视图
func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "common/login.tmpl", gin.H{})
}

// PostLogin 用户登录处理
func PostLogin(c *gin.Context) {
	var json database.User
	var saltInst user.Salt

	if dbe := database.DB.First(&json, c.PostForm("id")); dbe.Error != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": dbe.Error.Error()})
		return
	}

	var secret auth.Secret
	var result bool
	if secret, result = auth.GetSignKey(); !(result && user.GetSalt(&saltInst)) {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": "获取用户生成信息失败！"})
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
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": "密码错误！"})
		return
	}

	j := middleware.JWT{
		SigningKey: []byte(secret.SignKey),
	}

	claims := middleware.CustomClaims{
		ID:    json.ID,
		Email: json.Email,
		Role:  json.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
			Issuer:    "ThePupilOfTheOcean",
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 0, "Authrization": token})
}
