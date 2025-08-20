package main

import (
	"chatApp/message-service/controllers"
	"chatApp/message-service/initializers"
	//"chatApp/message-service/middleware"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main(){
	r := gin.Default()

	r.POST("/create-chat", controllers.CreateChat)
	r.Run()
}