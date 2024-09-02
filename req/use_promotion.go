package req

// GETUsePromotion ..
type GETUsePromotion struct {
	CustomerID string `bson:"customer_id" json:"customer_id"`
}

// POSTUsePromotion ..
type POSTUsePromotion struct {
	CustomerID  string `bson:"customer_id" json:"customer_id"`
	PromotionID string `bson:"promotion_id" json:"promotion_id"`
	Status      bool   `bson:"status" json:"status" validate:"required"`
}
