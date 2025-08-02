package initializers

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb(){
	var err error

	// usuario:contrasenia@tcp(host:puerto)/basededatos?opciones
  	dsn := os.Getenv("DB")
  	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{
		panic("fail to connect to db")
	}
}
