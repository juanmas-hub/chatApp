package database

import (
	"chatApp/message-service/initializers"
	"chatApp/message-service/internal/models"
)

func CreateChat(users []uint) int {
	var chat models.Chat = models.Chat{}

	chat.UsersIDs = users

	result := initializers.DB.Create(&chat)

	if result.Error != nil{
		return 1
	}
	return 0
}