package controllers

import (
	"net/http"
	"user-service/internal/models"
	"user-service/internal/database"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var body struct{
		Username string
		Email 	 string
		Password string
	}

	if Bind(c, &body) != 0 {
		return
	}
	
	var hash string
	if HashPassword(c, &hash, body.Password) != 0 {
		return
	}

	var user models.User = models.User{Username: body.Username, Email: body.Email, Password: string(hash)}
	if database.CreateUser(c, user) != 0 {
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func LogIn(c *gin.Context){
	var body struct{
		Email 	 string
		Password string
	}
	
	if Bind(c, &body) != 0 {
		return
	}
	
	var user models.User
	if database.CheckExistencebyEmail(c,&user, body.Email) != 0{
		return
	}

	if CheckPassword(c, user, body.Password) != 0{
		return
	}
	
	var tokenString string
	if GenerateToken(c, user, &tokenString) != 0{
		return
	}
	
	SetCookie(c, tokenString)

	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context){
	user, _ := c.Get("user")

	// user.(models.User)

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}