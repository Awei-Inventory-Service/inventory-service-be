package routes

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	auth_controller "github.com/inventory-service/handler/auth"
	branch_controller "github.com/inventory-service/handler/branch"
	branch_item_controller "github.com/inventory-service/handler/branch_item"
	inventory_stock_count_controller "github.com/inventory-service/handler/inventory_stock_count"
	invoice_controller "github.com/inventory-service/handler/invoice"
	item_controller "github.com/inventory-service/handler/item"
	product_controller "github.com/inventory-service/handler/product"
	purchase_controller "github.com/inventory-service/handler/purchase"

	sales_controller "github.com/inventory-service/handler/sales"
	stock_controller "github.com/inventory-service/handler/stock"
	supplier_controller "github.com/inventory-service/handler/supplier"
	upload_controller "github.com/inventory-service/handler/upload"
	routes "github.com/inventory-service/routes/middleware"

	branch_resource "github.com/inventory-service/resource/branch"
	inventory_stock_count_resource "github.com/inventory-service/resource/inventory_stock_count"
	invoice_resource "github.com/inventory-service/resource/invoice"
	item_resource "github.com/inventory-service/resource/item"
	item_composition_resource "github.com/inventory-service/resource/item_composition"
	item_purchase_chain_resource "github.com/inventory-service/resource/item_purchase_chain"
	"github.com/inventory-service/resource/mongodb"
	product_resource "github.com/inventory-service/resource/product"
	product_composition_resource "github.com/inventory-service/resource/product_composition"

	branch_item_resource "github.com/inventory-service/resource/branch_item"
	purchase_resource "github.com/inventory-service/resource/purchase"
	sales_resource "github.com/inventory-service/resource/sales"
	stock_transaction_resource "github.com/inventory-service/resource/stock_transaction"
	supplier_resource "github.com/inventory-service/resource/supplier"
	user_resource "github.com/inventory-service/resource/user"

	branch_domain "github.com/inventory-service/domain/branch"
	inventory_stock_count_domain "github.com/inventory-service/domain/inventory_stock_count"
	invoice_domain "github.com/inventory-service/domain/invoice"
	item_domain "github.com/inventory-service/domain/item"
	item_composition_domain "github.com/inventory-service/domain/item_composition"
	item_purchase_chain_domain "github.com/inventory-service/domain/item_purchase_chain"
	product_domain "github.com/inventory-service/domain/product"
	product_composition_domain "github.com/inventory-service/domain/product_composition"

	branch_item_domain "github.com/inventory-service/domain/branch_item"
	purchase_domain "github.com/inventory-service/domain/purchase"
	sales_domain "github.com/inventory-service/domain/sales"
	stock_transaction_domain "github.com/inventory-service/domain/stock_transaction"
	supplier_domain "github.com/inventory-service/domain/supplier"
	user_domain "github.com/inventory-service/domain/user"

	auth_service "github.com/inventory-service/usecase/auth"
	branch_service "github.com/inventory-service/usecase/branch"
	branch_item_usecase "github.com/inventory-service/usecase/branch_item"
	inventory_stock_count_service "github.com/inventory-service/usecase/inventory_stock_count"
	invoice_service "github.com/inventory-service/usecase/invoice"
	item_service "github.com/inventory-service/usecase/item"
	item_purchase_chain_service "github.com/inventory-service/usecase/item_purchase_chain"
	product_service "github.com/inventory-service/usecase/product"
	purchase_service "github.com/inventory-service/usecase/purchase"
	sales_service "github.com/inventory-service/usecase/sales"
	stock_service "github.com/inventory-service/usecase/stock"

	supplier_service "github.com/inventory-service/usecase/supplier"
	upload_service "github.com/inventory-service/usecase/upload"
	"gorm.io/gorm"
)

