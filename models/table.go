package models

import (
	"github.com/kamva/mgm/v3"
)

// Table ..
type Table struct {
	CustomerID       string `bson:"customer_id" json:"customer_id"`
	Bill             string `bson:"bill" json:"bill"` // ['waiting', 'completed']
	Status           bool   `bson:"status" json:"status"`
	mgm.DefaultModel `bson:",inline"`
}
