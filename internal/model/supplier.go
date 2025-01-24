package model

import "time"

type Supplier struct {
	UUID        string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name        string `gorm:"type:varchar(255);not null"`
	PhoneNumber string `gorm:"type:varchar(255);not null"`
	Address     string `gorm:"type:varchar(255);not null"`
	PICName     string `gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Supplier) TableName() string {
	return "suppliers"
}
