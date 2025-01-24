package user

import "gorm.io/gorm"

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
