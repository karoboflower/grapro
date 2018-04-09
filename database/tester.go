package database

import (
	"time"
)

// Tester 重构测试
type Tester struct {
	ID        int64  `gorm:"primary_key;AUTO_INCREMENT"`
	StudentID string `gorm:"type:char(13);not null"`
	Status    string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
