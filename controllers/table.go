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
	"github.com/ulule/deepcopier"
	"go.mongodb.org/mongo-driver/bson"
)

// TableController คือ struct สำหรับจัดการ request ที่เกี่ยวข้องกับผู้ใช้
type TableController struct{}

// GetTables ..
func (u *TableController) GetTables(c echo.Context) error {
	var tables []models.Table

	var body req.GETTable

	// ใช้ c.Bind เพื่อทำการแปลง JSON body เป็น struct
	if err := c.Bind(&body); err != nil {
		return c.JSONPretty(http.StatusBadRequest, resp.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
		}, " ")
	}

	var filter bson.M = bson.M{}

	// ตรวจสอบว่า FullName มีค่าหรือไม่
	if body.CustomerID != "" {
		filter = bson.M{"customer_id": bson.M{operator.Eq: body.CustomerID}}
	}
	if body.Bill != "" {
		filter = bson.M{"bill": bson.M{operator.Eq: body.Bill}}
	}
	err := mgm.Coll(&models.Table{}).SimpleFind(&tables, filter)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}

	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: tables,
	}, " ")
	// Table found, return the data
}

// GetTableByID ..
func (u *TableController) GetTableByID(c echo.Context) error {
	var table models.Table
	id := c.Param("id")
	coll := mgm.Coll(&table)

	err := coll.FindByID(id, &table)
	if err != nil {
		return c.JSONPretty(http.StatusOK, resp.ErrorResponse{
			Code:    http.StatusOK,
			Message: "Not Found Table",
		}, " ")
	}

	// Table found, return the data
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: table,
	}, " ")
}

// CreateTable ..
func (u *TableController) CreateTable(c echo.Context) error {

	var body req.POSTTable
	var newTable models.Table

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

	deepcopier.Copy(&newTable).From(&body)

	// แทรกข้อมูลลงใน MongoDB
	err := mgm.Coll(&newTable).Create(&newTable)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error: "+err.Error())
	}

	// ส่งคืนข้อมูลที่แทรกสำเร็จ
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: newTable,
	}, " ")
}

// UpdateTable ..
func (u *TableController) UpdateTable(c echo.Context) error {
	id := c.Param("id")
	existingTable := &models.Table{}

	var body req.POSTTable

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

	// ค้นหา Table ที่มี ID ตรงกับที่ให้มา
	err := mgm.Coll(existingTable).FindByID(id, existingTable)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Table not found")
	}

	// อัปเดตข้อมูลใน existingTable ด้วยข้อมูลจาก body
	existingTable.CustomerID = body.CustomerID
	existingTable.Bill = body.Bill
	existingTable.Status = body.Status

	// บันทึกการเปลี่ยนแปลงลงใน MongoDB
	err = mgm.Coll(existingTable).Update(existingTable)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error: "+err.Error())
	}

	// ส่งคืนข้อมูลที่อัปเดตสำเร็จ
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: existingTable,
	}, " ")

}

// DeleteTableByID ..
func (u *TableController) DeleteTableByID(c echo.Context) error {
	var table models.Table
	id := c.Param("id")

	coll := mgm.Coll(&table)

	// Find and decode the doc to a book model.
	err := coll.FindByID(id, &table)
	if err != nil {
		return c.JSONPretty(http.StatusOK, resp.ErrorResponse{
			Code:    http.StatusOK,
			Message: "Not Found Table",
		}, " ")
	}

	err = mgm.Coll(&table).Delete(&table)
	if err != nil {
		panic(err)
	}

	// Table found, return the data
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: table,
	}, " ")
}
