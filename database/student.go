package database

import (
	"log"
	"time"
)

// Student 学生信息
type Student struct {
	ID              string `gorm:"type:char(13);primary_key;not null;unique" form:"id" json:"id" binding:"required"`
	Name            string `gorm:"type:char(24);not null;" form:"name" json:"name" binding:"required"`
	Sex             string `gorm:"type:char(3);not null" form:"sex" json:"sex" binding:"required"`
	Birthdate       string `gorm:"not null" form:"birthdate" json:"birthdate" binding:"required"`
	Ethnic          string `gorm:"type:char(20);not null" form:"ethnic" json:"ethnic" binding:"required"`
	IDNumber        string `gorm:"type:char(18);not null;unique_index" form:"idnumber" json:"idnumber" binding:"required"`
	PoliticalStatus string `gorm:"type:char(20);not null" form:"politicalstatus" json:"politicalstatus" binding:"required"`
	College         string `gorm:"type:char(20);not null" form:"college" json:"college" binding:"required"`
	Profession      string `gorm:"type:char(20);not null" form:"profession" json:"profession" binding:"required"`
	Grade           string `gorm:"type:char(20);not null" form:"grade" json:"grade" binding:"required"`
	Class           string `gorm:"type:char(20);not null" form:"class" json:"class" binding:"required"`
	CounselorID     string `gorm:"type:char(13);not null" form:"counselorid" json:"counselorid" binding:"required"`
	Phone           string `gorm:"type:char(20);not null" form:"phone" json:"phone" binding:"required"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `sql:"index"`
}

// Create 创建记录
// @param in 需要创建记录的模型实例指针
func (s Student) Create(in interface{}) (result bool) {
	if dbe := DB.Create(in); dbe.Error != nil {
		log.Println(dbe.Error.Error())
		return false
	}
	return true
}

// Read 读取记录
// @param out 读取到的记录输出参数
func (s Student) Read(out interface{}) (result bool) {
	if dbe := DB.First(out); dbe.Error != nil {
		log.Println(dbe.Error.Error())
		return false
	}
	return true
}

// Update 更新记录
// @param in 需要被更新记录的模型实例指针
// @param out 记录更新后的模型实例指针
func (s Student) Update(in interface{}, out interface{}) (result bool) {
	if dbe := DB.Model(in).Updates(out); dbe.Error != nil {
		log.Println(dbe.Error.Error())
		return false
	}
	return true
}

// Delete 删除记录
// @param in 需要被删除的模型实例指针
func (s Student) Delete(in interface{}) (result bool) {
	if dbe := DB.Delete(in); dbe.Error != nil {
		log.Println(dbe.Error.Error())
		return false
	}
	return true
}
