package database

import (
	"time"

	"github.com/lib/pq"
)

// KindnessScholarship 应善良助学金
type KindnessScholarship struct {
	StudentID string         `gorm:"type:char(13);primary_key;not null;unique_index"`
	Accessory pq.StringArray `gorm:"type:varchar(255);not null"`
	Status    int            `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
