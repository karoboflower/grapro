package database

import (
	"time"

	"github.com/lib/pq"
)

// NIS 国家励志奖学金
type NIS struct {
	ID        uint32         `gorm:"primary_key;AUTO_INCREMENT"`
	StudentID string         `gorm:"type:char(13);not null"`
	Accessory pq.StringArray `gorm:"type:varchar(255);not null"`
	Status    string         `gorm:"type:char(4);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
