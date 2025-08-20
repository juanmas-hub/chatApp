package service

import (
	"chatApp/message-service/internal/database"
)

func CreateChat(users []uint) int{
	if database.CreateChat(users) != 0{
		return 1
	}
	return 0
}