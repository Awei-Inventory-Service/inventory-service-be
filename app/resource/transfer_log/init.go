package transferlog

import "gorm.io/gorm"

func NewTransferLogResource(db *gorm.DB) TransferLogResource {
	return &transferLogResource{db: db}
}
