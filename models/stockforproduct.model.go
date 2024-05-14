package models

import uuid "github.com/satori/go.uuid"

type Stockforproduct struct {
	ProductID     uuid.UUID `gorm:"type:varchar(255);not null"`
	ProductName   string    `gorm:"type:varchar(55);not null"`
	ProductAmount int       `gorm:"not null"`
	UnitPrice     int       `gorm:"not null"`
	Status        int
}
