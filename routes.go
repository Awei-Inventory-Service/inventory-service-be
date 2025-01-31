package main

import (
	"github.com/gin-gonic/gin"
	auth_controller "github.com/inventory-service/internal/controller/auth"
	branch_controller "github.com/inventory-service/internal/controller/branch"
	item_controller "github.com/inventory-service/internal/controller/item"
	purchase_controller "github.com/inventory-service/internal/controller/purchase"
	supplier_controller "github.com/inventory-service/internal/controller/supplier"

	product_controller "github.com/inventory-service/internal/controller/product"
	invoice_controller "github.com/inventory-service/internal/controller/invoice"

	"github.com/inventory-service/internal/middleware"
	branch_repository "github.com/inventory-service/internal/repository/branch"
	item_repository "github.com/inventory-service/internal/repository/item"
	product_repository "github.com/inventory-service/internal/repository/product"
	purchase_repository "github.com/inventory-service/internal/repository/purchase"
	supplier_repository "github.com/inventory-service/internal/repository/supplier"
	user_repository "github.com/inventory-service/internal/repository/user"
	invoice_repository "github.com/inventory-service/internal/repository/invoice"

	auth_service "github.com/inventory-service/internal/service/auth"
	branch_service "github.com/inventory-service/internal/service/branch"
	item_service "github.com/inventory-service/internal/service/item"
	purchase_service "github.com/inventory-service/internal/service/purchase"
	supplier_service "github.com/inventory-service/internal/service/supplier"
	invoice_service "github.com/inventory-service/internal/service/invoice"

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
	productRepository := product_repository.NewProductRepository(mongoDB, "inventory_service", "products")
	invoiceRepository := invoice_repository.NewInvoiceRepository(pgDB)

	// initialize service
	userService := auth_service.NewUserService(userRepository)
	supplierService := supplier_service.NewSupplierService(supplierRepository)
	itemService := item_service.NewItemService(itemRepository)
	branchService := branch_service.NewBranchService(branchRepository, userRepository)
	purchaseService := purchase_service.NewPurchaseService(purchaseRepository, supplierRepository, branchRepository, itemRepository)
	invoiceService := invoice_service.NewInvoiceService(invoiceRepository)

	// initialize controller
	authController := auth_controller.NewAuthController(userService)
	branchController := branch_controller.NewBranchController(branchService)
	supplierController := supplier_controller.NewSupplierController(supplierService)
	itemController := item_controller.NewItemController(itemService)
	purchaseController := purchase_controller.NewPurchaseController(purchaseService)
	productController := product_controller.NewProductController(productRepository)
	invoiceController := invoice_controller.NewInvoiceController(invoiceService)

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

		productRoutes := apiV1.Group("/product")
		{
			productRoutes.POST("/", productController.Create)
			productRoutes.GET("/", productController.FindAll)
			productRoutes.GET("/:id", productController.FindByID)
			productRoutes.PUT("/:id", productController.Update)
			productRoutes.DELETE("/:id", productController.Delete)
		}

		adminRoutes := apiV1.Group("/admin")
		{
			adminRoutes.Use(middleware.AuthMiddleware())
		}

		invoiceRoutes := apiV1.Group("/invoice")
		{
			invoiceRoutes.GET("/", invoiceController.GetInvoices)
			invoiceRoutes.GET("/:id", invoiceController.GetInvoice)
			invoiceRoutes.POST("/", invoiceController.CreateInvoice)
			invoiceRoutes.PUT("/:id", invoiceController.UpdateInvoice)
			invoiceRoutes.DELETE("/:id", invoiceController.DeleteInvoice)
		}
	}

	return router
}
