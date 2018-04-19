package database

import (
	"time"
)

// Application 申请起止记录
type Application struct {
	ApplicationID uint64    `gorm:"primary_key;AUTO_INCREMENT"`
	SSOID         string    `gorm:"type:char(13);not null"`
	Type          int       `gorm:"not null"`
	BeginTime     time.Time `gorm:"not null"`
	EndTime       time.Time `gorm:"not null"`
	Status        string    `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
}
