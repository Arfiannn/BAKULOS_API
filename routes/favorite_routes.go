package routes

import (
	"bakulos_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Favorite(router *gin.Engine, db *gorm.DB) {
	router.GET("/favorite", func(c *gin.Context) {
		var favorite []models.Favorite
		db.Preload("User").Preload("Product").Find(&favorite)

		var result []gin.H
		for _, favorite := range favorite {
			result = append(result, gin.H{
				"id_favorite": favorite.IDFavorite,
				"id_product":  favorite.IDProduct,
				"id_user":     favorite.IDUser,

				"user": gin.H{
					"nama":  favorite.User.Nama,
					"email": favorite.User.Email,
				},

				"product": gin.H{
					"id_product": favorite.Product.IDProduct,
					"brand":      favorite.Product.Brand,
					"kategori":   favorite.Product.Kategori,
					"size":       favorite.Product.Size,
					"price":      favorite.Product.Price,
					"deskripsi":  favorite.Product.Deskripsi,
					"image":      favorite.Product.Image,
				},
			})
		}
		c.JSON(http.StatusOK, gin.H{"data": result})
	})

	router.GET("/favorite/:id", func(c *gin.Context) {
		var favorite models.Favorite
		if err := db.Preload("User").Preload("Product").First(&favorite, "id_favorite = ?", c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Favorite not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"id_favorite": favorite.IDFavorite,
				"id_product":  favorite.IDProduct,
				"id_user":     favorite.IDUser,

				"user": gin.H{
					"nama":  favorite.User.Nama,
					"email": favorite.User.Email,

					"product": gin.H{
						"id_product": favorite.Product.IDProduct,
						"brand":      favorite.Product.Brand,
						"kategori":   favorite.Product.Kategori,
						"size":       favorite.Product.Size,
						"price":      favorite.Product.Price,
						"deskripsi":  favorite.Product.Deskripsi,
						"image":      favorite.Product.Image,
					},
				},
			},
		})
	})

}
