package req

// GETOrderDetail ..
type GETOrderDetail struct {
	TableID string `bson:"table_id" json:"table_id"  validate:"required"`
}

// POSTOrderDetail ..
type POSTOrderDetail struct {
	TableID   string `bson:"table_id" json:"table_id"  validate:"required"`
	ProductID string `bson:"product_id" json:"product_id"  validate:"required"`
	Quantity  uint32 `bson:"quantity" json:"quantity"  validate:"required"`
	Status    bool   `bson:"status" json:"status"  validate:"required"`
}
