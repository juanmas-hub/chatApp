package service

import (
	"chatApp/user-service/internal/models"
	"chatApp/user-service/internal/database"
	"github.com/gin-gonic/gin"
)

func SignUp(username string, email string, password string) int {
	var hash string
	if HashPassword(&hash, password) != 0 {
		return 1
	}

	var user models.User = models.User{Username: username, Email: email, Password: string(hash)}
	if database.CreateUser(&user) != 0 {
		return 1
	}
	return 0
}

func LogIn(c *gin.Context, email string, password string) int{

	var user models.User
	if database.CheckExistencebyEmail(&user, email) != 0{
		return 1
	}

	if CheckPassword(user, password) != 0{
		return 1
	}
	
	var tokenString string
	if GenerateToken(user, &tokenString) != 0{
		return 1
	}
	
	SetCookie(c, tokenString)

	return 0
}