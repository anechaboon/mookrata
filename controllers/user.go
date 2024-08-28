package controllers

import (
	"mookrata/models"
	"mookrata/req"
	"mookrata/resp"
	"net/http"

	"github.com/go-playground/validator"
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

	var filter bson.M = bson.M{}

	// ตรวจสอบว่า FullName มีค่าหรือไม่
	if body.FullName != "" {
		filter = bson.M{"full_name": bson.M{operator.Eq: body.FullName}}
	}
	err := mgm.Coll(&models.User{}).SimpleFind(&users, filter)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}

	// User found, return the data
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: resp.SuccessResponse{
			Data: users,
		},
	}, " ")
}

// GetUserByID ..
func (u *UserController) GetUserByID(c echo.Context) error {
	var user models.User
	id := c.Param("id")

	coll := mgm.Coll(&user)

	// Find and decode the doc to a book model.
	_ = coll.FindByID(id, &user)

	// User found, return the data
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: resp.SuccessResponse{
			Data: user,
		},
	}, " ")
}

// CreateUser ..
func (u *UserController) CreateUser(c echo.Context) error {

	var newUser models.User

	// ทำการ bind JSON body ของ request เป็น struct
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, "Error: "+err.Error())
	}

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
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: resp.SuccessResponse{
			Data: newUser,
		},
	}, " ")
}

// UpdateUser ..
func (u *UserController) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	existingUser := &models.User{}

	var body req.POSTUser

	// ใช้ c.Bind เพื่อทำการแปลง JSON body เป็น struct
	if err := c.Bind(&body); err != nil {
		return c.JSONPretty(http.StatusBadRequest, resp.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
		}, " ")
	}

	var validate = validator.New()

	// ตรวจสอบ validation
	if err := validate.Struct(&body); err != nil {
		return c.JSONPretty(http.StatusBadRequest, resp.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Validation failed: " + err.Error(),
		}, " ")
	}

	// ค้นหา User ที่มี ID ตรงกับที่ให้มา
	err := mgm.Coll(existingUser).FindByID(id, existingUser)
	if err != nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	password := []byte(body.Password)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	body.Password = string(hashedPassword)

	// อัปเดตข้อมูลใน existingUser ด้วยข้อมูลจาก body
	existingUser.FullName = body.FullName
	existingUser.UserName = body.UserName
	existingUser.Password = body.Password
	existingUser.Telephone = body.Telephone
	existingUser.Status = body.Status
	existingUser.RoleID = body.RoleID

	// บันทึกการเปลี่ยนแปลงลงใน MongoDB
	err = mgm.Coll(existingUser).Update(existingUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error: "+err.Error())
	}

	// ส่งคืนข้อมูลที่อัปเดตสำเร็จ
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: resp.SuccessResponse{
			Data: existingUser,
		},
	}, " ")

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
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: resp.SuccessResponse{
			Data: user,
		},
	}, " ")
}
