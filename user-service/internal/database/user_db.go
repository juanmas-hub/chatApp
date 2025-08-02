package database

import (
	"net/http"
	"user-service/initializers"
	"user-service/internal/models"

	"github.com/gin-gonic/gin"
)
func CreateUser[T any](c *gin.Context, user *T) int {

	result := initializers.DB.Create(user)

	if result.Error != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user",
		})

		return 1
	}
	return 0
}

func CheckExistencebyEmail(c *gin.Context, user *models.User, email string) int {
	initializers.DB.First(user, "email = ?", email)

	if user.ID == 0{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email",
		})

		return 1
	}
	return 0
}
