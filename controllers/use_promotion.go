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

// UsePromotionController คือ struct สำหรับจัดการ request ที่เกี่ยวข้องกับผู้ใช้
type UsePromotionController struct{}

// GetUsePromotions ..
func (u *UsePromotionController) GetUsePromotions(c echo.Context) error {
	var promotions []models.UsePromotion

	var body req.GETUsePromotion

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
		filter = bson.M{"meat_type_id": bson.M{operator.Eq: body.CustomerID}}
	}
	err := mgm.Coll(&models.UsePromotion{}).SimpleFind(&promotions, filter)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}

	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: promotions,
	}, " ")
	// UsePromotion found, return the data
}

// GetUsePromotionByID ..
func (u *UsePromotionController) GetUsePromotionByID(c echo.Context) error {
	var promotion models.UsePromotion
	id := c.Param("id")

	coll := mgm.Coll(&promotion)

	// Find and decode the doc to a book model.
	err := coll.FindByID(id, &promotion)
	if err != nil {
		return c.JSONPretty(http.StatusOK, resp.ErrorResponse{
			Code:    http.StatusOK,
			Message: "Not Found UsePromotion",
		}, " ")
	}
	// UsePromotion found, return the data
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: promotion,
	}, " ")
}

// CreateUsePromotion ..
func (u *UsePromotionController) CreateUsePromotion(c echo.Context) error {

	var body req.POSTUsePromotion
	var newUsePromotion models.UsePromotion

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

	deepcopier.Copy(&newUsePromotion).From(&body)

	// แทรกข้อมูลลงใน MongoDB
	err := mgm.Coll(&newUsePromotion).Create(&newUsePromotion)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error: "+err.Error())
	}

	// ส่งคืนข้อมูลที่แทรกสำเร็จ
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: newUsePromotion,
	}, " ")
}

// UpdateUsePromotion ..
func (u *UsePromotionController) UpdateUsePromotion(c echo.Context) error {
	id := c.Param("id")
	existingUsePromotion := &models.UsePromotion{}

	var body req.POSTUsePromotion

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

	// ค้นหา UsePromotion ที่มี ID ตรงกับที่ให้มา
	err := mgm.Coll(existingUsePromotion).FindByID(id, existingUsePromotion)
	if err != nil {
		return c.JSON(http.StatusNotFound, "UsePromotion not found")
	}

	// อัปเดตข้อมูลใน existingUsePromotion ด้วยข้อมูลจาก body
	existingUsePromotion.CustomerID = body.CustomerID
	existingUsePromotion.PromotionID = body.PromotionID
	existingUsePromotion.Status = body.Status

	// บันทึกการเปลี่ยนแปลงลงใน MongoDB
	err = mgm.Coll(existingUsePromotion).Update(existingUsePromotion)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error: "+err.Error())
	}

	// ส่งคืนข้อมูลที่อัปเดตสำเร็จ
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: existingUsePromotion,
	}, " ")

}

// DeleteUsePromotionByID ..
func (u *UsePromotionController) DeleteUsePromotionByID(c echo.Context) error {
	var promotion models.UsePromotion
	id := c.Param("id")

	coll := mgm.Coll(&promotion)

	// Find and decode the doc to a book model.
	err := coll.FindByID(id, &promotion)
	if err != nil {
		return c.JSONPretty(http.StatusOK, resp.ErrorResponse{
			Code:    http.StatusOK,
			Message: "Not Found UsePromotion",
		}, " ")
	}

	err = mgm.Coll(&promotion).Delete(&promotion)
	if err != nil {
		panic(err)
	}

	// UsePromotion found, return the data
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: promotion,
	}, " ")
}
