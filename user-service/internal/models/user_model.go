package models

import (

	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username  string    `json:"username" db:"username"`
    Email     string    `json:"email"`
    Password  string    `json:"password" gorm:"unique"`
    ChatIDs   []int     `json:"chat_ids" db:"chat_ids"`
}