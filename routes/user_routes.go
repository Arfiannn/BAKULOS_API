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

	
}
