package models

import (
	"github.com/kamva/mgm/v3"
)

// MeatType ..
type MeatType struct {
	Name             string `bson:"name" json:"name"`
	Status           bool   `bson:"status" json:"status"`
	mgm.DefaultModel `bson:",inline"`
}
