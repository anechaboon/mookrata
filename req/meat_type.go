package req

// GETMeatType ..
type GETMeatType struct {
	MeatTypeID uint `json:"name" validate:"required"`
}

// POSTMeatType ..
type POSTMeatType struct {
	Name   string `bson:"name" json:"name" validate:"required"`
	Status bool   `bson:"status" json:"status" validate:"required"`
}
