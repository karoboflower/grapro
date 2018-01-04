package controller

import (
	"crypto/md5"
	"fmt"
	"gra-pro/config/user"
	"gra-pro/database"
	"gra-pro/database/models"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterUser 注册用户
func RegisterUser(c *gin.Context) {
	var json models.User
	var saltInst user.Salt

	if c.BindJSON(&json) == nil {
		if user.GetSalt(&saltInst) {
			// 检测是否存在记录
			if !database.DB.NewRecord(json) {
				h := md5.New()
				io.WriteString(h, json.Password)
				pwmd5 := fmt.Sprintf("%x", h.Sum(nil))
				io.WriteString(h, saltInst.Salt1)
				io.WriteString(h, json.Email)
				io.WriteString(h, saltInst.Salt2)
				io.WriteString(h, pwmd5)
				last := fmt.Sprintf("%x", h.Sum(nil))
				json.Password = last
				if dbc := database.DB.Create(&json); dbc.Error != nil {
					c.JSON(http.StatusOK, gin.H{"status": "您已经注册，请登录！"})
				} else {
					c.JSON(http.StatusOK, gin.H{"status": "您已成功注册！"})
				}
			}
		}
	}
}
