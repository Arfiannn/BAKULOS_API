package routes

import (
	"bakulos_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Product(router *gin.Engine, db *gorm.DB) {
	
   router.GET("/product", func(c *gin.Context) {
	var product []models.Product
	db.Preload("Penjual").Find(&product)
	var result []gin.H
	for _, product := range product {
		result = append(result, gin.H{
			"id_product": product.IDProduct,
			"id_penjual": product.IDPenjual,
			"kategori":   product.Kategori,
			"size":       product.Size,
			"deskripsi":  product.Deskripsi,
			"brand":      product.Brand,
			"price":      product.Price,
			"image":      product.Image,
			"warna":      product.Warna,
			"penjual": gin.H{
				"id_penjual": product.Penjual.IDPenjual,
				"nama":       product.Penjual.Nama,
			},
		})
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
  })

  router.GET("/product/:id", func(c *gin.Context) {
	var product models.Product
	if err := db.Preload("Penjual").First(&product, "id_product = ?", c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id_product": product.IDProduct,
			"id_penjual": product.IDPenjual,
			"kategori":   product.Kategori,
			"size":       product.Size,
			"deskripsi":  product.Deskripsi,
			"brand":      product.Brand,
			"price":      product.Price,
			"image":      product.Image,
			"warna":      product.Warna,
			
			"penjual": gin.H{
				"id_penjual": product.Penjual.IDPenjual,
				"nama":       product.Penjual.Nama,
			},
		},
	})
})
 
  router.POST("/product", func(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	db.Create(&product)
	db.Create(&models.Search{
		IDProduct: product.IDProduct,
	})

	c.JSON(http.StatusCreated, gin.H{"data": product})
})

  router.PUT("/product/:id", func(c *gin.Context) {
	var product models.Product
	if err := db.First(&product, "id_product = ?", c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	db.Save(&product)
	c.JSON(http.StatusOK, gin.H{"data": product})
})

	router.DELETE("/product/:id", func(c *gin.Context) {
	db.Delete(&models.Product{}, "id_product = ?", c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
})
}