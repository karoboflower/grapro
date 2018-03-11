package controller

import (
	"crypto/md5"
	"fmt"
	"gra-pro/config/user"
	"gra-pro/database"
	"gra-pro/middleware"
	"io"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// LoginGET 返回用户登录视图
func LoginGET(c *gin.Context) {
	c.HTML(http.StatusOK, "common/login.tmpl", gin.H{})
}

// LoginPOST 用户登录处理
func LoginPOST(c *gin.Context) {
	var json database.User
	var saltInst user.Salt

	if dbe := database.DB.First(&json, c.PostForm("id")); dbe.Error == nil {
		if user.GetSalt(&saltInst) {
			h := md5.New()
			io.WriteString(h, c.PostForm("password"))
			pwmd5 := fmt.Sprintf("%x", h.Sum(nil))
			io.WriteString(h, saltInst.Salt1)
			io.WriteString(h, c.PostForm("id"))
			io.WriteString(h, saltInst.Salt2)
			io.WriteString(h, pwmd5)
			last := fmt.Sprintf("%x", h.Sum(nil))
			if json.Password == last {
				j := middleware.JWT{
					SigningKey: []byte("secret"),
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
				if err == nil {
					c.SetCookie("Authorization", token, 1, "/", "localhost", true, true)
					c.Redirect(http.StatusMovedPermanently, "/auth/"+c.PostForm("role")+"/"+c.PostForm("id"))
				}
			}
		} else {
			// c.JSON(http.StatusOK, gin.H{"msg": "密码盐"})
		}
	} else {
		// c.JSON(http.StatusOK, gin.H{"msg": dbe.Error})
	}
}
