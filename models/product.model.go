package models

import "github.com/satori/go.uuid"

type Product struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string    `gorm:"type:varchar(55);not null"`
	Amount    int       `gorm:"not null"`
	UnitPrice int       `gorm:"not null"`
	Type      string    `gorm:"type:varchar(15);not null"`
}

type UpdateProduct struct {
	Name      string `gorm:"type:varchar(55);not null"`
	Amount    int    `gorm:"not null"`
	UnitPrice int    `gorm:"not null"`
	Type      string `gorm:"type:varchar(15);not null"`
}
