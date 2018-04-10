package database

import (
	"time"

	"github.com/lib/pq"
)

// KindnessScholarship 应善良助学金
type KindnessScholarship struct {
	KSID      uint32         `gorm:"primary_key;AUTO_INCREMENT"`
	StudentID string         `gorm:"type:char(13);not null"`
	Accessory pq.StringArray `gorm:"type:varchar(255);not null"`
	Status    string         `gorm:"type:char(4);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
