package routers

import (
	"mookrata/controllers" // import controller

	"github.com/labstack/echo/v4"
)

// InitRoutes ..
func InitRoutes(e *echo.Echo) {
	userController := new(controllers.UserController)
	productController := new(controllers.ProductController)
	meatTypeController := new(controllers.MeatTypeController)
	customerController := new(controllers.CustomerController)
	promotionController := new(controllers.PromotionController)
	userRouter := e.Group("/user")
	{
		userRouter.GET("", userController.GetUsers)
		userRouter.GET("/:id", userController.GetUserByID)
		userRouter.POST("", userController.CreateUser)
		userRouter.PUT("/:id", userController.UpdateUser)
		userRouter.DELETE("/:id", userController.DeleteUserByID)
	}

	productRouter := e.Group("/product")
	{
		productRouter.GET("", productController.GetProducts)
		productRouter.GET("/:id", productController.GetProductByID)
		productRouter.POST("", productController.CreateProduct)
		productRouter.PUT("/:id", productController.UpdateProduct)
		productRouter.DELETE("/:id", productController.DeleteProductByID)
	}

	meatTypeRouter := e.Group("/meat-type")
	{
		meatTypeRouter.GET("", meatTypeController.GetMeatTypes)
		meatTypeRouter.GET("/:id", meatTypeController.GetMeatTypeByID)
		meatTypeRouter.POST("", meatTypeController.CreateMeatType)
		meatTypeRouter.PUT("/:id", meatTypeController.UpdateMeatType)
		meatTypeRouter.DELETE("/:id", meatTypeController.DeleteMeatTypeByID)
	}

	customerRouter := e.Group("/customer")
	{
		customerRouter.GET("", customerController.GetCustomers)
		customerRouter.GET("/:id", customerController.GetCustomerByID)
		customerRouter.POST("", customerController.CreateCustomer)
		customerRouter.PUT("/:id", customerController.UpdateCustomer)
		customerRouter.DELETE("/:id", customerController.DeleteCustomerByID)
	}

	promotionRouter := e.Group("/promotion")
	{
		promotionRouter.GET("", promotionController.GetPromotions)
		promotionRouter.GET("/:id", promotionController.GetPromotionByID)
		promotionRouter.POST("", promotionController.CreatePromotion)
		promotionRouter.PUT("/:id", promotionController.UpdatePromotion)
		promotionRouter.DELETE("/:id", promotionController.DeletePromotionByID)
	}

}
