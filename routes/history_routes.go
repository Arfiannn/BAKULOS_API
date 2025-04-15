package routes

import (
	"bakulos_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func History(router *gin.Engine, db *gorm.DB) {
	router.GET("/history", func(c *gin.Context) {
		var history []models.History
		db.Preload("User").Preload("Checkout.User").Preload("Checkout.Product").Find(&history)

		var result []gin.H
		for _, history := range history {
			result = append(result, gin.H{
				"id_history":  history.IDHistory,
				"id_user":     history.IDUser,
				"id_checkout": history.IDCheckout,
				"user": gin.H{
					"id_user": history.User.IDUser,
					"nama":    history.User.Nama,
				},
				"checkout": gin.H{
					"id_checkout":       history.Checkout.IDCheckout,
					"id_user":           history.Checkout.IDUser,
					"id_product":        history.Checkout.IDProduct,
					"alamat":            history.Checkout.Alamat,
					"metode_pengiriman": history.Checkout.MetodePengiriman,
					"pembayaran":        history.Checkout.Pembayaran,
					"jumlah":            history.Checkout.Jumlah,

					"user": gin.H{
						"id_user": history.Checkout.User.IDUser,
						"nama":    history.Checkout.User.Nama,
					},

					"product": gin.H{
						"id_product": history.Checkout.Product.IDProduct,
						"brand":      history.Checkout.Product.Brand,
					},
				},
			})
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	})

}
