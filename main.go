package main

import (
	"mookrata/database"
	"mookrata/routers"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	database.InitMGM()
	e := echo.New()

	// เรียกใช้ฟังก์ชัน InitRoutes เพื่อกำหนด Routing
	routers.InitRoutes(e)

	// เริ่มเซิร์ฟเวอร์
	e.Start(":8080")
}

var secretKey = []byte("secret-key")

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
