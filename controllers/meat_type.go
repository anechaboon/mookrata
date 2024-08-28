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
)

// MeatTypeController คือ struct สำหรับจัดการ request ที่เกี่ยวข้องกับผู้ใช้
type MeatTypeController struct{}

// GetMeatTypes ..
func (u *MeatTypeController) GetMeatTypes(c echo.Context) error {
	var meatTypes []models.MeatType

	var body req.GETMeatType

	// ใช้ c.Bind เพื่อทำการแปลง JSON body เป็น struct
	if err := c.Bind(&body); err != nil {
		return c.JSONPretty(http.StatusBadRequest, resp.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
		}, " ")
	}

	err := mgm.Coll(&models.MeatType{}).SimpleFind(&meatTypes, bson.M{"meat_type_id": bson.M{operator.Eq: body.MeatTypeID}})

	if err != nil {
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}

	// MeatType found, return the data
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: meatTypes,
	}, " ")
}

// GetMeatTypeByID ..
func (u *MeatTypeController) GetMeatTypeByID(c echo.Context) error {
	var meatType models.MeatType
	id := c.Param("id")

	coll := mgm.Coll(&meatType)

	// Find and decode the doc to a book model.
	_ = coll.FindByID(id, &meatType)

	// MeatType found, return the data
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: meatType,
	}, " ")
}

// CreateMeatType ..
func (u *MeatTypeController) CreateMeatType(c echo.Context) error {

	var newMeatType models.MeatType

	// ทำการ bind JSON body ของ request เป็น struct
	if err := c.Bind(&newMeatType); err != nil {
		return c.JSON(http.StatusBadRequest, "Error: "+err.Error())
	}

	// แทรกข้อมูลลงใน MongoDB
	err := mgm.Coll(&newMeatType).Create(&newMeatType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error: "+err.Error())
	}

	// ส่งคืนข้อมูลที่แทรกสำเร็จ
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: newMeatType,
	}, " ")
}

// UpdateMeatType ..
func (u *MeatTypeController) UpdateMeatType(c echo.Context) error {
	id := c.Param("id")
	existingMeatType := &models.MeatType{}

	var body req.POSTMeatType

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

	// ค้นหา MeatType ที่มี ID ตรงกับที่ให้มา
	err := mgm.Coll(existingMeatType).FindByID(id, existingMeatType)
	if err != nil {
		return c.JSON(http.StatusNotFound, "MeatType not found")
	}

	// อัปเดตข้อมูลใน existingMeatType ด้วยข้อมูลจาก body
	existingMeatType.Name = body.Name
	existingMeatType.Status = body.Status

	// บันทึกการเปลี่ยนแปลงลงใน MongoDB
	err = mgm.Coll(existingMeatType).Update(existingMeatType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error: "+err.Error())
	}

	// ส่งคืนข้อมูลที่อัปเดตสำเร็จ
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: existingMeatType,
	}, " ")

}

// DeleteMeatTypeByID ..
func (u *MeatTypeController) DeleteMeatTypeByID(c echo.Context) error {
	var meatType models.MeatType
	id := c.Param("id")

	coll := mgm.Coll(&meatType)

	// Find and decode the doc to a book model.
	_ = coll.FindByID(id, &meatType)

	err := mgm.Coll(&meatType).Delete(&meatType)
	if err != nil {
		panic(err)
	}

	// MeatType found, return the data
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: meatType,
	}, " ")
}
