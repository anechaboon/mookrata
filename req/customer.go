package req

// GETCustomer ..
type GETCustomer struct {
	Telephone string `json:"telephone"`
}

// POSTCustomer ..
type POSTCustomer struct {
	Telephone string `bson:"telephone" json:"telephone" validate:"required"`
	Count     uint   `bson:"count" json:"count" validate:"omitempty"`
	Time      uint   `bson:"time" json:"time"  validate:"omitempty"`
	Status    bool   `bson:"status" json:"status"  validate:"required"`
}
