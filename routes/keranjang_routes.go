package routes

import (
	"bakulos_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Keranjang(router *gin.Engine, db *gorm.DB) {
	router.GET("/keranjang", func(c *gin.Context) {
		var keranjang []models.Keranjang
		db.Preload("Product").Preload("User").Find(&keranjang)

		var result []gin.H
		for _, keranjang := range keranjang {
			result = append(result, gin.H{
				"id_keranjang": keranjang.IDKeranjang,
				"id_product":   keranjang.IDProduct,
				"id_user":      keranjang.IDUser,
				"jumlah":       keranjang.Jumlah,
				"product":      keranjang.Product,
				
				"user": gin.H{
					"id_user": keranjang.User.IDUser,
					"nama":    keranjang.User.Nama,
				},
			})
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	})
}




