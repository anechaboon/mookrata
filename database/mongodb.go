package database

import (
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitMGM ..
func InitMGM() {
	// ตั้งค่า mgm เพื่อเชื่อมต่อกับ MongoDB
	err := mgm.SetDefaultConfig(nil, "mookrata", options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
}
