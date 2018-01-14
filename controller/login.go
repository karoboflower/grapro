package controller

import (
	"gra-pro/config/user"
	"gra-pro/database"
	"log"

	"github.com/gin-gonic/gin"
)

// UserPOST 用户登录处理
func UserPOST(c *gin.Context) {
	var json database.User
	var saltInst user.Salt

	if user.GetSalt(&saltInst) {
		if c.BindJSON(&json) == nil {
			// h := md5.New()
			// io.WriteString(h, json.Password)
			// pwmd5 := fmt.Sprintf("%x", h.Sum(nil))
			// io.WriteString(h, saltInst.Salt1)
			// io.WriteString(h, json.Email)
			// io.WriteString(h, saltInst.Salt2)
			// io.WriteString(h, pwmd5)
			// last := fmt.Sprintf("%x", h.Sum(nil))

			// h.Write([]byte(json.Password))
			// x := h.Sum(nil)
			// y := make([]byte, 32)
			// hex.Encode(y, x)
			// fmt.Println(y)
			// if json.ID == "manu" && json.Password == "123" {
			// 	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			// } else {
			// 	c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			// }
		}
	} else {
		log.Println("读取密码盐失败！")
	}
}
