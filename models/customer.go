package models

import (
	"github.com/kamva/mgm/v3"
)

// Customer ..
type Customer struct {
	Telephone        string `bson:"telephone" json:"telephone"`
	Count            uint   `bson:"count" json:"count"`
	Time             uint   `bson:"time" json:"time"`
	Status           bool   `bson:"status" json:"status"`
	mgm.DefaultModel `bson:",inline"`
}
