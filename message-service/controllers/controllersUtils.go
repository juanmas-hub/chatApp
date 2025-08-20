package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Bind[T any](c *gin.Context, body *T) int {
	if c.Bind(body) != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return 1
	}
	return 0
}