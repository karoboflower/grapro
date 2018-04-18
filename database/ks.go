package database

import (
	"time"

	"github.com/lib/pq"
)

// KS 应善良助学金
type KS struct {
	KSID              uint64         `gorm:"primary_key;AUTO_INCREMENT"`
	StudentID         string         `gorm:"type:char(13);not null"`
	Filename          string         `gorm:"type:varchar(100);not null"`
	Accessory         pq.StringArray `gorm:"type:varchar(255);not null"`
	Status            string         `gorm:"type:char(4);not null"`
	CounselorDesc     string         `gorm:"type:TEXT"`
	CounselorFile     pq.StringArray `gorm:"type:varchar(255)"`
	StudentOfficeDesc string         `gorm:"type:TEXT"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time `sql:"index"`
}
