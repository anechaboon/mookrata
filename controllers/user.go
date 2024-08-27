package controllers

import (
	"mookrata/database"
	"mookrata/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UserController คือ struct สำหรับจัดการ request ที่เกี่ยวข้องกับผู้ใช้
type UserController struct{}

// GetUsers ..
func (u *UserController) GetUsers(c echo.Context) error {
	var users []models.User

	db := database.DB()
	// Query เพื่อดึงข้อมูลผู้ใช้จากฐานข้อมูล
	result := db.Find(&users)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return c.String(http.StatusNotFound, "User not found")
		}
		return c.String(http.StatusInternalServerError, "Error: "+result.Error.Error())
	}

	// User found, return the data
	return c.JSON(http.StatusOK, users)
}

// GetUserByID ..
func (u *UserController) GetUserByID(c echo.Context) error {
	var user models.User
	id := c.Param("id")

	db := database.DB()
	// Query เพื่อดึงข้อมูลผู้ใช้จากฐานข้อมูล
	result := db.First(&user, id)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return c.String(http.StatusNotFound, "User not found")
		}
		return c.String(http.StatusInternalServerError, "Error: "+result.Error.Error())
	}

	// User found, return the data
	return c.JSON(http.StatusOK, user)
}

// CreateUser ..
func (u *UserController) CreateUser(c echo.Context) error {

	var newUser models.User

	if err := c.Bind(&newUser); err != nil {
		return c.String(http.StatusBadRequest, "Error: "+err.Error())
	}

	db := database.DB()

	if err := db.Create(&newUser).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": newUser,
	}

	return c.JSON(http.StatusOK, response)
}
