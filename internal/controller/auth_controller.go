package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/internal/dto"
	"github.com/inventory-service/internal/service/auth"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService auth.UserService
}

func NewAuthController(authService auth.UserService) AuthController {
	return &authController{
		authService: authService,
	}
}

// Todo: standardize the error
func (a *authController) Login(ctx *gin.Context) {
	var loginRequest dto.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := a.authService.Login(loginRequest.Identifier, loginRequest.Password)
	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"token": token})
}

func (a *authController) Register(ctx *gin.Context) {
	var registerRequest dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&registerRequest); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := a.authService.Register(registerRequest.Name, registerRequest.Username, registerRequest.Email, registerRequest.Password); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "success"})
}