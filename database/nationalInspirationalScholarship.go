package database

import "time"

// NIS 国家励志奖学金
type NIS struct {
	ID        int `gorm:"primary_key;AUTO_INCREMENT"`
	status    int `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
