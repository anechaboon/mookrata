package models

import (
	"time"

	"github.com/kamva/mgm/v3"
)

type M map[string]interface{}

// Define our errors:
var internalError = M{"message": "internal error"}
var userNotFound = M{"message": "user not found"}

type User struct {
	mgm.DefaultModel `bson:",inline"`
	ID               uint      `gorm:"primaryKey"`
	FullName         string    `gorm:"size:255;not null"`
	UserName         string    `gorm:"size:255;not null"`
	Password         string    `gorm:"size:255;not null"`
	RoleID           uint      `gorm:"primaryKey"`
	Telephone        string    `gorm:"size:255;"`
	Status           bool      `gorm:"not null;default:true"`
	CreatedAt        time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt        time.Time `gorm:"not null;default:current_timestamp"`
}
