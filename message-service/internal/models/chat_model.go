package models

import "gorm.io/gorm"

type Chat struct {
    gorm.Model
    // IDs de usuarios que participan (sin traer entidades completas)
    ParticipantIDs []uint `json:"participantIds" gorm:"-"`
    // Relación con mensajes (propios del microservicio)
    Messages []Message `json:"messages" gorm:"foreignKey:ChatID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}