package database

import "time"

// Notify 通知
type Notify struct {
	ID        int    `gorm:"primary_key;AUTO_INCREMENT"`
	Content   string `gorm:"type:TEXT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
