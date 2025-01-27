package auth

import "github.com/inventory-service/internal/service/auth"

func NewAuthController(authService auth.UserService) AuthController {
	return &authController{
		authService: authService,
	}
}
