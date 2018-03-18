package database

import "time"

// NIS 国家励志奖学金
type NIS struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
