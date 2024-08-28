package req

// GETTable ..
type GETTable struct {
	CustomerID string `json:"customer_id"`
	Bill       string `json:"bill"`
}

// POSTTable ..
type POSTTable struct {
	CustomerID string `bson:"customer_id" json:"customer_id" validate:"required"`
	Bill       string `bson:"bill" json:"bill" validate:"required,eq=waiting|eq=completed" `
	Status     bool   `bson:"status" json:"status" validate:"required"`
}
