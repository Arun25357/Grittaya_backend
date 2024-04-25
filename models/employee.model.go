package models

import (
	"github.com/google/uuid"
)

type Employee struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Username string    `gorm:"type:varchar(55);uniqueIndex;not null"`
	Nickname string    `gorm:"type:varchar(25)"`
	Position int       `gorm:"not null"`
	Password string    `gorm:"not null"`
}

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetUser struct {
	ID       uuid.UUID `json:"id,omitempty" uri:"id"`
	Username string    `json:"username,omitempty"`
	Nickname string    `json:"nickname,omitempty"`
	Position int       `json:"position,omitempty"`
}
