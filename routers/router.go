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

}
