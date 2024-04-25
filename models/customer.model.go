package models

import (
	"github.com/satori/go.uuid"
)

type Customer struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Username string    `gorm:"type:varchar(55);not null"`
	Phone    string    `gorm:"type:varchar(15);not null"`
	Email    string    `gorm:"type:varchar(55);uniqueIndex;not null"`
	Location string    `gorm:"not null"`
}
