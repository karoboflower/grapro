package controller

import (
	"crypto/md5"
	"fmt"
	"gra-pro/config/auth"
	"gra-pro/config/user"
	"gra-pro/database"
	"gra-pro/middleware"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// RegisterGET 返回用户注册视图
func RegisterGET(c *gin.Context) {
	c.HTML(http.StatusOK, "common/register.tmpl", gin.H{})
}

// RegisterPOST 用户注册处理
func RegisterPOST(c *gin.Context) {
	var form database.User
	var saltInst user.Salt

	if err := c.ShouldBindWith(&form, binding.FormPost); err == nil {
		if !form.Exists() {
			if secret, result := auth.GetSignKey(); result && user.GetSalt(&saltInst) {
				h := md5.New()
				io.WriteString(h, form.Password)
				pwmd5 := fmt.Sprintf("%x", h.Sum(nil))
				io.WriteString(h, saltInst.Salt1)
				io.WriteString(h, form.ID)
				io.WriteString(h, saltInst.Salt2)
				io.WriteString(h, pwmd5)
				last := fmt.Sprintf("%x", h.Sum(nil))
				form.Password = last
				if dbe := database.DB.Create(&form); dbe.Error == nil {
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
							ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
							Issuer:    "ThePupilOfTheOcean",
						},
					}

					token, err := j.CreateToken(claims)

					if err == nil {
						c.SetCookie("Authorization", token, 1, "/", "localhost", true, true)
						c.Redirect(http.StatusMovedPermanently, "/auth/"+c.PostForm("role")+"/"+c.PostForm("id")+"/?Authorization="+token)
					}
				} else {
					log.Println(dbe.Error)
				}
			} else {
				// 获取盐失败，会无法入库。可以重定向到登录界面
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"msg": "您已经注册，请登录！"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "获取提交信息错误"})
	}
}
