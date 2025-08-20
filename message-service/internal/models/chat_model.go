package models

import "gorm.io/gorm"

type Chat struct {
    gorm.Model
    // IDs de usuarios que participan (sin traer entidades completas)
    UsersIDs []uint `json:"usersIds" gorm:"-"`
    // Relación con mensajes (propios del microservicio)
    Messages []Message `json:"messages" gorm:"foreignKey:ChatID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}