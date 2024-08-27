package req

// GETMeatType ..
type GETMeatType struct {
	MeatTypeID uint `json:"name" validate:"required"`
}
