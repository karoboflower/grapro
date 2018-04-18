package database

import (
	"time"

	"github.com/lib/pq"
)

// SG 国家助学金
type SG struct {
	SGID              uint64         `gorm:"primary_key;AUTO_INCREMENT"`
	StudentID         string         `gorm:"type:char(13);not null"`
	FSQuestionnaire   string         `gorm:"type:varchar(100);not null"`
	Accessory         pq.StringArray `gorm:"type:varchar(255)"`
	Level             string         `gorm:"type:char(4);not null"`
	Status            string         `gorm:"type:char(4);not null"`
	CounselorDesc     string         `gorm:"type:TEXT"`
	StudentOfficeDesc string         `gorm:"type:TEXT"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time `sql:"index"`
}
