package req

// GETPromotion ..
type GETPromotion struct {
	Title string `bson:"title" json:"title" validate:"required"`
}

// POSTPromotion ..
type POSTPromotion struct {
	Title           string  `bson:"title" json:"title" validate:"required"`
	Count           uint32  `bson:"count" json:"count" validate:"required"`
	DiscountAmount  float64 `bson:"discount_amount" json:"discount_amount" validate:"omitempty"`
	DiscountPercent float64 `bson:"discount_percent" json:"discount_percent" validate:"omitempty"`
	Status          bool    `bson:"status" json:"status" validate:"required"`
}
