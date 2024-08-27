package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB
var e error

// DatabaseInit ..
func DatabaseInit() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslmode := "disable"

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", user, password, dbName, sslmode)
	database, e = gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if e != nil {
		fmt.Printf("\n log:can't connect db \n")
		panic(e)
	}
}

// DB ..
func DB() *gorm.DB {
	return database
}
