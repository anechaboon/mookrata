package models

import (
	"github.com/kamva/mgm/v3"
)

// Product ..
type Product struct {
	Name             string  `bson:"name" json:"name"`
	MeatTypeID       string  `bson:"meat_type_id" json:"meat_type_id"`
	Weight           float64 `bson:"weight" json:"weight"`
	Price            float64 `bson:"price" json:"price"`
	Status           bool    `bson:"status" json:"status"`
	mgm.DefaultModel `bson:",inline"`
}
