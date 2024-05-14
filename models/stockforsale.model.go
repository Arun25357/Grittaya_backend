package models

import uuid "github.com/satori/go.uuid"

type Stockforsale struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Status       int
	Type         string    `gorm:"type:varchar(55);not null"`
	Category     string    `gorm:"type:varchar(55);not null"`
	SetproductID uuid.UUID `gorm:"type:varchar(255);not null"`
}
