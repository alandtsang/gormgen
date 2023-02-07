package model

import (
	"time"
)

// Contact Information.
type Contact struct {
	ID              uint64    `gorm:"column:id;primary_key;AUTO_INCREMENT"`                 // id
	Name            string    `gorm:"column:name;NOT NULL"`                                 // name
	Mobile          string    `gorm:"column:mobile;NOT NULL"`                               // mobile
	MobileConfirmed int       `gorm:"column:mobile_confirmed;default:0"`                    // mobile_confirmed
	Email           string    `gorm:"column:email"`                                         // email
	EmailConfirmed  int       `gorm:"column:email_confirmed;default:0"`                     // email_confirmed
	CreatedAt       time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL"` // created_at
	UpdatedAt       time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL"` // updated_at
}

func (m *Contact) TableName() string {
	return "contact"
}
