package database

import (
	"time"
)

// Student 学生信息
type Student struct {
	ID              string    `gorm:"type:char(13);primary_key;not null;unique" form:"id" json:"id" binding:"required"`
	Name            string    `gorm:"type:char(24);not null" form:"name" json:"name" binding:"required"`
	Sex             string    `gorm:"type:char(3);not null" form:"sex" json:"sex" binding:"required"`
	Birthdate       time.Time `gorm:"not null" form:"birthdate" json:"birthdate" binding:"required"`
	Ethnic          string    `gorm:"type:char(20);not null" form:"ethnic" json:"ethnic" binding:"required"`
	IDNumber        string    `gorm:"type:char(18);not null;unique_index" form:"idnumber" json:"idnumber" binding:"required"`
	PoliticalStatus string    `gorm:"type:char(20);not null" form:"politicalstatus" json:"politicalstatus" binding:"required"`
	College         string    `gorm:"type:char(20);not null" form:"college" json:"college" binding:"required"`
	Profession      string    `gorm:"type:char(20);not null" form:"profession" json:"profession" binding:"required"`
	Grade           string    `gorm:"type:char(20);not null" form:"grade" json:"grade" binding:"required"`
	Class           string    `gorm:"type:char(20);not null" form:"class" json:"class" binding:"required"`
	Phone           string    `gorm:"type:char(20);not null" form:"phone" json:"phone" binding:"required"`
}

// Exists 检测学生表是否已经存在该记录,存在返回true
func (s Student) Exists() bool {
	var student Student

	if DB.Where("id = ?", s.ID).First(&student); student.ID == s.ID {
		return true
	}
	return false
}
