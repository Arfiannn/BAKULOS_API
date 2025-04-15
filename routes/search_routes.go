package routes

import (
	"bakulos_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Search(router *gin.Engine, db *gorm.DB) {
	router.GET("/search", func(c *gin.Context) {
		var search []models.Search
		db.Preload("Product").Find(&search)
		c.JSON(http.StatusOK, gin.H{"data": search})
	})

	router.GET("/search/:id", func(c *gin.Context) {
		var search models.Search
		if err := db.Preload("Product").First(&search, "id_search = ?", c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Search not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": search})
	})

}