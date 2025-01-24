package main

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/internal/controller"
	"github.com/inventory-service/internal/middleware"
	branch_repository "github.com/inventory-service/internal/repository/branch"
	item_repository "github.com/inventory-service/internal/repository/item"
	purchase_repository "github.com/inventory-service/internal/repository/purchase"
	supplier_repository "github.com/inventory-service/internal/repository/supplier"
	user_repository "github.com/inventory-service/internal/repository/user"
	auth_service "github.com/inventory-service/internal/service/auth"
	branch_service "github.com/inventory-service/internal/service/branch"
	item_service "github.com/inventory-service/internal/service/item"
	purchase_service "github.com/inventory-service/internal/service/purchase"
	supplier_service "github.com/inventory-service/internal/service/supplier"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func InitRoutes(pgDB *gorm.DB, mongoDB *mongo.Client) *gin.Engine {
	router := gin.Default()

	// initialize repository
	userRepository := user_repository.NewUserRepository(pgDB)
	supplierRepository := supplier_repository.NewSupplierRepository(pgDB)
	itemRepository := item_repository.NewItemRepository(pgDB)
	branchRepository := branch_repository.NewBranchRepository(pgDB)
	purchaseRepository := purchase_repository.NewPurchaseRepository(pgDB)

	// initialize service
	userService := auth_service.NewUserService(userRepository)
	supplierService := supplier_service.NewSupplierService(supplierRepository)
	itemService := item_service.NewItemService(itemRepository)
	branchService := branch_service.NewBranchService(branchRepository, userRepository)
	purchaseService := purchase_service.NewPurchaseService(purchaseRepository, supplierRepository, branchRepository, itemRepository)

	// initialize controller
	authController := controller.NewAuthController(userService)
	branchController := controller.NewBranchController(branchService)
	supplierController := controller.NewSupplierController(supplierService)
	itemController := controller.NewItemController(itemService)
	purchaseController := controller.NewPurchaseController(purchaseService)

	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Healthcheck success",
		})
	})

	apiV1 := router.Group("/api/v1")
	{
		authRoutes := apiV1.Group("/auth")
		{
			authRoutes.POST("/login", authController.Login)
			authRoutes.POST("/register", authController.Register)
		}

		branchRoutes := apiV1.Group("/branch")
		{
			branchRoutes.GET("/", branchController.GetBranches)
			branchRoutes.GET("/:id", branchController.GetBranch)
			branchRoutes.POST("/", branchController.CreateBranch)
			branchRoutes.PUT("/:id", branchController.UpdateBranch)
			branchRoutes.DELETE("/:id", branchController.DeleteBranch)
		}

		supplierRoutes := apiV1.Group("/supplier")
		{
			supplierRoutes.GET("/", supplierController.GetSuppliers)
			supplierRoutes.GET("/:id", supplierController.GetSupplier)
			supplierRoutes.POST("/", supplierController.CreateSupplier)
			supplierRoutes.PUT("/:id", supplierController.UpdateSupplier)
			supplierRoutes.DELETE("/:id", supplierController.DeleteSupplier)
		}

		itemRoutes := apiV1.Group("/item")
		{
			itemRoutes.GET("/", itemController.GetItems)
			itemRoutes.GET("/:id", itemController.GetItem)
			itemRoutes.POST("/", itemController.CreateItem)
			itemRoutes.PUT("/:id", itemController.UpdateItem)
			itemRoutes.DELETE("/:id", itemController.DeleteItem)
		}

		purchaseRoutes := apiV1.Group("/purchase")
		{
			purchaseRoutes.GET("/", purchaseController.GetPurchases)
			purchaseRoutes.GET("/:id", purchaseController.GetPurchase)
			purchaseRoutes.POST("/", purchaseController.CreatePurchase)
			purchaseRoutes.PUT("/:id", purchaseController.UpdatePurchase)
			purchaseRoutes.DELETE("/:id", purchaseController.DeletePurchase)
		}

		adminRoutes := apiV1.Group("/admin")
		{
			adminRoutes.Use(middleware.AuthMiddleware())
		}
	}

	return router
}
