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

// OrderDetailController คือ struct สำหรับจัดการ request ที่เกี่ยวข้องกับผู้ใช้
type OrderDetailController struct{}

// GetOrderDetails ..
func (u *OrderDetailController) GetOrderDetails(c echo.Context) error {
	var orderDetails []models.OrderDetail

	var body req.GETOrderDetail

	// ใช้ c.Bind เพื่อทำการแปลง JSON body เป็น struct
	if err := c.Bind(&body); err != nil {
		return c.JSONPretty(http.StatusBadRequest, resp.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
		}, " ")
	}

	var filter bson.M = bson.M{}

	// ตรวจสอบว่า FullName มีค่าหรือไม่
	if body.TableID != "" {
		filter = bson.M{"table_id": bson.M{operator.Eq: body.TableID}}
	}
	err := mgm.Coll(&models.OrderDetail{}).SimpleFind(&orderDetails, filter)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}

	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: orderDetails,
	}, " ")
	// OrderDetail found, return the data
}

// GetOrderDetailByID ..
func (u *OrderDetailController) GetOrderDetailByID(c echo.Context) error {
	var orderDetail models.OrderDetail
	id := c.Param("id")

	coll := mgm.Coll(&orderDetail)

	// Find and decode the doc to a book model.
	err := coll.FindByID(id, &orderDetail)
	if err != nil {
		return c.JSONPretty(http.StatusOK, resp.ErrorResponse{
			Code:    http.StatusOK,
			Message: "Not Found OrderDetail",
		}, " ")
	}
	// OrderDetail found, return the data
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: orderDetail,
	}, " ")
}

// CreateOrderDetail ..
func (u *OrderDetailController) CreateOrderDetail(c echo.Context) error {

	var body req.POSTOrderDetail
	var newOrderDetail models.OrderDetail

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

	deepcopier.Copy(&newOrderDetail).From(&body)

	// แทรกข้อมูลลงใน MongoDB
	err := mgm.Coll(&newOrderDetail).Create(&newOrderDetail)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error: "+err.Error())
	}

	// ส่งคืนข้อมูลที่แทรกสำเร็จ
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: newOrderDetail,
	}, " ")
}

// UpdateOrderDetail ..
func (u *OrderDetailController) UpdateOrderDetail(c echo.Context) error {
	id := c.Param("id")
	existingOrderDetail := &models.OrderDetail{}

	var body req.POSTOrderDetail

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

	// ค้นหา OrderDetail ที่มี ID ตรงกับที่ให้มา
	err := mgm.Coll(existingOrderDetail).FindByID(id, existingOrderDetail)
	if err != nil {
		return c.JSON(http.StatusNotFound, "OrderDetail not found")
	}

	// อัปเดตข้อมูลใน existingOrderDetail ด้วยข้อมูลจาก body
	existingOrderDetail.TableID = body.TableID
	existingOrderDetail.ProductID = body.ProductID
	existingOrderDetail.Quantity = body.Quantity
	existingOrderDetail.Status = body.Status

	// บันทึกการเปลี่ยนแปลงลงใน MongoDB
	err = mgm.Coll(existingOrderDetail).Update(existingOrderDetail)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error: "+err.Error())
	}

	// ส่งคืนข้อมูลที่อัปเดตสำเร็จ
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: existingOrderDetail,
	}, " ")

}

// DeleteOrderDetailByID ..
func (u *OrderDetailController) DeleteOrderDetailByID(c echo.Context) error {
	var orderDetail models.OrderDetail
	id := c.Param("id")

	coll := mgm.Coll(&orderDetail)

	// Find and decode the doc to a book model.
	err := coll.FindByID(id, &orderDetail)
	if err != nil {
		return c.JSONPretty(http.StatusOK, resp.ErrorResponse{
			Code:    http.StatusOK,
			Message: "Not Found OrderDetail",
		}, " ")
	}

	err = mgm.Coll(&orderDetail).Delete(&orderDetail)
	if err != nil {
		panic(err)
	}

	// OrderDetail found, return the data
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: orderDetail,
	}, " ")
}
