package model

import "time"

type ContactGroups struct {
	ID        int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name      string    `gorm:"column:name"`
	Remarks   string    `gorm:"column:remarks"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (ContactGroups) TableName() string {
	return "ContactGroups"
}
