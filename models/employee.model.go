package models

import (
	
	"github.com/google/uuid"
)

type Employee struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Username	string	`gorm:"type:varchar(55);uniqueIndex;not null"`
	Password	string	`gorm:"uniqueIndex;not null"`
}

