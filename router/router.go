package router

import (
	"final-project-4/controllers"
	"final-project-4/database"
	"final-project-4/middlewares"
	"final-project-4/repositories"
	"final-project-4/services"

	"github.com/gin-gonic/gin"
)

func MulaiApp() *gin.Engine{
	db := database.MulaiDB()

	// repository
	userRepository := repositories.NewUserRepository(db)
	transactionRepository := repositories.NewTransactionHistoryRepository(db)
	productRepository := repositories.NewProductRepository(db)
	categoriesRepository := repositories.NewCategoryRepository(db)

	// service
	userService := services.NewUserService(userRepository)
	transactionService := services.NewTransactionHistoryService(transactionRepository, productRepository, userRepository)
	categoriesService := services.NewCategoryService(categoriesRepository)
	productService := services.NewProductService(productRepository)

	// controller
	userController := controllers.NewUserController(userService)
	transactionController := controllers.NewTransactionHistoryController(transactionService, userService, productService)
	categoriesController := controllers.NewCategoryController(categoriesService, userService)
	productController := controllers.NewProductController(productService, userService)


	// routing
	router := gin.Default()
	// user
	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", userController.RegisterUser)
		userRouter.POST("/login", userController.Login)
		userRouter.POST("/topup", middlewares.AuthMiddleware(), userController.TopUpSaldo)
	}

	// transactions
	transactionRouter := router.Group("/transactions")
	{
		transactionRouter.POST("/", middlewares.AuthMiddleware(), transactionController.NewTransaction)
		transactionRouter.POST("/my-transactions", middlewares.AuthMiddleware(), transactionController.GetMyTransaction)
		transactionRouter.POST("/user-transactions", middlewares.AuthMiddleware(), transactionController.GetUserTransaction)
	}

	// categories
	categoriesRouter := router.Group("/categories")
	{
		categoriesRouter.POST("/", middlewares.AuthMiddleware(), categoriesController.CreateCategory)
		categoriesRouter.GET("/", middlewares.AuthMiddleware(), categoriesController.GetAllCategory)
		categoriesRouter.PATCH("/:id", middlewares.AuthMiddleware(), categoriesController.UpdateCategory)
		categoriesRouter.DELETE("/:id", middlewares.AuthMiddleware(), categoriesController.DeleteCategory)
	}

	// products
	productsRouter := router.Group("/products")
	{
		productsRouter.POST("/", middlewares.AuthMiddleware(), productController.CreateProduct)
		productsRouter.PUT("/:id", middlewares.AuthMiddleware(), productController.UpdateProduct)
		productsRouter.GET("/", middlewares.AuthMiddleware(), productController.GetAllProduct)
		productsRouter.DELETE("/:id", middlewares.AuthMiddleware(), productController.DeleteProduct)
	}
	return router
}