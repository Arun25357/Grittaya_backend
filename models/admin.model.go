package models

import (
	"github.com/google/uuid"
)

type Admin struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Username string    `gorm:"type:varchar(55);uniqueIndex;not null"`
	Password string    `gorm:"not null"`
	Nickname string		`gorm:"type:varchar(25)"`
	Position string  	`gorm:"not null"`
	Phone    string    `gorm:"type:varchar(15);not null"`
}

type AdminSignInInput struct {
	Username string    `json:"username" binding:"required"`
	Password string    `json:"password" binding:"required"`
}

type GetAdmin struct {
	ID		 uuid.UUID 
    Username  string    `json:"username,omitempty"`
    Nickname string		`json:"type:varchar(25)"`
	Position string  	`json:"not null"`
}

type ForgotAdminPasswordInput struct {
	Phone string `gorm:"type:varchar(15);not null"`
}

type ResetAdminPasswordInput struct {
	Password        string `gorm:"password" binding:"required,min=8"`
	PasswordConfirm string `gorm:"passwordConfirm" binding:"required"`
}
