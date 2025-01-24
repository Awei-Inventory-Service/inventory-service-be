package auth

import "github.com/inventory-service/internal/model"

type AdjustmentLogRepository interface {
	Create(adjustment model.AdjustmentLog) error
	FindAll() ([]model.AdjustmentLog, error)
	FindByID(id string) (*model.AdjustmentLog, error)
	Delete(id string) error
}

type userRepository interface {
	Create(name, username, email, role, password string) error
	FindById(id string) (*model.User, error)
	FindUserByIdentifier(identifier string) (*model.User, error)
}
