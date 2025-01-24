package transferlog

import "gorm.io/gorm"

func NewTransferLogRepository(db *gorm.DB) TransferLogRepository {
	return &transferLogRepository{db: db}
}
