package initializers

import (
	"user-service/internal/models"
)

func SyncDatabase(){
	DB.AutoMigrate(&models.User{})
}