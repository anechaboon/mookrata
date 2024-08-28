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

// ProductController คือ struct สำหรับจัดการ request ที่เกี่ยวข้องกับผู้ใช้
type ProductController struct{}

// GetProducts ..
func (u *ProductController) GetProducts(c echo.Context) error {
	var products []models.Product

	var body req.GETProduct

	// ใช้ c.Bind เพื่อทำการแปลง JSON body เป็น struct
	if err := c.Bind(&body); err != nil {
		return c.JSONPretty(http.StatusBadRequest, resp.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
		}, " ")
	}

	var filter bson.M = bson.M{}

	// ตรวจสอบว่า FullName มีค่าหรือไม่
	if body.MeatTypeID != "" {
		filter = bson.M{"meat_type_id": bson.M{operator.Eq: body.MeatTypeID}}
	}
	err := mgm.Coll(&models.Product{}).SimpleFind(&products, filter)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}

	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: products,
	}, " ")
	// Product found, return the data
}

// GetProductByID ..
func (u *ProductController) GetProductByID(c echo.Context) error {
	var product models.Product
	id := c.Param("id")

	coll := mgm.Coll(&product)

	// Find and decode the doc to a book model.
	_ = coll.FindByID(id, &product)

	// Product found, return the data
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: product,
	}, " ")
}

// CreateProduct ..
func (u *ProductController) CreateProduct(c echo.Context) error {

	var newProduct models.Product

	// ทำการ bind JSON body ของ request เป็น struct
	if err := c.Bind(&newProduct); err != nil {
		return c.JSON(http.StatusBadRequest, "Error: "+err.Error())
	}

	// แทรกข้อมูลลงใน MongoDB
	err := mgm.Coll(&newProduct).Create(&newProduct)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error: "+err.Error())
	}

	// ส่งคืนข้อมูลที่แทรกสำเร็จ
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: newProduct,
	}, " ")
}

// UpdateProduct ..
func (u *ProductController) UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	existingProduct := &models.Product{}

	var body req.POSTProduct

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

	// ค้นหา Product ที่มี ID ตรงกับที่ให้มา
	err := mgm.Coll(existingProduct).FindByID(id, existingProduct)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}

	// อัปเดตข้อมูลใน existingProduct ด้วยข้อมูลจาก body
	existingProduct.Name = body.Name
	existingProduct.MeatTypeID = body.MeatTypeID
	existingProduct.Weight = body.Weight
	existingProduct.Price = body.Price
	existingProduct.Status = body.Status

	// บันทึกการเปลี่ยนแปลงลงใน MongoDB
	err = mgm.Coll(existingProduct).Update(existingProduct)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error: "+err.Error())
	}

	// ส่งคืนข้อมูลที่อัปเดตสำเร็จ
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: existingProduct,
	}, " ")

}

// DeleteProductByID ..
func (u *ProductController) DeleteProductByID(c echo.Context) error {
	var product models.Product
	id := c.Param("id")

	coll := mgm.Coll(&product)

	// Find and decode the doc to a book model.
	_ = coll.FindByID(id, &product)

	err := mgm.Coll(&product).Delete(&product)
	if err != nil {
		panic(err)
	}

	// Product found, return the data
	return c.JSONPretty(http.StatusOK, resp.SuccessResponse{
		Code: http.StatusOK,
		Data: product,
	}, " ")
}
