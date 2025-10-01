package routes

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	auth_controller "github.com/inventory-service/handler/auth"
	branch_controller "github.com/inventory-service/handler/branch"
	inventory_controller "github.com/inventory-service/handler/inventory"
	inventory_stock_count_controller "github.com/inventory-service/handler/inventory_stock_count"
	invoice_controller "github.com/inventory-service/handler/invoice"
	item_controller "github.com/inventory-service/handler/item"
	product_controller "github.com/inventory-service/handler/product"
	production_controller "github.com/inventory-service/handler/production"
	purchase_controller "github.com/inventory-service/handler/purchase"

	sales_controller "github.com/inventory-service/handler/sales"
	stock_controller "github.com/inventory-service/handler/stock"
	supplier_controller "github.com/inventory-service/handler/supplier"
	upload_controller "github.com/inventory-service/handler/upload"
	routes "github.com/inventory-service/routes/middleware"

	branch_resource "github.com/inventory-service/resource/branch"
	branch_product_reosurce "github.com/inventory-service/resource/branch_product"
	inventory_stock_count_resource "github.com/inventory-service/resource/inventory_stock_count"
	invoice_resource "github.com/inventory-service/resource/invoice"
	item_resource "github.com/inventory-service/resource/item"
	"github.com/inventory-service/resource/mongodb"
	product_resource "github.com/inventory-service/resource/product"
	production_resource "github.com/inventory-service/resource/production"
	production_item_resource "github.com/inventory-service/resource/production_item"

	product_recipe_resource "github.com/inventory-service/resource/product_recipe"

	inventory_resource "github.com/inventory-service/resource/inventory"
	purchase_resource "github.com/inventory-service/resource/purchase"
	sales_resource "github.com/inventory-service/resource/sales"
	stock_transaction_resource "github.com/inventory-service/resource/stock_transaction"
	supplier_resource "github.com/inventory-service/resource/supplier"
	user_resource "github.com/inventory-service/resource/user"

	branch_domain "github.com/inventory-service/domain/branch"
	inventory_stock_count_domain "github.com/inventory-service/domain/inventory_stock_count"
	invoice_domain "github.com/inventory-service/domain/invoice"
	item_domain "github.com/inventory-service/domain/item"
	product_domain "github.com/inventory-service/domain/product"
	product_recipe_domain "github.com/inventory-service/domain/product_recipe"

	branch_product_domain "github.com/inventory-service/domain/branch_product"
	inventory_domain "github.com/inventory-service/domain/inventory"
	production_domain "github.com/inventory-service/domain/production"
	purchase_domain "github.com/inventory-service/domain/purchase"
	sales_domain "github.com/inventory-service/domain/sales"
	stock_transaction_domain "github.com/inventory-service/domain/stock_transaction"
	supplier_domain "github.com/inventory-service/domain/supplier"
	user_domain "github.com/inventory-service/domain/user"

	auth_service "github.com/inventory-service/usecase/auth"
	branch_service "github.com/inventory-service/usecase/branch"
	inventory_usecase "github.com/inventory-service/usecase/inventory"
	inventory_stock_count_service "github.com/inventory-service/usecase/inventory_stock_count"
	invoice_service "github.com/inventory-service/usecase/invoice"
	item_service "github.com/inventory-service/usecase/item"
	product_service "github.com/inventory-service/usecase/product"
	production_usecase "github.com/inventory-service/usecase/production"
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
	inventoryResource := inventory_resource.NewItemBranchResource(pgDB)
	mongodbResource, err := mongodb.InitMongoDB()

	if err != nil {
		fmt.Println("Error iniitalizing mongo db")
	}

	productResource := product_resource.NewProductResource(pgDB)
	inventoryStockCountResource := inventory_stock_count_resource.NewInventoryStockCountResource(mongodbResource, "inventory_service", "inventory_stock_counts")
	invoiceResource := invoice_resource.NewInvoiceResource(pgDB)
	stockTransactionResource := stock_transaction_resource.NewStockTransactionResource(pgDB)
	salesResource := sales_resource.NewSalesResource(pgDB)
	productCompositionResource := product_recipe_resource.NewProductRecipeResource(pgDB)
	branchProductResource := branch_product_reosurce.NewBranchProductResource(pgDB)
	productionResource := production_resource.NewProductionResource(pgDB)
	productionItemResource := production_item_resource.NewProductionItemResource(pgDB)

	// Initialize usecase
	userDomain := user_domain.NewUserDomain(userResource)
	supplierDomain := supplier_domain.NewSupplierDomain(supplierResource)
	itemDomain := item_domain.NewItemDomain(itemResource, purchaseResource, inventoryResource)
	branchDomain := branch_domain.NewBranchDomain(branchResource)
	purchaseDomain := purchase_domain.NewPurchaseDomain(purchaseResource, inventoryResource, stockTransactionResource, itemResource)
	inventoryStockCountDomain := inventory_stock_count_domain.NewInventoryStockCountDomain(inventoryStockCountResource)
	productDomain := product_domain.NewProductDomain(productResource, itemResource, productCompositionResource, inventoryResource)
	invoiceDomain := invoice_domain.NewInvoiceDomain(invoiceResource)
	stockTransactionDomain := stock_transaction_domain.NewStockTransactionDomain(stockTransactionResource)
	salesDomain := sales_domain.NewSalesDomain(salesResource, productResource, branchProductResource)
	inventoryDomain := inventory_domain.NewBranchItemDomain(inventoryResource, stockTransactionResource, itemResource, purchaseResource)
	productCompositionDomain := product_recipe_domain.NewProductCompositionDomain(productCompositionResource)
	branchProductDomain := branch_product_domain.NewBranchProductDomain(branchProductResource)
	productionDomain := production_domain.NewProductionDomain(productionResource, productionItemResource)

	// initialize service
	userService := auth_service.NewUserService(userDomain)
	supplierService := supplier_service.NewSupplierService(supplierDomain)
	itemService := item_service.NewItemService(itemDomain)
	branchService := branch_service.NewBranchService(branchDomain, userDomain)
	purchaseService := purchase_service.NewPurchaseService(purchaseDomain, supplierDomain, branchDomain, itemDomain, inventoryDomain)
	inventoryStockCountService := inventory_stock_count_service.NewInventoryStockCountService(inventoryStockCountDomain, branchDomain, itemDomain)
	productService := product_service.NewProductservice(productDomain, itemDomain, productCompositionDomain, branchProductDomain)
	invoiceService := invoice_service.NewInvoiceService(invoiceDomain)
	stockService := stock_service.NewStockService(stockTransactionDomain)
	salesUsecase := sales_service.NewSalesUsecase(salesDomain, productDomain, branchProductDomain, stockTransactionDomain, inventoryDomain)
	uploadService := upload_service.NewUploadService(salesResource, productResource, salesUsecase)
	inventoryUsecase := inventory_usecase.NewInventoryUsecase(inventoryDomain, itemDomain, stockTransactionDomain)
	productionUsecase := production_usecase.NewProductionUsecase(productionDomain, stockTransactionDomain, inventoryDomain)

	// initialize controller
	authController := auth_controller.NewAuthController(userService)
	branchController := branch_controller.NewBranchController(branchService)
	supplierController := supplier_controller.NewSupplierController(supplierService)
	itemController := item_controller.NewItemController(itemService)
	purchaseController := purchase_controller.NewPurchaseController(purchaseService)
	productController := product_controller.NewProductController(productService)
	invoiceController := invoice_controller.NewInvoiceController(invoiceService)
	salesController := sales_controller.NewSalesController(salesUsecase)
	inventoryController := inventory_controller.NewInventoryHandler(inventoryUsecase)

	inventoryStockCountController := inventory_stock_count_controller.NewInventoryStockCountController(inventoryStockCountService)
	stockController := stock_controller.NewStockController(stockService)
	uploadController := upload_controller.NewUploadController(uploadService)
	productionController := production_controller.NewProductionHandler(productionUsecase)

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
			purchaseRoutes.DELETE("/:id", purchaseController.Delete)
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
			salesRoutes.GET("/", salesController.FindAll)
			salesRoutes.GET("/grouped-by-date", salesController.FindGroupedByDate)
			salesRoutes.GET("/grouped-by-date-and-branch", salesController.FindGroupedByDateAndBranch)
		}

		uploadRoutes := apiV1.Group("/upload")
		{
			uploadRoutes.POST("/transaction", uploadController.UploadTransaction)
		}

		branchItemRoutes := apiV1.Group("/inventory")
		{
			branchItemRoutes.GET("/", inventoryController.FindAll)
			branchItemRoutes.POST("/sync", inventoryController.SyncBalance)
			branchItemRoutes.POST("/", inventoryController.Create)
		}

		productionRoutes := apiV1.Group("/production")
		{
			productionRoutes.POST("/create", productionController.Create)
			productionRoutes.POST("/", productionController.GetProductionList)
		}
	}

	return router
}
