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

	db.AutoMigrate(&models.User{}, &models.Penjual{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Keranjang{})
	db.AutoMigrate(&models.Search{})
	db.AutoMigrate(&models.Checkout{}, &models.History{})
	db.AutoMigrate(&models.Favorite{})

	router := gin.Default()

	routes.User(router, db)
	routes.Penjual(router, db)
	routes.Product(router, db)
	routes.Keranjang(router, db)
	routes.Search(router, db)
	routes.Checkout(router, db)
	routes.History(router, db)
	routes.Favorite(router, db)

	router.Run(":3000")
}
