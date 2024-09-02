package models

import (
	"github.com/kamva/mgm/v3"
)

// UsePromotion ..
type UsePromotion struct {
	CustomerID       string `bson:"customer_id" json:"customer_id"`
	PromotionID      string `bson:"promotion_id" json:"promotion_id"`
	Status           bool   `bson:"status" json:"status"`
	mgm.DefaultModel `bson:",inline"`
}
