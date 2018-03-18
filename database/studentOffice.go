package database

// StudentOffice .
type StudentOffice struct {
	ID    string `gorm:"type:char(13);primary_key;not null;unique" form:"id" json:"id" binding:"required"`
	Phone string `gorm:"type:char(20);not null" form:"phone" json:"phone" binding:"required"`
}

// Exists .
func (so StudentOffice) Exists() bool {
	var studentoffice StudentOffice

	if DB.Where("id = ?", so.ID).First(&studentoffice); studentoffice.ID == so.ID {
		return true
	}
	return false
}
