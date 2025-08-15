package models

import "gorm.io/gorm"

type Message struct {
    gorm.Model
    ChatID  uint   `json:"chatId" gorm:"column:chat_id"`
    OwnerID uint   `json:"ownerId" gorm:"column:owner_id"` // ID del usuario emisor
    Content string `json:"content" gorm:"column:content"`
}
