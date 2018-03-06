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
	// j := middleware.JWT{
	// 	SigningKey: []byte("huanglachuan"),
	// }

	// claims := middleware.CustomClaims{
	// 	ID:    "2014051609033",
	// 	Email: "898859616@qq.com",
	// 	StandardClaims: jwt.StandardClaims{
	// 		ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
	// 		Issuer:    "huanglachuan",
	// 	},
	// }

	// token, err := j.CreateToken(claims)

	// if err != nil {
	// 	c.String(http.StatusOK, err.Error())
	// 	c.Abort()
	// }

	// c.Request.Header.Set("Authorization", token)
	// res, err := j.ParseToken(token)

	// if err != nil {
	// 	if err == middleware.TokenExpired {
	// 		newToken, err := j.RefreshToken(token)
	// 		if err != nil {
	// 			c.String(http.StatusOK, err.Error())
	// 		} else {
	// 			c.String(http.StatusOK, "Refresh Token <-------------------------------->"+newToken)
	// 		}
	// 	} else {
	// 		c.String(http.StatusOK, err.Error())
	// 	}
	// } else {
	// 	c.JSON(http.StatusOK, res)
	// }
}

// LoginPOST 用户登录处理
func LoginPOST(c *gin.Context) {
	var json database.User
	var saltInst user.Salt

	if database.DB.First(&json, c.Params.ByName("id")); json.ID == c.Params.ByName("id") {
		if user.GetSalt(&saltInst) {
			h := md5.New()
			io.WriteString(h, c.Params.ByName("password"))
			pwmd5 := fmt.Sprintf("%x", h.Sum(nil))
			io.WriteString(h, saltInst.Salt1)
			io.WriteString(h, c.Params.ByName("id"))
			io.WriteString(h, saltInst.Salt2)
			io.WriteString(h, json.Email)
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

				if err != nil {
					c.JSON(http.StatusOK, err.Error())
					c.Abort()
				}

				c.Header("Authorization", token)

				c.JSON(http.StatusOK, gin.H{"message": "Login Success"})
			}
		}
	}
	// var json database.User
	// var LoggedUser database.User
	// var saltInst user.Salt

	// if c.BindJSON(&json) == nil {
	// 	if database.DB.First(&LoggedUser, json.ID); LoggedUser.ID != json.ID {
	// 		if user.GetSalt(&saltInst) {
	// 			h := md5.New()
	// 			io.WriteString(h, json.Password)
	// 			pwmd5 := fmt.Sprintf("%x", h.Sum(nil))
	// 			io.WriteString(h, saltInst.Salt1)
	// 			io.WriteString(h, json.ID)
	// 			io.WriteString(h, saltInst.Salt2)
	// 			io.WriteString(h, LoggedUser.Email)
	// 			io.WriteString(h, pwmd5)
	// 			last := fmt.Sprintf("%x", h.Sum(nil))
	// 			if LoggedUser.Password == last {

	// 			}
	// 		}
	// 	} else {
	// 		c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/register")
	// 	}
	// }
}
