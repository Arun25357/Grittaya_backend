package models

import (
	"github.com/google/uuid"
)

type Admin struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Username string    `gorm:"type:varchar(55);uniqueIndex;not null"`
	Password string    `gorm:"uniqueIndex;not null"`
	Phone    string    `gorm:"type:varchar(15);not null"`
}

type ForgotAdminPasswordInput struct {
	Phone string `gorm:"type:varchar(15);not null"`
}

type ResetAdminPasswordInput struct {
	Password        string `gorm:"password" binding:"required,min=8"`
	PasswordConfirm string `gorm:"passwordConfirm" binding:"required"`
}
