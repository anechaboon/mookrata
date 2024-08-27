package controllers

import (
	"mookrata/models"
	"mookrata/req"
	"mookrata/resp"
	"net/http"

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

	err := mgm.Coll(&models.Product{}).SimpleFind(&products, bson.M{"meat_type_id": bson.M{operator.Eq: body.MeatTypeID}})

	if err != nil {
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}

	// Product found, return the data
	return c.JSON(http.StatusOK, products)
}

// GetProductByID ..
func (u *ProductController) GetProductByID(c echo.Context) error {
	var product models.Product
	id := c.Param("id")

	coll := mgm.Coll(&product)

	// Find and decode the doc to a book model.
	_ = coll.FindByID(id, &product)

	// Product found, return the data
	return c.JSON(http.StatusOK, product)
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
	return c.JSON(http.StatusOK, newProduct)
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
	return c.JSON(http.StatusOK, product)
}
