package routers

import (
	"mookrata/controllers" // import controller

	"github.com/labstack/echo/v4"
)

// InitRoutes ..
func InitRoutes(e *echo.Echo) {
	userController := new(controllers.UserController)
	e.GET("/users", userController.GetUsers)
	e.GET("/user/:id", userController.GetUserByID)
	e.POST("/user", userController.CreateUser)
	e.DELETE("/user/:id", userController.DeleteUserByID)
}
