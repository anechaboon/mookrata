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
	PromotionController := new(controllers.PromotionController)
	usePromotionController := new(controllers.UsePromotionController)
	tableController := new(controllers.TableController)
	orderDetailController := new(controllers.OrderDetailController)

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

	PromotionRouter := e.Group("/promotion")
	{
		PromotionRouter.GET("", PromotionController.GetPromotions)
		PromotionRouter.GET("/:id", PromotionController.GetPromotionByID)
		PromotionRouter.POST("", PromotionController.CreatePromotion)
		PromotionRouter.PUT("/:id", PromotionController.UpdatePromotion)
		PromotionRouter.DELETE("/:id", PromotionController.DeletePromotionByID)
	}

	usePromotionRouter := e.Group("/use-promotion")
	{
		usePromotionRouter.GET("", usePromotionController.GetUsePromotions)
		usePromotionRouter.GET("/:id", usePromotionController.GetUsePromotionByID)
		usePromotionRouter.POST("", usePromotionController.CreateUsePromotion)
		usePromotionRouter.PUT("/:id", usePromotionController.UpdateUsePromotion)
		usePromotionRouter.DELETE("/:id", usePromotionController.DeleteUsePromotionByID)
	}

	tableRouter := e.Group("/table")
	{
		tableRouter.GET("", tableController.GetTables)
		tableRouter.GET("/:id", tableController.GetTableByID)
		tableRouter.POST("", tableController.CreateTable)
		tableRouter.PUT("/:id", tableController.UpdateTable)
		tableRouter.DELETE("/:id", tableController.DeleteTableByID)
	}

	orderdDetailRouter := e.Group("/order-detail")
	{
		orderdDetailRouter.GET("", orderDetailController.GetOrderDetails)
		orderdDetailRouter.GET("/:id", orderDetailController.GetOrderDetailByID)
		orderdDetailRouter.POST("", orderDetailController.CreateOrderDetail)
		orderdDetailRouter.PUT("/:id", orderDetailController.UpdateOrderDetail)
		orderdDetailRouter.DELETE("/:id", orderDetailController.DeleteOrderDetailByID)
	}

}
