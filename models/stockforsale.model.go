package models

import uuid "github.com/satori/go.uuid"

type Stockforsale struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name         string    `gorm:"type:varchar(55);not null"`
	Amount       int       `gorm:"not null"`
	UnitPrice    int       `gorm:"not null"`
	Status       int
	Category     string    `gorm:"type:varchar(55);not null"`
	SetproductID uuid.UUID `gorm:"type:varchar(255);not null"`
}
