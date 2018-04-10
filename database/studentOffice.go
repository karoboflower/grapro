package database

import "time"

// StudentOffice .
type StudentOffice struct {
	StudentOfficeID string `gorm:"type:char(13);primary_key;not null;unique" form:"id" json:"id" binding:"required"`
	Grade           string `gorm:"type:char(20);not null" form:"grade" json:"grade" binding:"required"`
	College         string `gorm:"type:varchar(100);not null" form:"college" json:"college" binding:"required"`
	Phone           string `gorm:"type:char(20);not null" form:"phone" json:"phone" binding:"required"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `sql:"index"`
}
