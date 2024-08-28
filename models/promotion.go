package models

import (
	"github.com/kamva/mgm/v3"
)

// Promotion ..
type Promotion struct {
	Title            string  `bson:"title" json:"title"`
	Count            uint32  `bson:"count" json:"count"`
	DiscountAmount   float64 `bson:"discount_amount" json:"discount_amount"`
	DiscountPercent  float64 `bson:"discount_percent" json:"discount_percent"`
	Status           bool    `bson:"status" json:"status"`
	mgm.DefaultModel `bson:",inline"`
}
