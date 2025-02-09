package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	auth_controller "github.com/inventory-service/internal/controller/auth"
	branch_controller "github.com/inventory-service/internal/controller/branch"
	inventory_stock_count_controller "github.com/inventory-service/internal/controller/inventory_stock_count"
	invoice_controller "github.com/inventory-service/internal/controller/invoice"
	item_controller "github.com/inventory-service/internal/controller/item"
	product_controller "github.com/inventory-service/internal/controller/product"
	purchase_controller "github.com/inventory-service/internal/controller/purchase"
	stock_controller "github.com/inventory-service/internal/controller/stock"
	supplier_controller "github.com/inventory-service/internal/controller/supplier"

	"github.com/inventory-service/internal/middleware"
	branch_repository "github.com/inventory-service/internal/repository/branch"
	inventory_stock_count_repository "github.com/inventory-service/internal/repository/inventory_stock_count"
	invoice_repository "github.com/inventory-service/internal/repository/invoice"
	item_repository "github.com/inventory-service/internal/repository/item"
	"github.com/inventory-service/internal/repository/mongodb"
	product_repository "github.com/inventory-service/internal/repository/product"
	purchase_repository "github.com/inventory-service/internal/repository/purchase"
	stock_transaction_repository "github.com/inventory-service/internal/repository/stock_transaction"
	supplier_repository "github.com/inventory-service/internal/repository/supplier"
	user_repository "github.com/inventory-service/internal/repository/user"

	auth_service "github.com/inventory-service/internal/service/auth"
	branch_service "github.com/inventory-service/internal/service/branch"
	inventory_stock_count_service "github.com/inventory-service/internal/service/inventory_stock_count"
	invoice_service "github.com/inventory-service/internal/service/invoice"
	item_service "github.com/inventory-service/internal/service/item"
	product_service "github.com/inventory-service/internal/service/product"
	purchase_service "github.com/inventory-service/internal/service/purchase"
	supplier_service "github.com/inventory-service/internal/service/supplier"
	stock_service "github.com/inventory-service/internal/service/stock"

	"gorm.io/gorm"
)

