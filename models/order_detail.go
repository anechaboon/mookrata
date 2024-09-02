package models

import (
	"github.com/kamva/mgm/v3"
)

// OrderDetail ..
type OrderDetail struct {
	TableID          string `bson:"table_id" json:"table_id"`
	ProductID        string `bson:"product_id" json:"product_id"`
	Quantity         uint32 `bson:"quantity" json:"quantity"`
	Status           bool   `bson:"status" json:"status"`
	mgm.DefaultModel `bson:",inline"`
}
