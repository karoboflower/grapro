package controller

import (
	"crypto/md5"
	"fmt"
	"gra-pro/config/user"
	"gra-pro/database"
	"gra-pro/middleware"
	"io"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

// LoginGET 返回用户登录视图
func LoginGET(c *gin.Context) {
	// c.HTML(http.StatusOK, "common/login.tmpl", gin.H{
	// 	"message": "huanglachuan",
	// })
	j := middleware.JWT{
		[]byte("test"),
	}

	claims := middleware.CustomClaims{
		1,
		"awh521",
		"1044176017@qq.com",
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.String(http.StatusOK, err.Error())
		c.Abort()
	}

	c.String(http.StatusOK, token+"----------------------------</br>")
	res, err := j.ParseToken(token)

	if err != nil {
		if err == middleware.TokenExpired {
			newToken, err := j.RefreshToken(token)
			if err != nil {
				c.String(http.StatusOK, err.Error())
			} else {
				c.String(http.StatusOK, newToken)
			}
		} else {
			c.String(http.StatusOK, err.Error())
		}
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// LoginPOST 用户登录处理
func LoginPOST(c *gin.Context) {
	var json database.User
	var LoggedUser database.User
	var saltInst user.Salt

	if c.BindJSON(&json) == nil {
		if database.DB.First(&LoggedUser, json.ID); LoggedUser.ID != json.ID {
			if user.GetSalt(&saltInst) {
				h := md5.New()
				io.WriteString(h, json.Password)
				pwmd5 := fmt.Sprintf("%x", h.Sum(nil))
				io.WriteString(h, saltInst.Salt1)
				io.WriteString(h, json.ID)
				io.WriteString(h, saltInst.Salt2)
				io.WriteString(h, LoggedUser.Email)
				io.WriteString(h, pwmd5)
				last := fmt.Sprintf("%x", h.Sum(nil))
				if LoggedUser.Password == last {

				}
			}
		} else {
			c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/register")
		}
	}
}
