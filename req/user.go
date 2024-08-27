package req

// GETUser ..
type GETUser struct {
	FullName string `json:"full_name" validate:"required"`
}
