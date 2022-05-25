package model

import (
	"time"
)

type Contacts struct {
	ID              int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name            string    `gorm:"column:name"`
	Mobile          string    `gorm:"column:mobile"`
	MobileConfirmed int       `gorm:"column:mobile_confirmed"`
	Email           string    `gorm:"column:email"`
	EmailConfirmed  int       `gorm:"column:email_confirmed"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}

func (Contacts) TableName() string {
	return "contacts"
}
