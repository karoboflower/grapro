package models

import (
	"time"
)

// User 用户模型
type User struct {
	ID        string `gorm:"type:char(13);primary_key;not null;unique"`
	Email     string `gorm:"type:char(25);unique_index;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
