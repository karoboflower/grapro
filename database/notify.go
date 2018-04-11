package database

import "time"

// Notify 通知
type Notify struct {
	NotifyID  uint64 `gorm:"primary_key;AUTO_INCREMENT"`
	Content   string `gorm:"type:TEXT;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
