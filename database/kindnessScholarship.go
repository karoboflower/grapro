package database

import "time"

// KindnessScholarship 应善良助学金
type KindnessScholarship struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
