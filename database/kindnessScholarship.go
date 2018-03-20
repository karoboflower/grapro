package database

import "time"

// KindnessScholarship 应善良助学金
type KindnessScholarship struct {
	ID        int `gorm:"primary_key;AUTO_INCREMENT"`
	status    int `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
