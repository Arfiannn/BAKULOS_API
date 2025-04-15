package routes

import (
	"bakulos_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func User(router *gin.Engine, db *gorm.DB) {

	router.GET("/user", func(c *gin.Context) {
		var user []models.User
		db.Find(&user)

		var result []gin.H
		for _, user := range user {
			result = append(result, gin.H{
				"id_user": user.IDUser,
				"nama":    user.Nama,
				"email":   user.Email,
			})
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	})

	router.GET("/user/:id", func(c *gin.Context) {
		var user models.User
		if err := db.First(&user, "id_user = ?", c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"id_user": user.IDUser,
				"nama":    user.Nama,
				"email":   user.Email,
			},
		})
	})

	router.POST("/user", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		db.Create(&user)
		c.JSON(http.StatusCreated, gin.H{
			"data": gin.H{
				"id_user": user.IDUser,
				"nama":    user.Nama,
				"email":   user.Email,
			},
		})
	})

	router.PUT("/user/:id", func(c *gin.Context) {
		var user models.User
		if err := db.First(&user, "id_user = ?", c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}

		var input models.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		db.Model(&user).Updates(models.User{
			Nama:     input.Nama,
			Password: input.Password,
		})

		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"id_user": user.IDUser,
				"nama":    user.Nama,
				"email":   user.Email,
			},
		})
	})
	
	router.DELETE("/user/:id", func(c *gin.Context) {
		db.Delete(&models.User{}, "id_user = ?", c.Param("id"))
		c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
	})
}
