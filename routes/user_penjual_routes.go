package routes

import (
	"bakulos_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Penjual(router *gin.Engine, db *gorm.DB) {
	router.GET("/penjual", func(c *gin.Context) {
		var penjual []models.Penjual
		db.Find(&penjual)

		var result []gin.H
		for _, penjual := range penjual {
			result = append(result, gin.H{
				"id_penjual": penjual.IDPenjual,
				"nama":       penjual.Nama,
				"email":      penjual.Email,
			})
		}
		c.JSON(http.StatusOK, gin.H{"data": result})
	})

	router.GET("/penjual/:id", func(c *gin.Context) {
		var penjual models.Penjual
		if err := db.First(&penjual, "id_penjual = ?", c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Penjual not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"id_penjual": penjual.IDPenjual,
				"nama":       penjual.Nama,
				"email":      penjual.Email,
			},
		})
	})

	router.POST("/penjual", func(c *gin.Context) {
		var penjual models.Penjual
		if err := c.ShouldBindJSON(&penjual); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		db.Create(&penjual)
		c.JSON(http.StatusCreated, gin.H{
			"data": gin.H{
				"id_penjual": penjual.IDPenjual,
				"nama":       penjual.Nama,
				"email":      penjual.Email,
			},
		})
	})

	router.PUT("/penjual/:id", func(c *gin.Context) {
		var penjual models.Penjual
		if err := db.First(&penjual, "id_penjual = ?", c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Penjual not found"})
			return
		}

		var input models.Penjual
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		db.Model(&penjual).Updates(models.Penjual{
			Nama:     input.Nama,
			Password: input.Password,
		})

		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"id_penjual": penjual.IDPenjual,
				"nama":       penjual.Nama,
				"email":      penjual.Email,
			},
		})
	})

	router.DELETE("/penjual/:id", func(c *gin.Context) {
		db.Delete(&models.Penjual{}, "id_penjual = ?", c.Param("id"))
		c.JSON(http.StatusOK, gin.H{"message": "Penjual deleted"})
	})
}
