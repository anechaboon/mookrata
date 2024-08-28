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

// PromotionController คือ struct สำหรับจัดการ request ที่เกี่ยวข้องกับผู้ใช้
type PromotionController struct{}

// GetPromotions ..
func (u *PromotionController) GetPromotions(c echo.Context) error {
	var promotions []models.Promotion

	var body req.GETPromotion

	// ใช้ c.Bind เพื่อทำการแปลง JSON body เป็น struct
	if err := c.Bind(&body); err != nil {
		return c.JSONPretty(http.StatusBadRequest, resp.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
		}, " ")
	}

	var filter bson.M = bson.M{}

	// ตรวจสอบว่า FullName มีค่าหรือไม่
	if body.Title != "" {
		filter = bson.M{"meat_type_id": bson.M{operator.Eq: body.Title}}
	}
	err := mgm.Coll(&models.Promotion{}).SimpleFind(&promotions, filter)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}

	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: promotions,
	}, " ")
	// Promotion found, return the data
}

// GetPromotionByID ..
func (u *PromotionController) GetPromotionByID(c echo.Context) error {
	var promotion models.Promotion
	id := c.Param("id")

	coll := mgm.Coll(&promotion)

	// Find and decode the doc to a book model.
	_ = coll.FindByID(id, &promotion)

	// Promotion found, return the data
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: promotion,
	}, " ")
}

// CreatePromotion ..
func (u *PromotionController) CreatePromotion(c echo.Context) error {

	var newPromotion models.Promotion

	// ทำการ bind JSON body ของ request เป็น struct
	if err := c.Bind(&newPromotion); err != nil {
		return c.JSON(http.StatusBadRequest, "Error: "+err.Error())
	}

	// แทรกข้อมูลลงใน MongoDB
	err := mgm.Coll(&newPromotion).Create(&newPromotion)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error: "+err.Error())
	}

	// ส่งคืนข้อมูลที่แทรกสำเร็จ
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: newPromotion,
	}, " ")
}

// UpdatePromotion ..
func (u *PromotionController) UpdatePromotion(c echo.Context) error {
	id := c.Param("id")
	existingPromotion := &models.Promotion{}

	var body req.POSTPromotion

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

	// ค้นหา Promotion ที่มี ID ตรงกับที่ให้มา
	err := mgm.Coll(existingPromotion).FindByID(id, existingPromotion)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Promotion not found")
	}

	// อัปเดตข้อมูลใน existingPromotion ด้วยข้อมูลจาก body
	existingPromotion.Title = body.Title
	existingPromotion.Count = body.Count
	existingPromotion.DiscountAmount = body.DiscountAmount
	existingPromotion.DiscountPercent = body.DiscountPercent
	existingPromotion.Status = body.Status

	// บันทึกการเปลี่ยนแปลงลงใน MongoDB
	err = mgm.Coll(existingPromotion).Update(existingPromotion)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error: "+err.Error())
	}

	// ส่งคืนข้อมูลที่อัปเดตสำเร็จ
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: existingPromotion,
	}, " ")

}

// DeletePromotionByID ..
func (u *PromotionController) DeletePromotionByID(c echo.Context) error {
	var promotion models.Promotion
	id := c.Param("id")

	coll := mgm.Coll(&promotion)

	// Find and decode the doc to a book model.
	_ = coll.FindByID(id, &promotion)

	err := mgm.Coll(&promotion).Delete(&promotion)
	if err != nil {
		panic(err)
	}

	// Promotion found, return the data
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: promotion,
	}, " ")
}
