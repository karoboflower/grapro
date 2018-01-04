package models

import (
	"gra-pro/config/user"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// User 用户模型
type User struct {
	ID        string `gorm:"type:char(13);primary_key;not null;unique" form:"id" json:"id" binding:"required"`
	Email     string `gorm:"type:char(25);unique_index;not null" form:"email" json:"email" binding:"required"`
	Password  string `gorm:"type:char(32);not null" form:"password" json:"password" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// UserPOST 用户登录处理
func UserPOST(c *gin.Context) {
	var json User
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
