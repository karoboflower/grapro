package database

import (
	"time"
)

// User 用户模型
type User struct {
	ID        string `gorm:"type:char(13);primary_key;not null;unique" form:"id" json:"id" binding:"required"`
	Email     string `gorm:"type:char(32);not null" form:"email" json:"email" binding:"required"`
	Password  string `gorm:"type:char(32);not null" form:"password" json:"password" binding:"required"`
	Role      string `gorm:"type:char(15);not null" form:"role" json:"role" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// Exists 检测用户是否存在
func (u User) Exists() bool {
	var user User
	if DB.Where("id = ?", u.ID).First(&user); user.ID != u.ID {
		return true
	}
	return false
}
