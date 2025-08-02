package auth

import "github.com/inventory-service/usecase/auth"

func NewAuthController(authService auth.UserService) AuthController {
	return &authController{
		authService: authService,
	}
}