func InitRoutes(pgDB *gorm.DB) *gin.Engine {
	router := gin.Default()

	// initialize repository
	userRepository := user_repository.NewUserRepository(pgDB)
	supplierRepository := supplier_repository.NewSupplierRepository(pgDB)
	itemRepository := item_repository.NewItemRepository(pgDB)
	branchRepository := branch_repository.NewBranchRepository(pgDB)
	purchaseRepository := purchase_repository.NewPurchaseRepository(pgDB)
	mongodbRepository, err := mongodb.InitMongoDB()

	if err != nil {
		fmt.Println("Error iniitalizing mongo db")
	}

	productRepository := product_repository.NewProductRepository(mongodbRepository, "inventory_service", "products")
	inventoryStockCountRepository := inventory_stock_count_repository.NewInventoryStockCountRepository(mongodbRepository, "inventory_service", "inventory_stock_counts")
	invoiceRepository := invoice_repository.NewInvoiceRepository(pgDB)
	stockTransactionRepository := stock_transaction_repository.NewStockTransactionRepository(pgDB)

	// initialize service
	userService := auth_service.NewUserService(userRepository)
	supplierService := supplier_service.NewSupplierService(supplierRepository)
	itemService := item_service.NewItemService(itemRepository)
	branchService := branch_service.NewBranchService(branchRepository, userRepository)
	purchaseService := purchase_service.NewPurchaseService(purchaseRepository, supplierRepository, branchRepository, itemRepository)
	inventoryStockCountService := inventory_stock_count_service.NewInventoryStockCountService(inventoryStockCountRepository, branchRepository, itemRepository)
	productService := product_service.NewProductservice(productRepository, itemRepository)
	invoiceService := invoice_service.NewInvoiceService(invoiceRepository)
	stockService := stock_service.NewStockService(stockTransactionRepository)

	// initialize controller
	authController := auth_controller.NewAuthController(userService)
	branchController := branch_controller.NewBranchController(branchService)
	supplierController := supplier_controller.NewSupplierController(supplierService)
	itemController := item_controller.NewItemController(itemService)
	purchaseController := purchase_controller.NewPurchaseController(purchaseService)
	productController := product_controller.NewProductController(productService)
	invoiceController := invoice_controller.NewInvoiceController(invoiceService)

	inventoryStockCountController := inventory_stock_count_controller.NewInventoryStockCountController(inventoryStockCountService)
	stockController := stock_controller.NewStockController(stockService)

	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Healthcheck success",
		})
	})

	apiV1 := router.Group("/api/v1")
	{
		apiV1.Use(middleware.BasicMiddleware())
		authRoutes := apiV1.Group("/auth")
		{
			authRoutes.POST("/login", authController.Login)
			authRoutes.POST("/register", authController.Register)
		}

		branchRoutes := apiV1.Group("/branch")
		{
			// branchRoutes.Use(middleware.AuthMiddleware())
			branchRoutes.GET("/", branchController.GetBranches)
			branchRoutes.GET("/:id", branchController.GetBranch)
			branchRoutes.POST("/", branchController.CreateBranch)
			branchRoutes.PUT("/:id", branchController.UpdateBranch)
			branchRoutes.DELETE("/:id", branchController.DeleteBranch)
		}

		supplierRoutes := apiV1.Group("/supplier")
		{
			// supplierRoutes.Use(middleware.AuthMiddleware())
			supplierRoutes.GET("/", supplierController.GetSuppliers)
			supplierRoutes.GET("/:id", supplierController.GetSupplier)
			supplierRoutes.POST("/", supplierController.CreateSupplier)
			supplierRoutes.PUT("/:id", supplierController.UpdateSupplier)
			supplierRoutes.DELETE("/:id", supplierController.DeleteSupplier)
		}

		itemRoutes := apiV1.Group("/item")
		{
			// itemRoutes.Use(middleware.AuthMiddleware())
			itemRoutes.GET("/", itemController.GetItems)
			itemRoutes.GET("/:id", itemController.GetItem)
			itemRoutes.POST("/", itemController.CreateItem)
			itemRoutes.PUT("/:id", itemController.UpdateItem)
			itemRoutes.DELETE("/:id", itemController.DeleteItem)
		}

		purchaseRoutes := apiV1.Group("/purchase")
		{
			// purchaseRoutes.Use(middleware.AuthMiddleware())
			purchaseRoutes.GET("/", purchaseController.GetPurchases)
			purchaseRoutes.GET("/:id", purchaseController.GetPurchase)
			purchaseRoutes.POST("/", purchaseController.CreatePurchase)
			purchaseRoutes.PUT("/:id", purchaseController.UpdatePurchase)
			purchaseRoutes.DELETE("/:id", purchaseController.DeletePurchase)
		}

		productRoutes := apiV1.Group("/product")
		{
			// productRoutes.Use(middleware.AuthMiddleware())
			productRoutes.POST("/", productController.Create)
			productRoutes.GET("/", productController.FindAll)
			productRoutes.GET("/:id", productController.FindByID)
			productRoutes.PUT("/:id", productController.Update)
			productRoutes.DELETE("/:id", productController.Delete)
		}

		inventoryStockCountRoutes := apiV1.Group("/inventory-stock-count")
		{
			// inventoryStockCountRoutes.Use(middleware.AuthMiddleware())
			inventoryStockCountRoutes.POST("/", inventoryStockCountController.Create)
			inventoryStockCountRoutes.GET("/", inventoryStockCountController.FindAll)
			inventoryStockCountRoutes.GET("/:id", inventoryStockCountController.FindByID)
			inventoryStockCountRoutes.PUT("/:id", inventoryStockCountController.Update)
			inventoryStockCountRoutes.GET("/branch/:id", inventoryStockCountController.FilterByBranch)
			inventoryStockCountRoutes.DELETE("/:id", inventoryStockCountController.Delete)
		}

		adminRoutes := apiV1.Group("/admin")
		{
			adminRoutes.Use(middleware.AuthMiddleware())
		}

		invoiceRoutes := apiV1.Group("/invoice")
		{
			// invoiceRoutes.Use(middleware.AuthMiddleware())
			invoiceRoutes.GET("/", invoiceController.GetInvoices)
			invoiceRoutes.GET("/:id", invoiceController.GetInvoice)
			invoiceRoutes.POST("/", invoiceController.CreateInvoice)
			invoiceRoutes.PUT("/:id", invoiceController.UpdateInvoice)
			invoiceRoutes.DELETE("/:id", invoiceController.DeleteInvoice)
		}

		stockRoutes := apiV1.Group("/stock")
		{
			stockRoutes.GET("/:id/qty", stockController.GetStockByItemID)
		}
	}

	return router
}
