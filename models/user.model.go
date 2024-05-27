package models

import uuid "github.com/satori/go.uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Username string    `gorm:"type:varchar(55);uniqueIndex;not null"`
	Password string    `gorm:"not null"`
	Nickname string    `gorm:"type:varchar(25)"`
	Position int       `gorm:"not null"`
	Phone    string    `gorm:"type:varchar(10);not null"`
}

type UserSignUpInput struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}

type UserSignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetUser struct {
	ID       uuid.UUID
	Username string `json:"username,omitempty"`
	Nickname string `json:"type:varchar(25)"`
	Position string `json:"not null"`
}

type ForgotUserPasswordInput struct {
	Phone string `gorm:"type:varchar(10);not null"`
}

type ResetUserPasswordInput struct {
	Password        string `gorm:"password" binding:"required,min=8"`
	PasswordConfirm string `gorm:"passwordConfirm" binding:"required"`
}

type Userlogout struct {
	
}
