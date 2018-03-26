package database

import "time"

// StateGrants 国家助学金
type StateGrants struct {
	ID              int    `gorm:"primary_key;AUTO_INCREMENT"`
	StudentID       string `gorm:"type:char(13);not null;unique_index"`
	Description     string `gorm:"type:TEXT"`
	FSQuestionnaire string `gorm:"type:varchar(100);not null"`
	Accessory       string `gorm:"type:varchar(100)"`
	Status          int    `gorm:"not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `sql:"index"`
}
