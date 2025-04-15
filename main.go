package main

import (
	"bakulos_api/models"
	"bakulos_api/routes"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bakulos?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Gagal terhubung ke database:", err)
		return
	}

	db.AutoMigrate(&models.User{})

	router := gin.Default()

	// Registrasi routes
	routes.User(router, db)

	router.Run(":3000")
}
