package database

import (
	"time"

	"github.com/lib/pq"
)

// StateGrants 国家助学金
type StateGrants struct {
	StudentID       string         `gorm:"type:char(13);primary_key;not null;unique_index"`
	FSQuestionnaire string         `gorm:"type:varchar(100);not null"`
	Accessory       pq.StringArray `gorm:"type:varchar(255)"`
	Status          int            `gorm:"not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `sql:"index"`
}
