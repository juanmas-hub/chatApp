package initializers

import (
	"chatApp/user-service/internal/models"
)

func SyncDatabase(){
	DB.AutoMigrate(&models.User{})
}