package models

import (
	"github.com/kamva/mgm/v3"
)

type User struct {
	FullName         string `bson:"full_name" json:"full_name"`
	UserName         string `bson:"user_name" json:"user_name"`
	Password         string `bson:"password" json:"password"`
	Telephone        string `bson:"telephone" json:"telephone"`
	RoleID           uint   `bson:"role_id" json:"role_id"`
	Status           bool   `bson:"status" json:"status"`
	mgm.DefaultModel `bson:",inline"`
}
