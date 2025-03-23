package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserRole string

// Define the possible values for the Role ENUM
const (
	RoleAdmin         UserRole = "Admin"
	RoleBusinessOwner UserRole = "Business Owner"
	RoleBranchManager UserRole = "Branch Manager"
	RolePurchasing    UserRole = "Purchasing"
	RoleFinance       UserRole = "Finance"
	RoleGuest         UserRole = "Guest"
)

type User struct {
	UUID      string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Username  string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	Email     string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	Role      UserRole  `gorm:"type:user_role;not null;default:'Guest'"`
	Password  string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) HashPassword(password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashed)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
