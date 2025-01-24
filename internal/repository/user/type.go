package user

import (
	"github.com/inventory-service/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(name, username, email, role, password string) error
	FindById(id string) (*model.User, error)
	FindUserByIdentifier(identifier string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}
