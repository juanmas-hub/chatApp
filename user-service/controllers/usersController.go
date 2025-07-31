package controllers

import (
	"net/http"
	"user-service/initializers"
	"user-service/internal/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// Get data off req body
	var body struct{
		Username string
		Email 	 string
		Password string
	}

	if c.Bind(&body) != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})

		return
	}

	
	// Hash de password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10) // default cost

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash password",
		})

		return
	}

	// Create the user
	user := models.User{Username: body.Username, Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{})
}