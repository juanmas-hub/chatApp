package models

import "gorm.io/gorm"


type Chat struct {
    gorm.Model
    Messages []Message `json:"messages" gorm:"foreignKey:ChatID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Message struct {
    gorm.Model
    ChatID  uint   `json:"chatId"  gorm:"column:chat_id"`
    OwnerID int    `json:"ownerId" gorm:"column:owner_id"`
    Content string `json:"content" gorm:"column:content"`
}

