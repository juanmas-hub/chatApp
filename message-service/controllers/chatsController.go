package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"chatApp/message-service/internal/service"
)

func CreateChat(c *gin.Context) {
	var body struct{
		users []uint
	}

	if Bind(c, &body) != 0 {
		return
	}

	if service.CreateChat(body.users) != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user",
		})
		return
	}
	

	c.JSON(http.StatusOK, gin.H{})
}