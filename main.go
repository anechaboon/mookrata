package main

import (
	"mookrata/database"
	"mookrata/routers"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware" // เพิ่มการ import middleware สำหรับ CORS
)

func main() {
	// โหลด environment variables จากไฟล์ .env
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	// เริ่มการเชื่อมต่อกับ MongoDB (หรือฐานข้อมูลอื่น ๆ)
	database.InitMGM()

	// สร้าง instance ของ Echo
	e := echo.New()

	// เพิ่ม CORS middleware ที่นี่
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},                    // อนุญาตเฉพาะ origin นี้
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE}, // กำหนด method ที่อนุญาต
	}))

	// กำหนด Routing โดยเรียกฟังก์ชัน InitRoutes
	routers.InitRoutes(e)

	// เริ่มเซิร์ฟเวอร์บนพอร์ต 8080
	e.Start(":8080")
}

var secretKey = []byte("secret-key")

// ฟังก์ชันสำหรับสร้าง JWT token
func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token หมดอายุใน 24 ชั่วโมง
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
