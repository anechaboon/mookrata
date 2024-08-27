package req

// GETUser ..
type GETUser struct {
	FullName string `json:"full_name" validate:"required"`
}

// POSTUser ..
type POSTUser struct {
	FullName  string `bson:"full_name" json:"full_name" validate:"required"`
	UserName  string `bson:"user_name" json:"user_name" validate:"required"`
	Password  string `bson:"password" json:"password" validate:"required"`
	Telephone string `bson:"telephone" json:"telephone" validate:"required"`
	RoleID    uint   `bson:"role_id" json:"role_id" validate:"required"`
	Status    bool   `bson:"status" json:"status" validate:"required"`
}
