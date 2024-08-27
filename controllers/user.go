package controllers

import (
	"fmt"
	"mookrata/models"
	"mookrata/req"
	"mookrata/resp"
	"net/http"

	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/operator"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// UserController คือ struct สำหรับจัดการ request ที่เกี่ยวข้องกับผู้ใช้
type UserController struct{}

// GetUsers ..
func (u *UserController) GetUsers(c echo.Context) error {
	var users []models.User

	var body req.GETUser

	// ใช้ c.Bind เพื่อทำการแปลง JSON body เป็น struct
	if err := c.Bind(&body); err != nil {
		return c.JSONPretty(http.StatusBadRequest, resp.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
		}, " ")
	}

	err := mgm.Coll(&models.User{}).SimpleFind(&users, bson.M{"full_name": bson.M{operator.Eq: body.FullName}})

	if err != nil {
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}

	// User found, return the data
	return c.JSON(http.StatusOK, users)
}

// GetUserByID ..
func (u *UserController) GetUserByID(c echo.Context) error {
	var user models.User
	id := c.Param("id")

	coll := mgm.Coll(&user)

	// Find and decode the doc to a book model.
	_ = coll.FindByID(id, &user)

	// User found, return the data
	return c.JSON(http.StatusOK, user)
}

// CreateUser ..
func (u *UserController) CreateUser(c echo.Context) error {

	var newUser models.User

	// ทำการ bind JSON body ของ request เป็น struct
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, "Error: "+err.Error())
	}

	fmt.Printf("\n log:newUser.Password %v \n", newUser.Password)
	password := []byte(newUser.Password)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	newUser.Password = string(hashedPassword)

	// แทรกข้อมูลลงใน MongoDB
	err = mgm.Coll(&newUser).Create(&newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error: "+err.Error())
	}

	// ส่งคืนข้อมูลที่แทรกสำเร็จ
	return c.JSON(http.StatusOK, newUser)
}

// DeleteUserByID ..
func (u *UserController) DeleteUserByID(c echo.Context) error {
	var user models.User
	id := c.Param("id")

	coll := mgm.Coll(&user)

	// Find and decode the doc to a book model.
	_ = coll.FindByID(id, &user)

	err := mgm.Coll(&user).Delete(&user)
	if err != nil {
		panic(err)
	}

	// User found, return the data
	return c.JSON(http.StatusOK, user)
}
