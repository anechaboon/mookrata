package controllers

import (
	"fmt"
	"mookrata/models"
	"mookrata/req"
	"mookrata/resp"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/operator"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// CustomerController คือ struct สำหรับจัดการ request ที่เกี่ยวข้องกับผู้ใช้
type CustomerController struct{}

// GetCustomers ..
func (u *CustomerController) GetCustomers(c echo.Context) error {
	var customers []models.Customer

	var body req.GETCustomer

	// ใช้ c.Bind เพื่อทำการแปลง JSON body เป็น struct
	if err := c.Bind(&body); err != nil {
		return c.JSONPretty(http.StatusBadRequest, resp.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
		}, " ")
	}

	var filter bson.M = bson.M{}

	// ตรวจสอบว่า FullName มีค่าหรือไม่
	if body.Telephone != "" {
		filter = bson.M{"telephone": bson.M{operator.Eq: body.Telephone}}
	}
	err := mgm.Coll(&models.Customer{}).SimpleFind(&customers, filter)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}

	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: customers,
	}, " ")
	// Customer found, return the data
}

// GetCustomerByID ..
func (u *CustomerController) GetCustomerByID(c echo.Context) error {
	var customer models.Customer
	id := c.Param("id")

	coll := mgm.Coll(&customer)

	// Find and decode the doc to a book model.
	_ = coll.FindByID(id, &customer)

	// Customer found, return the data
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: customer,
	}, " ")
}

// CreateCustomer ..
func (u *CustomerController) CreateCustomer(c echo.Context) error {

	var newCustomer models.Customer

	// ทำการ bind JSON body ของ request เป็น struct
	if err := c.Bind(&newCustomer); err != nil {
		return c.JSON(http.StatusBadRequest, "Error: "+err.Error())
	}

	// แทรกข้อมูลลงใน MongoDB
	err := mgm.Coll(&newCustomer).Create(&newCustomer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error: "+err.Error())
	}

	// ส่งคืนข้อมูลที่แทรกสำเร็จ
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: newCustomer,
	}, " ")
}

// UpdateCustomer ..
func (u *CustomerController) UpdateCustomer(c echo.Context) error {
	id := c.Param("id")
	existingCustomer := &models.Customer{}

	var body req.POSTCustomer

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

	// ค้นหา Customer ที่มี ID ตรงกับที่ให้มา
	err := mgm.Coll(existingCustomer).FindByID(id, existingCustomer)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Customer not found")
	}

	fmt.Printf("\n log:body %v \n", body)
	// อัปเดตข้อมูลใน existingCustomer ด้วยข้อมูลจาก body
	existingCustomer.Telephone = body.Telephone
	existingCustomer.Count = body.Count
	existingCustomer.Time = body.Time
	existingCustomer.Status = body.Status

	// บันทึกการเปลี่ยนแปลงลงใน MongoDB
	err = mgm.Coll(existingCustomer).Update(existingCustomer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error: "+err.Error())
	}

	// ส่งคืนข้อมูลที่อัปเดตสำเร็จ
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: existingCustomer,
	}, " ")

}

// DeleteCustomerByID ..
func (u *CustomerController) DeleteCustomerByID(c echo.Context) error {
	var customer models.Customer
	id := c.Param("id")

	coll := mgm.Coll(&customer)

	// Find and decode the doc to a book model.
	_ = coll.FindByID(id, &customer)

	err := mgm.Coll(&customer).Delete(&customer)
	if err != nil {
		panic(err)
	}

	// Customer found, return the data
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: customer,
	}, " ")
}
