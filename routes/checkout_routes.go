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
	

	router.GET("/checkout/:id", func(c *gin.Context) {
		var checkout models.Checkout
		if err := db.Preload("User").Preload("Product").First(&checkout, "id_checkout = ?", c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Checkout not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
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
			},
		})
	})

	router.POST("/checkout", func(c *gin.Context) {
		var checkout models.Checkout
		if err := c.ShouldBindJSON(&checkout); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if checkout.IDKeranjang != nil {
			var kr models.Keranjang
			if err := db.First(&kr, *checkout.IDKeranjang).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"message": "Keranjang tidak ditemukan"})
				return
			}
			checkout.IDProduct = kr.IDProduct
		}

		if err := db.Create(&checkout).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menyimpan data checkout"})
			return
		}

		db.Preload("User").Preload("Product").First(&checkout, checkout.IDCheckout)
		c.JSON(http.StatusCreated, gin.H{
			"message": "Checkout dari keranjang berhasil",
			"data": gin.H{
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
			},
		})
	})
}
