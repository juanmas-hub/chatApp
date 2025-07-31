package main

import (
	"user-service/controllers"
	"user-service/initializers"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main(){
	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.Run()
}