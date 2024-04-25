package models

import "github.com/google/uuid"

type Token struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	User_ID   string
	Token     string
	CreatedAt int64 `json:"created_at"`
}
