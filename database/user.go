package database

import (
	"time"
)

// User .
type User struct {
	ID        string `gorm:"type:char(13);primary_key;not null;unique" form:"id" json:"id" binding:"required"`
	Email     string `gorm:"type:char(32);not null;unique" form:"email" json:"email" binding:"required"`
	Password  string `gorm:"type:char(32);not null" form:"password" json:"password" binding:"required"`
	Role      string `gorm:"type:char(1);not null" form:"role" json:"role" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// Exists 检测用户表是否存在该记录,存在返回true
func (u User) Exists() bool {
	var user User

	if DB.Where("id = ?", u.ID).First(&user); user.ID == u.ID {
		return true
	}
	return false
}
