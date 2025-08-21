package controllers

import (
	"net/http"
	"chatApp/user-service/internal/service"
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
	
	if service.SignUp(body.Username, body.Email, body.Password) != 0{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to SignUp",
		})
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
	
	if service.LogIn(c, body.Email, body.Password) != 0{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to LogIn",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context){
	user, _ := c.Get("user")

	// user.(models.User)

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}