func InitRoutes(pgDB *gorm.DB) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// initialize repository
	userResource := user_resource.NewUserResource(pgDB)
	supplierResource := supplier_resource.NewSupplierResource(pgDB)
	itemResource := item_resource.NewItemResource(pgDB)
	branchResource := branch_resource.NewBranchResource(pgDB)
	purchaseResource := purchase_resource.NewPurchaseResource(pgDB)
	itemBalanceResource := branch_item_resource.NewItemBranchResource(pgDB)
	mongodbResource, err := mongodb.InitMongoDB()

	if err != nil {
		fmt.Println("Error iniitalizing mongo db")
	}

	productResource := product_resource.NewProductResource(pgDB)
	inventoryStockCountResource := inventory_stock_count_resource.NewInventoryStockCountResource(mongodbResource, "inventory_service", "inventory_stock_counts")
	invoiceResource := invoice_resource.NewInvoiceResource(pgDB)
	stockTransactionResource := stock_transaction_resource.NewStockTransactionResource(pgDB)
	salesResource := sales_resource.NewSalesResource(pgDB)
	itemPurchaseChainResource := item_purchase_chain_resource.NewItemPurchaseChainResource(mongodbResource, "inventory_service", "itempurchasechain")
	productCompositionResource := product_composition_resource.NewProductCompositionResource(pgDB)
	itemCompositionResource := item_composition_resource.NewItemCompositionResource(pgDB)

	// Initialize usecase
	userDomain := user_domain.NewUserDomain(userResource)
	supplierDomain := supplier_domain.NewSupplierDomain(supplierResource)
	itemDomain := item_domain.NewItemDomain(itemResource, itemCompositionResource, purchaseResource, itemBalanceResource)
	branchDomain := branch_domain.NewBranchDomain(branchResource)
	purchaseDomain := purchase_domain.NewPurchaseDomain(purchaseResource, itemBalanceResource, stockTransactionResource)
	inventoryStockCountDomain := inventory_stock_count_domain.NewInventoryStockCountDomain(inventoryStockCountResource)
	productDomain := product_domain.NewProductDomain(productResource, itemResource, productCompositionResource)
	invoiceDomain := invoice_domain.NewInvoiceDomain(invoiceResource)
	stockDomain := stock_transaction_domain.NewStockTransactionDomain(stockTransactionResource)
	itemPurchaseChainDomain := item_purchase_chain_domain.NewItemPurchaseChainDomain(itemPurchaseChainResource)
	salesDomain := sales_domain.NewSalesDomain(salesResource)
	itemBranchDomain := branch_item_domain.NewItemBranchDomain(itemBalanceResource, stockTransactionResource, itemResource)
	productCompositionDomain := product_composition_domain.NewProductCompositionDomain(productCompositionResource)
	itemCompositionDomain := item_composition_domain.NewItemCompositionDomain(itemCompositionResource)

	// initialize service
	userService := auth_service.NewUserService(userDomain)
	supplierService := supplier_service.NewSupplierService(supplierDomain)
	itemService := item_service.NewItemService(itemDomain, itemCompositionDomain)
	branchService := branch_service.NewBranchService(branchDomain, userDomain)
	purchaseService := purchase_service.NewPurchaseService(purchaseDomain, supplierDomain, branchDomain, itemDomain, itemBranchDomain)
	inventoryStockCountService := inventory_stock_count_service.NewInventoryStockCountService(inventoryStockCountDomain, branchDomain, itemDomain)
	productService := product_service.NewProductservice(productDomain, itemDomain, productCompositionDomain)
	invoiceService := invoice_service.NewInvoiceService(invoiceDomain)
	stockService := stock_service.NewStockService(stockDomain)
	itemPurchaseChainService := item_purchase_chain_service.NewItemPurchaseChainService(itemPurchaseChainDomain, purchaseDomain, itemDomain, branchDomain)
	salesService := sales_service.NewSalesService(salesDomain, productDomain, itemPurchaseChainDomain, itemPurchaseChainService)
	uploadService := upload_service.NewUploadService(salesResource, productResource, salesService)
	itemBranchUsecase := branch_item_usecase.NewBranchItemUsecase(itemBranchDomain)

	// initialize controller
	authController := auth_controller.NewAuthController(userService)
	branchController := branch_controller.NewBranchController(branchService)
	supplierController := supplier_controller.NewSupplierController(supplierService)
	itemController := item_controller.NewItemController(itemService)
	purchaseController := purchase_controller.NewPurchaseController(purchaseService)
	productController := product_controller.NewProductController(productService)
	invoiceController := invoice_controller.NewInvoiceController(invoiceService)
	salesController := sales_controller.NewSalesController(salesService)
	stockBalanceController := branch_item_controller.NewBranchItemHandler(itemBranchUsecase)

	inventoryStockCountController := inventory_stock_count_controller.NewInventoryStockCountController(inventoryStockCountService)
	stockController := stock_controller.NewStockController(stockService)
	uploadController := upload_controller.NewUploadController(uploadService)

	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Healthcheck success",
		})
	})

	apiV1 := router.Group("/api/v1")
	{
		apiV1.Use(routes.BasicMiddleware())
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
			purchaseRoutes.POST("/", purchaseController.Create)
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
			adminRoutes.Use(routes.AuthMiddleware())
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

		salesRoutes := apiV1.Group("/sales")
		{
			salesRoutes.POST("/", salesController.Create)
		}

		uploadRoutes := apiV1.Group("/upload")
		{
			uploadRoutes.POST("/transaction", uploadController.UploadTransaction)
		}

		stockBalanceRoutes := apiV1.Group("/stock-balance")
		{
			stockBalanceRoutes.GET("/", stockBalanceController.FindAllStockBalance)
		}
	}

	return router
}
