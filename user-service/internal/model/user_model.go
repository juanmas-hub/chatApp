package model

import(
	"time"
)

type User struct {
    ID        int       `json:"id" db:"id"`
    Username  string    `json:"username" db:"username"`
    Email     string    `json:"email" db:"email"`
    Password  string    `json:"password" db:"password"`
    ChatIDs   []int     `json:"chat_ids" db:"chat_ids"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}