package req

// GETProduct ..
type GETProduct struct {
	MeatTypeID string `json:"meat_type_id" validate:"required"`
}
