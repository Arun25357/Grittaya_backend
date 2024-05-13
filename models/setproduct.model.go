package models

import uuid "github.com/satori/go.uuid"

type Setproduct struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string    `gorm:"type:varchar(55);not null"`
	Amount    int       `gorm:"not null"`
	UnitPrice int       `gorm:"not null"`
	ProductID uuid.UUID `gorm:"type:varchar(255);not null"`
	Status    int
}
