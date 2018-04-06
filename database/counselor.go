package database

import (
	"time"

	"github.com/lib/pq"
)

// Counselor 辅导员信息
type Counselor struct {
	ID        string         `gorm:"type:char(13);primary_key;not null;unique" form:"id" json:"id" binding:"required"`
	Name      string         `gorm:"type:varchar(50);not null" form:"name" json:"name" binding:"required"`
	Grade     string         `gorm:"type:char(20);not null" form:"grade" json:"grade" binding:"required"`
	College   string         `gorm:"type:varchar(100);not null" form:"college" json:"college" binding:"required"`
	Classes   pq.StringArray `gorm:"type:varchar(255)" form:"class[]" json:"class[]" binding:"required"`
	Phone     string         `gorm:"type:char(11);not null" form:"phone" json:"phone" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
