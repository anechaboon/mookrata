package req

// GETProduct ..
type GETProduct struct {
	MeatTypeID string `json:"meat_type_id" validate:"required"`
}

// POSTProduct ..
type POSTProduct struct {
	Name       string  `bson:"name" json:"name" validate:"required"`
	MeatTypeID string  `bson:"meat_type_id" json:"meat_type_id" validate:"required"`
	Weight     float64 `bson:"weight" json:"weight" validate:"required"`
	Price      float64 `bson:"price" json:"price" validate:"required"`
	Status     bool    `bson:"status" json:"status" validate:"required"`
}
