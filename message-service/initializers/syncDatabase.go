package initializers

import (
	"chatApp/message-service/internal/models"
)

func SyncDatabase(){
	DB.AutoMigrate(&models.Message{})
	DB.AutoMigrate(&models.Chat{})
}