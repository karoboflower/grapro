package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterGET 返回用户注册视图
func RegisterGET(c *gin.Context) {
	c.HTML(http.StatusOK, "common/register.tmpl", gin.H{})
}

// RegisterPOST 用户注册处理
func RegisterPOST(c *gin.Context) {
	// var json database.User
	// var saltInst user.Salt

	id := c.PostForm("id")
	c.JSON(http.StatusOK, gin.H{"message": "hello " + id})

	// if err := c.BindJSON(&json); err == nil && user.GetSalt(&saltInst) {
	// 	// 检测是否存在记录
	// 	if json.Exists() {
	// 		h := md5.New()
	// 		io.WriteString(h, json.Password)
	// 		pwmd5 := fmt.Sprintf("%x", h.Sum(nil))
	// 		io.WriteString(h, saltInst.Salt1)
	// 		io.WriteString(h, json.ID)
	// 		io.WriteString(h, saltInst.Salt2)
	// 		io.WriteString(h, json.Email)
	// 		io.WriteString(h, pwmd5)
	// 		last := fmt.Sprintf("%x", h.Sum(nil))
	// 		json.Password = last
	// 		database.DB.Create(&json)
	// 		database.AuthEnforcer.AddPermissionForUser(json.ID, "data/*", "(GET)|(POST)")
	// 		// if database.AuthEnforcer.Enforce(json.ID, "data/something", "GET") {
	// 		// 	fmt.Println("access confirmed")
	// 		// }
	// 		database.AuthEnforcer.LoadPolicy()
	// 		c.JSON(http.StatusOK, gin.H{"status": "您已成功注册！"})
	// 	} else {
	// 		c.JSON(http.StatusOK, gin.H{"status": "您已经注册，请登录！"})
	// 	}
	// }
}
