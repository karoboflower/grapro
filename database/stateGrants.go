package database

import "time"

// StateGrants 国家助学金
type StateGrants struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
