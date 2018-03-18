package database

import "time"

// Counselor 辅导员信息
type Counselor struct {
	ID        string `gorm:"type:char(13);primary_key;not null;unique" form:"id" json:"id" binding:"required"`
	Name      string `gorm:"type:char(24);not null" form:"name" json:"name" binding:"required"`
	College   string `gorm:"type:char(20);not null" form:"college" json:"college" binding:"required"`
	Grade     string `gorm:"type:char(20);not null" form:"grade" json:"grade" binding:"required"`
	Class     string `gorm:"type:char(20);not null" form:"class" json:"class" binding:"required"`
	Phone     string `gorm:"type:char(20);not null" form:"phone" json:"phone" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// Exists .
func (c Counselor) Exists() bool {
	var counselor Counselor

	if DB.Where("id = ?", c.ID).First(&counselor); counselor.ID == c.ID {
		return true
	}
	return false
}
