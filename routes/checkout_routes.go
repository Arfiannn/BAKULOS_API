package routes

import (
	"bakulos_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Checkout(router *gin.Engine, db *gorm.DB) {

	router.GET("/checkout", func(c *gin.Context) {
		var checkout []models.Checkout
		db.Preload("User").Preload("Product").Find(&checkout)

		var result []gin.H
		for _, checkout := range checkout {
			result = append(result, gin.H{
				"id_checkout":       checkout.IDCheckout,
				"id_user":           checkout.IDUser,
				"id_product":        checkout.IDProduct,
				"id_keranjang":      checkout.IDKeranjang,
				"alamat":            checkout.Alamat,
				"metode_pengiriman": checkout.MetodePengiriman,
				"pembayaran":        checkout.Pembayaran,
				"jumlah":            checkout.Jumlah,

				"user": gin.H{
					"id_user": checkout.User.IDUser,
					"nama":    checkout.User.Nama,
				},
				"product": checkout.Product,
			})
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	})

}
