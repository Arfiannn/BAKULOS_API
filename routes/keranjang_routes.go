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

	router.GET("/keranjang/:id", func(c *gin.Context) {
		var keranjang models.Keranjang
		if err := db.Preload("Product").Preload("User").First(&keranjang, "id_keranjang = ?", c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Keranjang not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"id_keranjang": keranjang.IDKeranjang,
				"id_product":   keranjang.IDProduct,
				"id_user":      keranjang.IDUser,
				"jumlah":       keranjang.Jumlah,
				"product":      keranjang.Product,

				"user": gin.H{
					"id_user": keranjang.User.IDUser,
					"nama":    keranjang.User.Nama,
				},
			},
		})
	})

	router.POST("/keranjang", func(c *gin.Context) {
		var keranjang models.Keranjang
		if err := c.ShouldBindJSON(&keranjang); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		db.Create(&keranjang)
		db.Preload("Product").Preload("User").First(&keranjang, keranjang.IDKeranjang)

		c.JSON(http.StatusCreated, gin.H{
			"data": gin.H{
				"id_keranjang": keranjang.IDKeranjang,
				"id_product":   keranjang.IDProduct,
				"id_user":      keranjang.IDUser,
				"jumlah":       keranjang.Jumlah,
				"product":      keranjang.Product,

				"user": gin.H{
					"id_user": keranjang.User.IDUser,
					"nama":    keranjang.User.Nama,
				},
			},
		})
	})

	router.PUT("/keranjang/:id", func(c *gin.Context) {
		var keranjang models.Keranjang
		if err := db.First(&keranjang, "id_keranjang = ?", c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Keranjang not found"})
			return
		}
		if err := c.ShouldBindJSON(&keranjang); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		db.Save(&keranjang)
	
		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"id_keranjang": keranjang.IDKeranjang,
				"id_product":   keranjang.IDProduct,
				"id_user":      keranjang.IDUser,
				"jumlah":       keranjang.Jumlah,
				"product":      keranjang.Product,
				"user": gin.H{
					"id_user": keranjang.User.IDUser,
					"nama":    keranjang.User.Nama,
				},
			},
		})
	})
	router.DELETE("/keranjang/:id", func(c *gin.Context) {
		db.Delete(&models.Keranjang{}, "id_keranjang = ?", c.Param("id"))
		c.JSON(http.StatusOK, gin.H{"message": "Keranjang deleted"})
	})
}




