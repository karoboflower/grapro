package database

import (
	"time"

	"github.com/lib/pq"
)

// NIS 国家励志奖学金
type NIS struct {
	StudentID string         `gorm:"type:char(13);primary_key;not null;unique_index"`
	Accessory pq.StringArray `gorm:"type:varchar(255);not null"`
	Status    int            `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
